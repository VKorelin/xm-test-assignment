package kafka

import (
	"fmt"
	"time"

	"github.com/IBM/sarama"
	"github.com/pkg/errors"
)

func NewAsyncProducer(brokers []string, opts ...Option) (sarama.AsyncProducer, error) {
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
			fmt.Println(err.Error())
		}
	}()

	go func() {
		for msg := range asyncProducer.Successes() {
			fmt.Println("Async success with key", msg.Key)
		}
	}()

	return asyncProducer, nil
}
