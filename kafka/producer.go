package kafka

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/awesome-sphere/as-seating/kafka/interfaces"
	"github.com/segmentio/kafka-go"
	"github.com/segmentio/kafka-go/snappy"
)

func pushMessage(topic_name string, key_parition string, value *interfaces.WriterInterface) {
	config := kafka.WriterConfig{
		Brokers:          []string{KAFKA_LOCATION},
		Topic:            topic_name,
		Balancer:         &kafka.LeastBytes{},
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	writer_connector := kafka.NewWriter(config)
	defer writer_connector.Close()

	new_byte_buffer := new(bytes.Buffer)
	json.NewEncoder(new_byte_buffer).Encode(value)

	err := writer_connector.WriteMessages(
		context.Background(),
		kafka.Message{
			Key:   []byte(key_parition),
			Value: new_byte_buffer.Bytes(),
		},
	)
	if err != nil {
		panic(err.Error())
	}
}

func SendToConsumer(theaterID int64, timeslotID int64, seatID int, status string) {
	partitionKey := fmt.Sprintf("%d", theaterID)
	pushMessage(TOPIC, partitionKey, &interfaces.WriterInterface{
		TheaterID:  theaterID,
		TimeSlotID: timeslotID,
		SeatID:     seatID,
		SeatStatus: status,
	})
}
