package usecase

import (
	"context"
	"jonathangunawan30/task-manager/internal/entity"
)

type TaskUsecase interface {
	Create(ctx context.Context, request *entity.TaskCreateRequest) (*entity.TaskResponse, error)
	GetAll(ctx context.Context, userID int) ([]*entity.TaskResponse, error)
	GetDetail(ctx context.Context, taskID, userID int) (*entity.TaskResponse, error)
	Update(ctx context.Context, request *entity.TaskUpdateRequest) (*entity.TaskResponse, error)
	Delete(ctx context.Context, taskID, userID int) error
}
