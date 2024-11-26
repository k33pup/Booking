package notificationsvc

import (
	"context"
	"encoding/json"
	"os"

	deliverysystem "github.com/k33pup/Booking.git/internal/notification/delivery_system"
	kafkahandler "github.com/k33pup/Booking.git/internal/notification/kafka_handler"
	logger "github.com/k33pup/Booking.git/pkg/logger"
)

func Run() {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("notificationTopic")
	groupID := os.Getenv("groupID")

	ds := deliverysystem.NewMockDeliverySystem()
	consumer := kafkahandler.GetKafkaReader(kafkaURL, topic, groupID)
	defer consumer.Close()

	for {
		m, err := consumer.ReadMessage(context.Background())
		if err != nil {
			logger.WriteLog(err.Error()) //TODO
		}
		if string(m.Value) == "Success" { //TODO

		}
		str := "Hello, mr Penis payment completed" //TODO
		bytes, err := json.Marshal(str)
		if err != nil {
			logger.WriteLog(err.Error()) //TODO
		}
		ds.SendMessage(bytes) //
	}
}
