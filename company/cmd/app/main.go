package main

import (
	"common/logger"
	grpcapp "xm/company/internal/pkg/app/grpc"
)

func main() {
	logger := logger.InitLogger()
	defer logger.Sync()

	grpcApp := grpcapp.NewApp(logger)

	err := grpcApp.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
