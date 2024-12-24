package services

import "DistributedQueueSystem/internal/services/tasks"

type ServiceContainer struct {
	TaskService *tasks.TaskService
}

func New() *ServiceContainer {
	return &ServiceContainer{
		TaskService: tasks.NewTaskService(),
	}
}
