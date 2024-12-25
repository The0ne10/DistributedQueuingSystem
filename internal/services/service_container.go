package services

import (
	"DistributedQueueSystem/internal/config"
	"DistributedQueueSystem/internal/services/kafka"
	task "DistributedQueueSystem/internal/services/tasks"
	"fmt"
)

type ServiceContainer struct {
	TaskService     *task.TaskService
	ProducerService *kafka.AsyncProducer
}

func New(cfg config.Config) (*ServiceContainer, error) {

	producer, err := kafka.NewProducer(cfg.Kafka.Address)
	if err != nil {
		return nil, fmt.Errorf("create kafka producer: %w", err)
	}

	return &ServiceContainer{
		TaskService:     task.New(),
		ProducerService: producer,
	}, nil
}
