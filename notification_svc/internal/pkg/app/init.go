package app

import (
	config "booking/notification_svc/internal/config"
	grpcClient "booking/notification_svc/internal/pkg/api/grpc"
	deliverysystem "booking/notification_svc/internal/usecases"
	kafkahandler "booking/notification_svc/pkg/kafka"
	"context"
	"fmt"

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
	kafkaReader := kafkahandler.NewKafkaReader(kafka_cfg.KafkaURL, kafka_cfg.Topic, kafka_cfg.GroupID)
	return &NotificationService{deliverySystemClient: ds, kafkaReader: kafkaReader}, nil
}

func (nsvc *NotificationService) Start(ctx context.Context) error {
	m, err := nsvc.kafkaReader.ReadMessage(context.Background())
	if err != nil {
		fmt.Println("Error in reading msg")
		return err
	}
	if string(m.Value) == "Success" {
		nsvc.deliverySystemClient.SendMail(string(m.Value))
	}
	return nil
}

func (nsvc *NotificationService) Close() {
	nsvc.kafkaReader.Close()
}
