package kafka

import (
	"time"

	"github.com/IBM/sarama"
)

func prepareProducerSaramaConfig(opts ...Option) *sarama.Config {
	c := sarama.NewConfig()

	c.Producer.Partitioner = sarama.NewHashPartitioner
	c.Producer.RequiredAcks = sarama.WaitForAll
	c.Producer.Idempotent = true
	c.Net.MaxOpenRequests = 1

	c.Producer.Retry.Max = 3
	c.Producer.Retry.Backoff = 5 * time.Millisecond

	c.Producer.Flush.Messages = 10
	c.Producer.Flush.Frequency = time.Second

	c.Producer.CompressionLevel = sarama.CompressionLevelDefault
	c.Producer.Compression = sarama.CompressionGZIP

	c.Producer.Return.Successes = true
	c.Producer.Return.Errors = true

	return c
}
