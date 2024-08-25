package kafka

import (
	"context"
	"encoding/json"
	"fmt"
	kafkainf "xm/company/internal/pkg/infrastructure/kafka"
	"xm/company/internal/pkg/services/notifications"

	"github.com/IBM/sarama"
	"github.com/gofrs/uuid"
	"go.uber.org/zap"
)

type OrderStatusProducer interface {
	Input() chan<- *sarama.ProducerMessage
}

type ChangeNotificationService struct {
	logger    *zap.Logger
	producer  OrderStatusProducer
	topicName string
}

func NewKafkaNotificationService(logger *zap.Logger, producer OrderStatusProducer, topicName string) *ChangeNotificationService {
	return &ChangeNotificationService{
		logger:    logger,
		producer:  producer,
		topicName: topicName,
	}
}

func (s *ChangeNotificationService) Notify(ctx context.Context, id uuid.UUID, changeType notifications.ChangeType) error {

	m := notifications.ChangeMessage{
		Id:         id.String(),
		ChangeType: changeType,
	}

	bytes, err := json.Marshal(m)
	if err != nil {
		s.logger.Error("Error occured on marshalling json", zap.Error(err))
		return err
	}

	msg, err := kafkainf.BuildMessage(s.topicName, fmt.Sprint(m.Id), bytes)
	if err != nil {
		s.logger.Error("Error occured on building kafka message", zap.Error(err))
		return err
	}

	s.producer.Input() <- msg

	return nil
}
