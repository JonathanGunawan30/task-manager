package repository

import (
	"context"
	"jonathangunawan30/task-manager/internal/model"
)

type UserRepository interface {
	Save(ctx context.Context, user *model.User) (*model.User, error)
	GetAll(ctx context.Context) ([]*model.User, error)
}
