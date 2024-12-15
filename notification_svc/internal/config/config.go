package config

import "os"

type KafkaConfig struct {
	KafkaURL string
	Topic    string
	GroupID  string
}

func LoadKafkaConfig() KafkaConfig {
	kafkaURL := os.Getenv("kafkaURL")
	topic := os.Getenv("topic")
	groupID := os.Getenv("groupID")

	return KafkaConfig{KafkaURL: kafkaURL, Topic: topic, GroupID: groupID}
}
