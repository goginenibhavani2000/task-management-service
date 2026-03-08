package domain

import (
	"context"
	"errors"
)

var ErrTaskNotFound = errors.New("task not found")

// Task is our central business entity
type Task struct {
	ID          string `gorm:"primaryKey"`
	Title       string
	Description string
	Completed   bool
	UserID      string
}

// TaskRepository defines how we talk to the database
type TaskRepository interface {
	Create(ctx context.Context, task *Task) error
	GetByID(ctx context.Context, id string) (*Task, error)
}

// TaskService defines our business logic entry points
type TaskService interface {
	CreateTask(ctx context.Context, title, desc, userID string) (*Task, error)
}
