package impl

import (
	"context"
	"errors"
	errDto "jonathangunawan30/task-manager/internal/errors"
	"jonathangunawan30/task-manager/internal/model"
	"jonathangunawan30/task-manager/internal/repository"

	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

type TaskRepository struct {
	log  *logrus.Logger
	base *repository.BaseRepository
}

func NewTaskRepository(log *logrus.Logger, base *repository.BaseRepository) repository.TaskRepository {
	return &TaskRepository{log: log, base: base}
}

func (t *TaskRepository) Save(ctx context.Context, task *model.Task) (*model.Task, error) {
	if err := t.base.GetDB(ctx).WithContext(ctx).Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (t *TaskRepository) GetAll(ctx context.Context, userID int) ([]*model.Task, error) {
	var tasks []*model.Task
	if err := t.base.GetDB(ctx).WithContext(ctx).Joins("JOIN projects ON projects.id = tasks.project_id").Where("projects.user_id = ?", userID).Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (t *TaskRepository) GetDetail(ctx context.Context, taskID, userID int) (*model.Task, error) {
	task := &model.Task{}
	if err := t.base.GetDB(ctx).WithContext(ctx).Joins("JOIN projects ON projects.id = tasks.project_id").Where("projects.user_id = ? AND tasks.id = ?", userID, taskID).First(task).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errDto.TaskNotFound
		}
		return nil, err
	}

	return task, nil
}

func (t *TaskRepository) Update(ctx context.Context, task *model.Task, userID int) (*model.Task, error) {
	result := t.base.GetDB(ctx).WithContext(ctx).Where("id = ?", task.ID).Updates(task)
	if result.RowsAffected == 0 {
		return nil, errDto.TaskNotFound
	}

	if result.Error != nil {
		return nil, result.Error
	}

	detail, err := t.GetDetail(ctx, task.ID, userID)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (t *TaskRepository) Delete(ctx context.Context, taskID int) error {
	result := t.base.GetDB(ctx).WithContext(ctx).Where("id = ?", taskID).Delete(&model.Task{})
	if result.RowsAffected == 0 {
		return errDto.TaskNotFound
	}

	return result.Error
}
