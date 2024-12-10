package app

import (
	config "booking/notification_svc/internal/config"
	grpcClient "booking/notification_svc/internal/pkg/api/grpc"
	deliverysystem "booking/notification_svc/internal/usecases"
	kafkahandler "booking/notification_svc/pkg/kafka"
	logger "booking/notification_svc/pkg/logger"
	"context"

	"github.com/segmentio/kafka-go"
)

type NotificationService struct {
	deliverySystemClient deliverysystem.IDeliverySystem
	kafkaReader          *kafka.Reader
	logger               *logger.Logger
}

func NewNotificationService() (*NotificationService, error) {
	logger := logger.Logger{}

	ds, err := grpcClient.NewMockDeliverySystemClient()
	if err != nil {
		return nil, err
	}
	kafka_cfg := config.LoadKafkaConfig()
	kafkaReader := kafkahandler.NewKafkaReader(kafka_cfg.KafkaURL, kafka_cfg.Topic, kafka_cfg.GroupID)
	return &NotificationService{deliverySystemClient: ds, kafkaReader: kafkaReader, logger: &logger}, nil
}

func (nsvc *NotificationService) Start(ctx context.Context) error {
	nsvc.logger.LogInfo("Start hotel svc...")
	for {
		if m, err := nsvc.kafkaReader.ReadMessage(context.Background()); err != nil {
			nsvc.logger.LogError(err)
			return err
		} else {
			if string(m.Value) == "Success" {
				nsvc.deliverySystemClient.SendMail(string(m.Value))
			}
		}
	}
}

func (nsvc *NotificationService) Close() {
	nsvc.logger.LogInfo("End hotel svc...")
	nsvc.kafkaReader.Close()
}
