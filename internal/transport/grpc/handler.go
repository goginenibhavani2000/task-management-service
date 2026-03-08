package grpc

import (
	"context"

	"github.com/goginenibhavani2000/task-management-service/internal/domain"
	"github.com/goginenibhavani2000/task-management-service/pb"
)

type TaskHandler struct {
	pb.UnimplementedTaskServiceServer
	svc domain.TaskService
}

func NewTaskHandler(svc domain.TaskService) *TaskHandler {
	return &TaskHandler{svc: svc}
}

func (h *TaskHandler) CreateTask(ctx context.Context, req *pb.CreateTaskRequest) (*pb.TaskResponse, error) {
	// userID would usually be extracted from JWT context here
	userID := "user-123"

	task, err := h.svc.CreateTask(ctx, req.Title, req.Description, userID)
	if err != nil {
		return nil, err
	}

	return &pb.TaskResponse{
		Id:    task.ID,
		Title: task.Title,
	}, nil
}
