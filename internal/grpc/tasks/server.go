package tasks

import (
	"DistributedQueueSystem/internal/services"
	"context"
	"github.com/The0ne10/grpc-for-DQS/grpc-for-DQS/queue"
	"google.golang.org/grpc"
)

type serverApi struct {
	queue.UnimplementedQueueServiceServer
	services *services.ServiceContainer
}

func Register(gRPC *grpc.Server, services *services.ServiceContainer) {
	queue.RegisterQueueServiceServer(gRPC, &serverApi{
		services: services,
	})
}

func (s *serverApi) PushTask(
	ctx context.Context,
	request *queue.PushTaskRequest,
) (*queue.PushTaskResponse, error) {
	panic("implement me")
}
