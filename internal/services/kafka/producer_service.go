package kafka

import (
	"errors"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"log/slog"
	"time"
)

const (
	flushTimeout = 5000 // ms
)

var (
	errUnknownProducerType = errors.New("unknown producer type")
)

type AsyncProducer struct {
	producer sarama.AsyncProducer
	log      *slog.Logger
}

func NewProducer(address []string) (*AsyncProducer, error) {
	const op = "kafka.NewProducer"

	cfg := sarama.NewConfig()
	cfg.Producer.RequiredAcks = sarama.WaitForAll
	cfg.Producer.Retry.Max = 5
	cfg.Producer.Return.Successes = true
	cfg.Producer.Return.Errors = true

	producer, err := sarama.NewAsyncProducer(address, cfg)
	if err != nil {
		return nil, fmt.Errorf("%s: %w", op, err)
	}

	return &AsyncProducer{producer: producer}, nil
}

func (ap *AsyncProducer) Produce(message, topic, key string) error {
	const op = "kafka.Produce"

	kafkaMsg := &sarama.ProducerMessage{
		Topic: topic,
		Value: sarama.StringEncoder(message),
		Key:   sarama.StringEncoder(key),
	}

	// Отправка сообщения
	ap.producer.Input() <- kafkaMsg

	// Ждем события об успешной отправке или ошибке
	select {
	case success := <-ap.producer.Successes():
		log.Printf("Message sent successfully to topic=%s partition=%d offset=%d",
			success.Topic, success.Partition, success.Offset)
		return nil
	case err := <-ap.producer.Errors():
		log.Printf("Failed to send message to topic=%s: %v", err.Msg.Topic, err.Err)
		return fmt.Errorf("%s: %w", op, err.Err)
	}
}

// Close завершает работу продюсера и очищает буферы.
func (ap *AsyncProducer) Close() {
	const op = "kafka.Close"

	log.Println("Flushing and closing Kafka producer...")
	timeout := time.After(flushTimeout)
	done := make(chan struct{})

	// Закрываем продюсер в отдельной горутине
	go func() {
		ap.producer.AsyncClose()
		close(done)
	}()

	select {
	case <-done:
		log.Println("Producer closed gracefully.")
	case <-timeout:
		log.Println("Timeout reached while closing producer.")
	}
}
