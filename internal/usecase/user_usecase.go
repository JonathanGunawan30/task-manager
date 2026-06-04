package usecase

import (
	"context"
	"jonathangunawan30/task-manager/internal/entity"
)

type UserUsecase interface {
	Register(ctx context.Context, request *entity.UserRegisterRequest) (*entity.UserResponse, error)
	GetAll(ctx context.Context) ([]*entity.UserResponse, error)
}
