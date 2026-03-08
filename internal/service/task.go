package service

import (
	"context"
	"errors"

	"github.com/goginenibhavani2000/task-management-service/internal/domain"
	"github.com/google/uuid"
)

type taskService struct {
	repo domain.TaskRepository
}

func NewTaskService(repo domain.TaskRepository) domain.TaskService {
	return &taskService{repo: repo}
}

func (s *taskService) CreateTask(ctx context.Context, title, desc, userID string) (*domain.Task, error) {
	task := &domain.Task{
		ID:          uuid.NewString(),
		Title:       title,
		Description: desc,
		UserID:      userID,
		Completed:   false,
	}

	// Logic: Ensure title isn't empty
	if title == "" {
		return nil, errors.New("title is required")
	}

	if err := s.repo.Create(ctx, task); err != nil {
		return nil, err
	}
	return task, nil
}
