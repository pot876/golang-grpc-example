package api

import (
	context "context"
	"fmt"
	"math/big"
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

func (q *GrpcImplementation) GetFiboNumbersStream(in *proto.FiboStreamRequest, stream proto.Fibo_GetFiboNumbersStreamServer) error {
	for i := uint64(0); i < min(2, in.N); i++ {
		time.Sleep(time.Millisecond * 100)
		if err := stream.Send(&proto.FiboStreamReply{Number: fmt.Sprintf("%d", i)}); err != nil {
			logrus.Warnf("GetFiboNumbersStream send err: [%v]", err)
			return nil
		}
	}

	fi := fibo.FiboIterator(big.NewInt(0), big.NewInt(1))
	for i := uint64(2); i < in.N; i++ {
		time.Sleep(time.Millisecond * 100)
		if err := stream.Send(&proto.FiboStreamReply{Number: fi().String()}); err != nil {
			logrus.Warnf("GetFiboNumbersStream send err: [%v]", err)
			return nil
		}
	}
	return nil
}

func min(a, b uint64) uint64 {
	if a < b {
		return a
	}

	return b
}
