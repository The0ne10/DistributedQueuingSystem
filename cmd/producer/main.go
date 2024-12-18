package main

import (
	"DistributedQueueSystem/internal/kafka"
	"fmt"
	"github.com/google/uuid"
	"log"
)

const (
	topic = "my-topic"
)

var (
	address = []string{"localhost:9091", "localhost:9092", "localhost:9093"}
)

func main() {
	producer, err := kafka.NewProducer(address)
	if err != nil {
		log.Fatalf("Failed to create producer: %v", err)
	}
	defer producer.Close()

	keys := generateUUID()

	for i := 0; i < 200; i++ {
		msg := fmt.Sprintf("Message #%d", i)
		key := keys[i%len(keys)]
		err = producer.Produce(
			msg,
			topic,
			key,
		)
		if err != nil {
			log.Printf("Failed to produce message: %v", err)
		} else {
			log.Println("Message sent successfully!", i, topic)
		}
	}
}

func generateUUID() [20]string {
	const op = "kafka.generateUUID"

	var uuids [20]string

	for i := 0; i < 20; i++ {
		uuids[i] = uuid.NewString()
	}

	return uuids
}
