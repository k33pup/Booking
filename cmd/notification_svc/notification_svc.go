package notificationsvc

import (
	"context"
	"encoding/json"
	"os"

	kafkahandler "github.com/k33pup/Booking.git/internal/notification/kafka_handler"
	deliverysystem "github.com/k33pup/Booking.git/pkg/deliverySystem"
	logger "github.com/k33pup/Booking.git/pkg/logger"
)

func Run() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("notificationTopic")
	groupID := os.Getenv("groupID")

	ds := deliverysystem.NewMockDeliverySystem()
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
		str := "Hello, mr Penis payment completed" //TODO
		bytes, err := json.Marshal(str)
		if err != nil {
			logger.LogError(err)
		}
		ds.SendMessage(bytes) //
	}
}
