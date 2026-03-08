package repository

import (
	"context"

	"github.com/goginenibhavani2000/task-management-service/internal/domain"
	"gorm.io/gorm"
)

type taskRepo struct {
	db *gorm.DB
}

func NewTaskRepository(db *gorm.DB) domain.TaskRepository {
	return &taskRepo{db: db}
}

func (r *taskRepo) Create(ctx context.Context, task *domain.Task) error {
	return r.db.WithContext(ctx).Create(task).Error
}

func (r *taskRepo) GetByID(ctx context.Context, id string) (*domain.Task, error) {
	var task domain.Task
	if err := r.db.WithContext(ctx).First(&task, "id = ?", id).Error; err != nil {
		return nil, domain.ErrTaskNotFound
	}
	return &task, nil
}
