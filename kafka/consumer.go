package kafka

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func initReader(topic_name string, partition int) {
	config := kafka.ReaderConfig{
		Brokers:  []string{KAFKA_LOCATION},
		Topic:    topic_name,
		MinBytes: 10e3, // 10KB
		MaxBytes: 10e6, // 10MB
		GroupID:  "seating-status-consumer",
	}
	reader_connector := kafka.NewReader(config)
	defer reader_connector.Close()

	for {
		msg, err := reader_connector.ReadMessage(context.Background())
		if err != nil {
			panic(err.Error())
		}
		fmt.Printf("Message on %s/%d/%d: %s = %s\n", topic_name, msg.Partition, partition, msg.Key, string(msg.Value))
	}

}

func ReadMessage(topic_name string) {
	for i := 0; i < PARTITION; i++ {
		go initReader(topic_name, i)
	}
}
