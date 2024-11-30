package app

import (
	config "booking/notification_svc/internal/config"
	kafkahandler "booking/notification_svc/internal/kafka_handler"
	grpcClient "booking/notification_svc/internal/pkg/api/grpc"
	deliverysystem "booking/notification_svc/internal/usecases"

	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	deliverySystemClient deliverysystem.IDeliverySystem
	kafkaReader          *kafka.Reader
}

func NewNotificationService() (*NotificationService, error) {
	ds, err := grpcClient.NewMockDeliverySystemClient()
	if err != nil {
		return nil, err
	}
	kafka_cfg := config.LoadKafkaConfig()
	kafkaReader := kafkahandler.GetKafkaReader(kafka_cfg.KafkaURL, kafka_cfg.Topic, kafka_cfg.GroupID)
	return &NotificationService{deliverySystemClient: ds, kafkaReader: kafkaReader}, nil
}
