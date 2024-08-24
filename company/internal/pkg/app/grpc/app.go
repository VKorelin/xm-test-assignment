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

	"github.com/jackc/pgx/v5/pgxpool"
	"go.uber.org/zap"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type config struct {
	port               string
	companyDSN         string
	companyStatusTopic string
	brokers            string
}

func newConfigFromFlags() config {
	const (
		defaultPort               = ":50050"
		defaultCompanyDSN         = "postgresql://postgres:password@company_db:5432/company"
		defaultCompanyStatusTopic = "company_status"
		defaultBrokers            = "kafka-broker-1:9091,kafka-broker-2:9092,kafka-broker-3:9093"
	)

	result := config{}
	flag.StringVar(&result.port, "port", defaultPort, "gRPC port, default: "+defaultPort)
	flag.StringVar(&result.companyDSN, "companyDSN", defaultCompanyDSN, "company DSN, default: "+defaultCompanyDSN)
	flag.StringVar(&result.companyStatusTopic, "compnayStatusTopic", defaultCompanyStatusTopic, "companyStatusTopic, default: "+defaultCompanyStatusTopic)
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

func initDbPool(databaseDSN string, logger *zap.Logger) *pgxpool.Pool {
	ctx := context.Background()

	dbpool, err := pgxpool.New(ctx, databaseDSN)
	if err != nil {
		logger.Fatal("Unable to create connection pool", zap.Error(err))
	}

	if err := dbpool.Ping(ctx); err != nil {
		panic(err)
	}

	return dbpool
}

func registerCompanyServer(grpcServer *grpc.Server, dbpool *pgxpool.Pool, logger *zap.Logger) {
	storage := repository.NewCompanyStorage(dbpool)
	companyRepository := repository.NewCompanyRepository(storage, logger)

	fetchService := services.NewFetchCompanyService(companyRepository)
	createService := services.NewCreateCompanyService(companyRepository)
	updateService := services.NewUpdateCompanyService(companyRepository)
	deleteService := services.NewDeleteCompanyService(companyRepository)

	companyServer := company.NewCompanyServerImpl(fetchService, createService, updateService, deleteService)
	desc.RegisterCompanyServiceServer(grpcServer, companyServer)
}

func (a App) Run() error {

	dbpool := initDbPool(a.config.companyDSN, a.logger)
	defer dbpool.Close()

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

	registerCompanyServer(grpcServer, dbpool, a.logger)

	a.logger.Info("Start server listening", zap.String("address", lis.Addr().String()))

	if err = grpcServer.Serve(lis); err != nil {
		a.logger.Fatal("failed to serve", zap.Error(err))
	}

	return nil
}
