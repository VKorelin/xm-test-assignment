package company

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func InternalError(err error) error {
	return status.Errorf(codes.Internal, err.Error())
}

func NotFound(err error) error {
	return status.Errorf(codes.NotFound, err.Error())
}
