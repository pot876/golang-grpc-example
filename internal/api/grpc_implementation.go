package api

import (
	context "context"
	"fmt"
	"time"

	"github.com/pot876/golang-grpc-example/internal/api/proto"
	"github.com/pot876/golang-grpc-example/internal/fibo"
	"github.com/sirupsen/logrus"
)

type GrpcImplementation struct {
	proto.UnimplementedFiboServer
}

func (q *GrpcImplementation) GetFiboNumbers(ctx context.Context, in *proto.FiboRequest) (*proto.FiboReply, error) {
	result, err := fibo.FromTo(in.From, in.To)
	if err != nil {
		return nil, err
	}

	return &proto.FiboReply{
		Numbers: result,
	}, nil
}

func (q *GrpcImplementation) GetFiboNumbersStream(in *proto.FiboRequest, stream proto.Fibo_GetFiboNumbersStreamServer) error {
	// TODO
	for i := 0; i < 1000; i++ {
		if err := stream.Send(&proto.FiboReply{Numbers: []string{fmt.Sprintf("%d", i)}}); err != nil {
			logrus.Warnf("GetFiboNumbersStream send err: [%v]", err)
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}
