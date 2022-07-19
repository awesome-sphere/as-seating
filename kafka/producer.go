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

func pushMessage(topic string, partition int, value *interfaces.WriterInterface) {
	config := kafka.WriterConfig{
		Brokers:          []string{KAFKA_LOCATION},
		Topic:            topic,
		Balancer:         &PartitionBalancer{},
		WriteTimeout:     10 * time.Second,
		ReadTimeout:      10 * time.Second,
		CompressionCodec: snappy.NewCompressionCodec(),
	}
	writer := kafka.NewWriter(config)
	defer writer.Close()

	byteBuffer := new(bytes.Buffer)
	json.NewEncoder(byteBuffer).Encode(value)

	fmt.Printf("Sending message to %s/%d\n", topic, partition)

	err := writer.WriteMessages(
		context.Background(),
		kafka.Message{
			Partition: partition,
			Value:     byteBuffer.Bytes(),
		},
	)
	if err != nil {
		panic(err.Error())
	}
}

func SendToConsumer(theaterID int64, timeslotID int64, seatID int, status string) {
	pushMessage(TOPIC, int(theaterID), &interfaces.WriterInterface{
		TheaterID:  theaterID,
		TimeSlotID: timeslotID,
		SeatID:     seatID,
		SeatStatus: status,
	})
}
