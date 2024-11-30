package notificationsvc

import (
	"context"
	"os"

	kafkahandler "booking/notification_svc/internal/kafka_handler"
	// deliverysystem "booking/notification_svc/pkg/deliverySystem"
	logger "booking/notification_svc/pkg/logger"
)

func Run() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("notificationTopic")
	groupID := os.Getenv("groupID")

	// ds := deliverysystem.NewMockDeliverySystem()
	consumer := kafkahandler.GetKafkaReader(kafkaURL, topic, groupID)
	logger := logger.Logger{} //TODO in pkg init
	defer consumer.Close()

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			logger.LogError(err)
		}
		if string(m.Value) != "Success" { //TODO

		}
	}
}
