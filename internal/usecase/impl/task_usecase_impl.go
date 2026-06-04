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
)

type TaskUsecase struct {
	log               *logrus.Logger
	tx                tx.TxManager
	taskRepository    repository.TaskRepository
	projectRepository repository.ProjectRepository
}

func NewTaskUsecase(log *logrus.Logger, tx tx.TxManager, taskRepository repository.TaskRepository, projectRepository repository.ProjectRepository) usecase.TaskUsecase {
	return &TaskUsecase{log: log, tx: tx, taskRepository: taskRepository, projectRepository: projectRepository}
}

func (t *TaskUsecase) Create(ctx context.Context, request *entity.TaskCreateRequest) (*entity.TaskResponse, error) {
	_, err := t.projectRepository.GetDetail(ctx, request.ProjectID, request.UserID)
	if err != nil {
		return nil, err
	}

	var taskResponse *entity.TaskResponse
	err = t.tx.RunInTx(ctx, func(ctx context.Context) error {

		task := &model.Task{
			ProjectID:   request.ProjectID,
			Title:       request.Title,
			Description: request.Description,
			Status:      model.Status(request.Status),
			Priority:    model.Priority(request.Priority),
			Deadline:    request.Deadline,
		}

		task, err = t.taskRepository.Save(ctx, task)
		if err != nil {
			t.log.Errorf("failed to save task %v", err)
			return err
		}

		taskResponse = converter.ToTaskResponse(task)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return taskResponse, nil
}

func (t *TaskUsecase) GetAll(ctx context.Context, userID int) ([]*entity.TaskResponse, error) {
	var taskResponses []*entity.TaskResponse

	err := t.tx.RunInTx(ctx, func(ctx context.Context) error {
		tasks, err := t.taskRepository.GetAll(ctx, userID)
		if err != nil {
			t.log.Errorf("failed to get all tasks: %v", err)
			return err
		}

		taskResponses = converter.ToTaskResponses(tasks)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return taskResponses, nil
}

func (t *TaskUsecase) GetDetail(ctx context.Context, taskID, userID int) (*entity.TaskResponse, error) {
	var taskResponse *entity.TaskResponse

	err := t.tx.RunInTx(ctx, func(ctx context.Context) error {

		task, err := t.taskRepository.GetDetail(ctx, taskID, userID)
		if err != nil {
			t.log.Errorf("failed to get task detail: %v", err)
			return err
		}

		taskResponse = converter.ToTaskResponse(task)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return taskResponse, nil
}

func (t *TaskUsecase) Update(ctx context.Context, request *entity.TaskUpdateRequest) (*entity.TaskResponse, error) {
	_, err := t.taskRepository.GetDetail(ctx, request.ID, request.UserID)
	if err != nil {
		return nil, err
	}

	var taskResponse *entity.TaskResponse
	err = t.tx.RunInTx(ctx, func(ctx context.Context) error {

		task := &model.Task{
			ID:          request.ID,
			ProjectID:   request.ProjectID,
			Title:       request.Title,
			Description: request.Description,
			Status:      model.Status(request.Status),
			Priority:    model.Priority(request.Priority),
			Deadline:    request.Deadline,
		}

		task, err = t.taskRepository.Update(ctx, task, request.UserID)
		if err != nil {
			t.log.Errorf("failed to save task %v", err)
			return err
		}

		taskResponse = converter.ToTaskResponse(task)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return taskResponse, nil
}

func (t *TaskUsecase) Delete(ctx context.Context, taskID, userID int) error {
	_, err := t.taskRepository.GetDetail(ctx, taskID, userID)
	if err != nil {
		return err
	}

	return t.tx.RunInTx(ctx, func(ctx context.Context) error {
		err = t.taskRepository.Delete(ctx, taskID)
		if err != nil {
			t.log.Errorf("failed to delete task: %v", err)
			return err
		}
		return nil
	})
}
