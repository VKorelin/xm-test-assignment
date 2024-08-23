package grpc

import (
	"flag"

	"go.uber.org/zap"
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

func (a App) Run() error {
	return nil
}
