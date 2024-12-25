package tasks

import (
	"DistributedQueueSystem/internal/services"
	"context"
	"encoding/json"
	"errors"
	"github.com/The0ne10/grpc-for-DQS/grpc-for-DQS/queue"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"log/slog"
)

type serverApi struct {
	log *slog.Logger
	queue.UnimplementedQueueServiceServer
	services *services.ServiceContainer
}

func Register(log *slog.Logger, gRPC *grpc.Server, services *services.ServiceContainer) {
	queue.RegisterQueueServiceServer(gRPC, &serverApi{
		log:      log,
		services: services,
	})
}

func (s *serverApi) PushTask(
	ctx context.Context,
	request *queue.PushTaskRequest,
) (*queue.PushTaskResponse, error) {
	const op = "transport.tasks.PushTask"

	log := s.log.With(slog.String("op", op))
	log.Info("Received PushTask request", "queue_name", request.QueueName, "task_id", request.TaskId)

	// Валидация входных данных
	if err := validatePushTask(op, log, request); err != nil {
		return nil, err
	}

	// Проверка JSON payload
	var payload map[string]interface{}
	if err := json.Unmarshal([]byte(request.TaskPayload), &payload); err != nil {
		log.Error("Invalid task payload", "error", err)
		return nil, errors.New("task_payload must be valid JSON")
	}

	// Обработка задачи (пример: добавление в очередь)
	log.Info("Processing task", "task_id", request.TaskId, "priority", request.Priority)

	// Логика добавления в очередь, в данном случае эмуляция
	taskProcessed := true // Например, задача обработана успешно
	if !taskProcessed {
		log.Error("Failed to process task", "task_id", request.TaskId)
		return nil, errors.New("failed to process task")
	}

	// Формирование ответа
	response := &queue.PushTaskResponse{
		Success: true,
	}

	log.Info("Task successfully added to tasks", "queue_name", request.QueueName, "task_id", request.TaskId)

	return response, nil
}

func validatePushTask(op string, log *slog.Logger, request *queue.PushTaskRequest) error {
	if request.QueueName == "" {
		log.Error("Queue name is empty", op)
		return status.Errorf(codes.InvalidArgument, "queue_name is required")
	}

	if request.TaskId == "" {
		log.Error("Task ID is empty", op)
		return status.Errorf(codes.InvalidArgument, "task_id must be provided")
	}

	return nil
}
