package main

import (
	"context"
	"fibo-prj/internal/api"
	"fibo-prj/internal/config"
	"fmt"
	"net"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
)

func f0() int {
	ctx, cancel := context.WithCancel(context.Background())

	systemStopSignals := make(chan os.Signal, 1)
	signal.Notify(systemStopSignals, syscall.SIGINT, syscall.SIGTERM)

	var wg sync.WaitGroup
	wg.Add(2)

	grpcRunHelper(ctx, cancel, &wg)
	httpRunHelper(ctx, cancel, &wg)

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
	go func() {
		grpcServer := grpc.NewServer()
		api.RegisterFiboServer(grpcServer, &api.GrpcImplementation{})

		if err := runGrpcServer(ctx, config.Cfg.GrpcAddr, grpcServer, 5*time.Second); err != nil {
			logrus.Errorf("grpc server finished with err: [%v]", err)
		}

		cancel()
		wg.Done()
	}()
}
func httpRunHelper(ctx context.Context, cancel context.CancelFunc, wg *sync.WaitGroup) {
	go func() {
		if err := runGinServer(ctx, config.Cfg.HttpAddr, api.CreateRestGin(), 5*time.Second); err != nil {
			logrus.Errorf("http server finished with err: [%v]", err)
		}

		cancel()
		wg.Done()
	}()
}

func runGinServer(_ctx context.Context, addr string, ginEngine *gin.Engine, sdTimeout time.Duration) error {
	srv := &http.Server{
		Addr:           addr,
		Handler:        ginEngine,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   25 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()

	var runError error
	var sdError error
	go func() {
		logrus.Infof("starting  gin server, addr: [%s]", addr)
		runError = srv.ListenAndServe()
		cancel()
	}()

	{
		<-ctx.Done()
		sdContext, cancel := context.WithTimeout(context.Background(), sdTimeout)
		defer cancel()

		sdError = srv.Shutdown(sdContext)
	}

	if sdError != nil || (runError != nil && runError != http.ErrServerClosed) {
		return fmt.Errorf("runError: [%v], sdError: [%v]", runError, sdError)
	}

	return nil
}

func runGrpcServer(_ctx context.Context, addr string, server *grpc.Server, sdTimeout time.Duration) error {
	lis, err := net.Listen("tcp", addr)
	if err != nil {
		return fmt.Errorf("failed to listen: %v", err)
	}

	ctx, cancel := context.WithCancel(_ctx)
	defer cancel()

	var runError error
	var sdError error
	go func() {
		logrus.Infof("starting grpc server, addr: [%s]", addr)
		runError = server.Serve(lis)
		cancel()
	}()
	{
		<-ctx.Done()
		sdContext, cancel := context.WithTimeout(context.Background(), sdTimeout)
		defer cancel()

		go func() {
			server.GracefulStop()
			cancel()
		}()

		<-sdContext.Done()
		if sdContext.Err() == context.DeadlineExceeded {
			sdError = context.DeadlineExceeded
		}
	}

	if sdError != nil || runError != nil {
		return fmt.Errorf("runError: [%v], sdError: [%v]", runError, sdError)
	}
	return nil
}