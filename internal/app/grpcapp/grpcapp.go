package grpcapp

import (
	"context"
	"log"
	"net"
	"sync"

	pb "github.com/The0ne10/grpc-for-DQS/grpc-for-DQS/queue"

	"google.golang.org/grpc"
)

type GRPCServer struct {
	pb.UnimplementedQueueServiceServer
	queues map[string][]*Task
	mu     sync.Mutex
}

type Task struct {
	ID       string
	Payload  string
	Priority int32
}

func New() (GRPCServer, error) {

}

func (s *GRPCServer) PushTask(ctx context.Context, req *pb.PushTaskRequest) (*pb.PushTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	if _, exists := s.queues[req.QueueName]; !exists {
		s.queues[req.QueueName] = []*Task{}
	}

	task := &Task{
		ID:       req.TaskId,
		Payload:  req.TaskPayload,
		Priority: req.Priority,
	}
	s.queues[req.QueueName] = append(s.queues[req.QueueName], task)

	log.Printf("Task %s added to queue %s", req.TaskId, req.QueueName)
	return &pb.PushTaskResponse{Success: true}, nil
}

func (s *GRPCServer) PopTask(ctx context.Context, req *pb.PopTaskRequest) (*pb.PopTaskResponse, error) {
	s.mu.Lock()
	defer s.mu.Unlock()

	queue, exists := s.queues[req.QueueName]
	if !exists || len(queue) == 0 {
		return nil, nil
	}

	task := queue[0]
	s.queues[req.QueueName] = queue[1:]

	return &pb.PopTaskResponse{
		TaskId:      task.ID,
		TaskPayload: task.Payload,
	}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	pb.RegisterQueueServiceServer(s, &GRPCServer{queues: make(map[string][]*Task)})

	log.Println("Server is running on port 50051")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
