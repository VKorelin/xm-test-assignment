package grpc

import (
	"common/mw/logging"
	panicInterceptor "common/mw/panic"
	"common/mw/validation"
	"context"
	"flag"
	"net"
	"xm/company/internal/pkg/api/company"
	"xm/company/internal/pkg/repository"
	"xm/company/internal/pkg/services"
	desc "xm/company/pkg/api/company/v1"

	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	port             string
	companyDSN       string
	orderStatusTopic string
	brokers          string
}

func newConfigFromFlags() config {
	const (
		defaultPort             = ":50050"
		defaultCompanyDSN       = "postgresql://postgres:password@company_db:5432/company"
		defaultOrderStatusTopic = "company_status"
		defaultBrokers          = "kafka-broker-1:9091,kafka-broker-2:9092,kafka-broker-3:9093"
	)

	result := config{}
	flag.StringVar(&result.port, "port", defaultPort, "gRPC port, default: "+defaultPort)
	flag.StringVar(&result.companyDSN, "companyDSN", defaultCompanyDSN, "company DSN, default: "+defaultCompanyDSN)
	flag.StringVar(&result.orderStatusTopic, "orderStatusTopic", defaultOrderStatusTopic, "orderStatusTopic, default: "+defaultOrderStatusTopic)
	flag.StringVar(&result.brokers, "brokers", defaultBrokers, "brokers, default: "+defaultBrokers)
	flag.Parse()
	return result
}

type unaryInterceptorWithLogger func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler, logger *zap.Logger) (resp any, err error)

type App struct {
	config config
	logger *zap.Logger
}

func NewApp(logger *zap.Logger) *App {
	return &App{
		config: newConfigFromFlags(),
		logger: logger,
	}
}

func registerCompanyServer(grpcServer *grpc.Server) {
	companyRepository := repository.NewOrderRepository()
	fetchService := services.NewFetchCompanyService(companyRepository)

	orderServer := company.NewOrderServerImpl(fetchService)
	desc.RegisterCompanyServiceServer(grpcServer, orderServer)
}

func (a App) Run() error {

	lis, err := net.Listen("tcp", a.config.port)
	if err != nil {
		a.logger.Fatal("failed to listen tcp", zap.String("port", a.config.port), zap.Error(err))
	}

	getUnaryInterceptorWithLogger := func(interceptor unaryInterceptorWithLogger) grpc.UnaryServerInterceptor {
		return func(ctx context.Context, req any, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (resp any, err error) {
			return interceptor(ctx, req, info, handler, a.logger)
		}
	}

	grpcServer := grpc.NewServer(
		grpc.ChainUnaryInterceptor(
			getUnaryInterceptorWithLogger(panicInterceptor.Interceptor),
			getUnaryInterceptorWithLogger(logging.Interceptor),
			getUnaryInterceptorWithLogger(validation.Interceptor),
		),
	)

	reflection.Register(grpcServer)

	registerCompanyServer(grpcServer)

	a.logger.Info("Start server listening", zap.String("address", lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		a.logger.Fatal("failed to serve", zap.Error(err))
	}

	return nil
}
