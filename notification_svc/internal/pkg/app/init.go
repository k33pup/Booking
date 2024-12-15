package app

import (
	"context"
	"errors"

	config "github.com/k33pup/Booking/notification_svc/internal/config"
	grpcClient "github.com/k33pup/Booking/notification_svc/internal/pkg/api/grpc"
	deliverysystem "github.com/k33pup/Booking/notification_svc/internal/usecases"
	kafkahandler "github.com/k33pup/Booking/notification_svc/pkg/kafka"
	logger "github.com/k33pup/Booking/notification_svc/pkg/logger"

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
	nsvc.logger.LogInfo("Start notify svc...")
	for {
		m, err := nsvc.kafkaReader.ReadMessage(ctx)
		if errors.Is(err, context.Canceled) {
			nsvc.logger.LogInfo("Context is canceled. Stop reading kafka...")
			return nil
		}
		if err != nil {
			nsvc.logger.LogError(err)
			return err
		}
		if string(m.Value) == "Success" {
			nsvc.deliverySystemClient.SendMail(string(m.Value))
		} else {
			nsvc.deliverySystemClient.SendMail("Payment failed")
		}
	}
}

func (nsvc *NotificationService) Stop() error {
	nsvc.logger.LogInfo("End notify svc...")
	err := nsvc.kafkaReader.Close()
	if err != nil {
		nsvc.logger.LogError(err)
		return err
	}
	return nil
}
