package impl

import (
	"context"
	"errors"
	errDto "jonathangunawan30/task-manager/internal/errors"
	"jonathangunawan30/task-manager/internal/model"
	"jonathangunawan30/task-manager/internal/repository"

	"github.com/go-sql-driver/mysql"
	"github.com/sirupsen/logrus"
)

type UserRepository struct {
	log  *logrus.Logger
	base *repository.BaseRepository
}

func NewUserRepository(log *logrus.Logger, base *repository.BaseRepository) repository.UserRepository {
	return &UserRepository{log: log, base: base}
}

func (u *UserRepository) Save(ctx context.Context, user *model.User) (*model.User, error) {
	if err := u.base.GetDB(ctx).WithContext(ctx).Create(user).Error; err != nil {
		var mysqlErr *mysql.MySQLError
		if errors.As(err, &mysqlErr) && mysqlErr.Number == 1062 {
			return nil, errDto.EmailAlreadyExists
		}
		return nil, err
	}
	return user, nil
}

func (u *UserRepository) GetAll(ctx context.Context) ([]*model.User, error) {
	var users []*model.User
	if err := u.base.GetDB(ctx).WithContext(ctx).Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}
