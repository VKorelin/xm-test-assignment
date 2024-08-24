package main

import (
	"common/logger"
	"fmt"
	"os"
	grpcapp "xm/company/internal/pkg/app/grpc"
)

func main() {
	logger := logger.InitLogger()

	defer func() {
		if err := logger.Sync(); err != nil {
			fmt.Fprintf(os.Stderr, "Error syncing logger: %v\n", err)
		}
	}()

	grpcApp := grpcapp.NewApp(logger)

	err := grpcApp.Run()
	if err != nil {
		logger.Fatal(err.Error())
	}
}
