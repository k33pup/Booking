package kafka

import (
	"context"
	"fmt"
	"strings"

	kafka "github.com/segmentio/kafka-go"
)

func NewKafkaReader(kafkaURL, topic, groupID string) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		GroupID:  groupID,
		Topic:    topic,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
	})
}

type KafkaWriter struct {
	writer *kafka.Writer
	topic  string
}

func NewKafkaWriter(kafkaURL, topic string) *KafkaWriter {
	ww := &kafka.Writer{
		Addr:  kafka.TCP(kafkaURL),
		Topic: topic,
	}
	return &KafkaWriter{ww, topic}
}

func (w *KafkaWriter) Write(ctx context.Context, key, value string) error {
	msg := kafka.Message{
		Key:   []byte(key),
		Value: []byte(fmt.Sprint(key)),
	}
	err := w.writer.WriteMessages(context.Background(), msg)
	return err
}
