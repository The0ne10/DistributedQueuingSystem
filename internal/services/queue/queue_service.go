package queue

import (
	"context"
	"log/slog"
)

type QueueService struct {
	ctx context.Context
	log *slog.Logger
}

type Response struct {
	Message string `json:"message"`
	Success bool   `json:"success"`
}

func New() *QueueService {
	return &QueueService{}
}
