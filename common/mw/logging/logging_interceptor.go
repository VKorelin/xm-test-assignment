package logging

import (
	"context"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
	"google.golang.org/protobuf/proto"
)

func Interceptor(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, logger *zap.Logger) (resp interface{}, err error) {
	raw, _ := protojson.Marshal((req).(proto.Message))
	logger.Info("Request received", zap.String("method", info.FullMethod), zap.String("req", string(raw)))

	resp, err = handler(ctx, req)
	if resp != nil {
		rawResp, _ := protojson.Marshal((resp).(proto.Message))
		logger.Info("Response", zap.String("method", info.FullMethod), zap.String("res", string(rawResp)))
	}

	if err != nil {
		logger.Info("Response with error", zap.Error(err))
	}

	return resp, err
}
