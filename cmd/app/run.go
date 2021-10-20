package main

import (
	"context"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/pot876/golang-grpc-example/internal/api"
	"github.com/pot876/golang-grpc-example/internal/api/proto"
	"github.com/pot876/golang-grpc-example/internal/config"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func f0() int {
	ctx, cancel := context.WithCancel(context.Background())

	var wg sync.WaitGroup
	grpcRunHelper(ctx, cancel, &wg)
	httpRunHelper(ctx, cancel, &wg)

	systemStopSignals := make(chan os.Signal, 1)
	signal.Notify(systemStopSignals, syscall.SIGINT, syscall.SIGTERM)

	select {
	case <-systemStopSignals:
		logrus.Infof("stop signal recieved, stopping...")
		cancel()
	case <-ctx.Done():
	}

	wg.Wait()
	return 0
}

func grpcRunHelper(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	f := func(ctx context.Context) {
		lis, err := net.Listen("tcp", config.Cfg.GrpcAddr)
		if err != nil {
			logrus.Errorf("net.Listen err:[%v]", err)
			return
		}

		grpcServer := grpc.NewServer()
		proto.RegisterFiboServer(grpcServer, &api.GrpcImplementation{})

		logrus.Infof("starting grpc server, addr: [%s]", config.Cfg.GrpcAddr)
		if err := runGrpcServer(ctx, lis, grpcServer, config.Cfg.GrpcShutdownTime); err != nil {
			logrus.Errorf("grpc server finished with err: [%v]", err)
		}
	}

	runAdopter(ctx, cancel, wg, f)
}

func httpRunHelper(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	f := func(ctx context.Context) {
		if err := runGinServer(ctx, config.Cfg.HttpAddr, api.CreateRestGin(), config.Cfg.HttpShutdownTime); err != nil {
			logrus.Errorf("http server finished with err: [%v]", err)
		}
	}
	runAdopter(ctx, cancel, wg, f)
}

func runAdopter(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup, f func(context.Context)) {
	wg.Add(1)
	go func() {
		f(ctx)
		cancel()
		wg.Done()
	}()
}

func runGinServer(ctx context.Context, addr string, ginEngine *gin.Engine, sdTimeout time.Duration) error {
	srv := &http.Server{
		Addr:           addr,
		Handler:        ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   25 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	var runErrorChan = make(chan error, 1)
	var runError, sdError error
	go func() {
		logrus.Infof("starting gin server, addr: [%s]", addr)
		runErrorChan <- srv.ListenAndServe()
	}()

	select {
	case <-ctx.Done():
		sdContext, cancel := context.WithTimeout(context.Background(), sdTimeout)
		defer cancel()

		if sdError = srv.Shutdown(sdContext); sdError == context.DeadlineExceeded {
			logrus.Warnf("http server shutdown deadline exceeded")
			sdError = nil
		}
	case runError = <-runErrorChan:
	}

	if sdError != nil || (runError != nil && runError != http.ErrServerClosed) {
		return fmt.Errorf("runError: [%v], sdError: [%v]", runError, sdError)
	}
	return nil
}

func runGrpcServer(ctx context.Context, lis net.Listener, server *grpc.Server, sdTimeout time.Duration) error {
	var runErrorChan = make(chan error, 1)
	var runError, sdError error
	go func() {
		runErrorChan <- server.Serve(lis)
	}()

	select {
	case <-ctx.Done():
		go server.GracefulStop()

		select {
		case <-time.After(sdTimeout):
			logrus.Warnf("grpc server shutdown deadline exceeded")
		case runError = <-runErrorChan:
		}
	case runError = <-runErrorChan:
	}

	if sdError != nil || runError != nil {
		return fmt.Errorf("runError: [%v], sdError: [%v]", runError, sdError)
	}
	return nil
}
