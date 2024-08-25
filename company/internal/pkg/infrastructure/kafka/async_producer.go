package kafka

import (
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
	"go.uber.org/zap"
)

func NewAsyncProducer(brokers []string, logger *zap.Logger, opts ...Option) (sarama.AsyncProducer, error) {
	config := prepareProducerSaramaConfig(opts...)

	var asyncProducer sarama.AsyncProducer
	var err error

	for i := 0; i < 20; i++ {
		asyncProducer, err = sarama.NewAsyncProducer(brokers, config)
		if err != nil {
			time.Sleep(500 * time.Microsecond)
		} else {
			err = nil
			break
		}
	}

	if err != nil {
		return nil, errors.Wrap(err, "error with async kafka-producer")
	}

	go func() {
		for err := range asyncProducer.Errors() {
			logger.Error(err.Error())
		}
	}()

	go func() {
		for msg := range asyncProducer.Successes() {
			logger.Info("Message was produced", zap.String("Topic", msg.Topic))
		}
	}()

	return asyncProducer, nil
}
