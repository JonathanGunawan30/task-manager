package usecase

import (
	"context"
	"jonathangunawan30/task-manager/internal/entity"
)

type ProjectUsecase interface {
	Create(ctx context.Context, request *entity.ProjectCreateRequest) (*entity.ProjectResponse, error)
	GetAll(ctx context.Context, userID int) ([]*entity.ProjectResponse, error)
	GetDetail(ctx context.Context, projectID, userID int) (*entity.ProjectResponse, error)
	Update(ctx context.Context, request *entity.ProjectUpdateRequest) (*entity.ProjectResponse, error)
	Delete(ctx context.Context, projectID, userID int) error
}
