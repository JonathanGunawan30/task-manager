package repository

import (
	"context"
	"jonathangunawan30/task-manager/internal/model"
)

type TaskRepository interface {
	Save(ctx context.Context, task *model.Task) (*model.Task, error)
	GetAll(ctx context.Context, userID int) ([]*model.Task, error)
	GetDetail(ctx context.Context, taskID, userID int) (*model.Task, error)
	Update(ctx context.Context, task *model.Task, userID int) (*model.Task, error)
	Delete(ctx context.Context, taskID int) error
}
