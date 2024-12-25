package services

import "DistributedQueueSystem/internal/services/queue"

type ServiceContainer struct {
	TaskService *queue.QueueService
}

func New() *ServiceContainer {
	return &ServiceContainer{
		TaskService: queue.New(),
	}
}
