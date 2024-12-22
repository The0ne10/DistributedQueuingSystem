package main

import (
	"DistributedQueueSystem/internal/kafka"
	"log"
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

}
