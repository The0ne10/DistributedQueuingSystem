package tasks

import (
	"context"
	"log/slog"
)

type TaskService struct {
	ctx context.Context
	log *slog.Logger
}

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func New() *TaskService {
	return &TaskService{}
}
