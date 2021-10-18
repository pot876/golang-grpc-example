package api

import (
	context "context"
	"fibo-prj/internal/fibo"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
)

type GrpcImplementation struct {
	UnimplementedFiboServer
}

func (q *GrpcImplementation) GetFiboNumbers(ctx context.Context, in *FiboRequest) (*FiboReply, error) {
	result, err := fibo.FromTo(in.From, in.To)
	if err != nil {
		return nil, err
	}

	return &FiboReply{
		Numbers: result,
	}, nil
}

func (q *GrpcImplementation) GetFiboNumbersStream(in *FiboRequest, stream Fibo_GetFiboNumbersStreamServer) error {
	// TODO
	for i := 0; i < 1000; i++ {
		if err := stream.Send(&FiboReply{Numbers: []string{fmt.Sprintf("%d", i)}}); err != nil {
			logrus.Warnf("GetFiboNumbersStream send err: [%v]", err)
			return nil
		}
		time.Sleep(time.Second)
	}
	return nil
}
