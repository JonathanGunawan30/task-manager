package impl

import (
	"context"
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/entity/converter"
	"jonathangunawan30/task-manager/internal/model"
	"jonathangunawan30/task-manager/internal/repository"
	"jonathangunawan30/task-manager/internal/tx"
	"jonathangunawan30/task-manager/internal/usecase"

	"github.com/sirupsen/logrus"
	"golang.org/x/crypto/bcrypt"
)

type UserUsecase struct {
	log            *logrus.Logger
	tx             tx.TxManager
	userRepository repository.UserRepository
}

func NewUserUsecase(log *logrus.Logger, tx tx.TxManager, userRepository repository.UserRepository) usecase.UserUsecase {
	return &UserUsecase{log: log, tx: tx, userRepository: userRepository}
}

func (u *UserUsecase) Register(ctx context.Context, request *entity.UserRegisterRequest) (*entity.UserResponse, error) {
	var userResponse *entity.UserResponse

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(request.Password), bcrypt.DefaultCost)
	if err != nil {
		u.log.Errorf("failed to hashing password: %v", err)
		return nil, err
	}

	err = u.tx.RunInTx(ctx, func(ctx context.Context) error {
		user := &model.User{
			Name:     request.Name,
			Email:    request.Email,
			Password: string(hashedPassword),
		}

		user, err = u.userRepository.Save(ctx, user)
		if err != nil {
			u.log.Errorf("failed to register user: %v", err)
			return err
		}

		userResponse = converter.ToUserResponse(user)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return userResponse, nil
}

func (u *UserUsecase) GetAll(ctx context.Context) ([]*entity.UserResponse, error) {
	var userResponse []*entity.UserResponse

	err := u.tx.RunInTx(ctx, func(ctx context.Context) error {
		users, err := u.userRepository.GetAll(ctx)
		if err != nil {
			u.log.Errorf("failed to get all users: %v", err)
			return err
		}
		userResponse = converter.ToUserResponses(users)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return userResponse, nil
}
