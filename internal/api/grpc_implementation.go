package api

import (
	context "context"
	"fibo-prj/internal/fibo"
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
