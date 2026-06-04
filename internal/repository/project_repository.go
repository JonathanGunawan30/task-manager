package repository

import (
	"context"
	"jonathangunawan30/task-manager/internal/model"
)

type ProjectRepository interface {
	Save(ctx context.Context, project *model.Project) (*model.Project, error)
	GetAll(ctx context.Context, userID int) ([]*model.Project, error)
	GetDetail(ctx context.Context, projectID, userID int) (*model.Project, error)
	Update(ctx context.Context, project *model.Project) (*model.Project, error)
	Delete(ctx context.Context, projectID, userID int) error
}
