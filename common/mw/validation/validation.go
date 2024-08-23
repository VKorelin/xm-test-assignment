package validation

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type validator interface {
	ValidateAll() error
}

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, logger *zap.Logger) (resp interface{}, err error) {

	if r, ok := req.(validator); ok {
		if err := r.ValidateAll(); err != nil {
			return nil, status.Errorf(codes.InvalidArgument, err.Error())
		}
	}

	return handler(ctx, req)
}
