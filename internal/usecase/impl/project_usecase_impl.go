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

type ProjectUsecase struct {
	log               *logrus.Logger
	tx                tx.TxManager
	projectRepository repository.ProjectRepository
}

func NewProjectUsecase(log *logrus.Logger, tx tx.TxManager, projectRepository repository.ProjectRepository) usecase.ProjectUsecase {
	return &ProjectUsecase{log: log, tx: tx, projectRepository: projectRepository}
}

func (p *ProjectUsecase) Create(ctx context.Context, request *entity.ProjectCreateRequest) (*entity.ProjectResponse, error) {
	var response *entity.ProjectResponse
	err := p.tx.RunInTx(ctx, func(ctx context.Context) error {

		project := &model.Project{
			UserID:      request.UserID,
			Title:       request.Title,
			Description: request.Description,
		}

		save, err := p.projectRepository.Save(ctx, project)
		if err != nil {
			p.log.Errorf("failed to save project: %v", err)
			return err
		}

		response = converter.ToProjectResponse(save)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return response, nil
}

func (p *ProjectUsecase) GetAll(ctx context.Context, userID int) ([]*entity.ProjectResponse, error) {
	var responses []*entity.ProjectResponse

	err := p.tx.RunInTx(ctx, func(ctx context.Context) error {
		projects, err := p.projectRepository.GetAll(ctx, userID)
		if err != nil {
			p.log.Errorf("failed to get all projects: %v", err)
			return err
		}

		responses = converter.ToProjectResponses(projects)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return responses, nil
}

func (p *ProjectUsecase) GetDetail(ctx context.Context, projectID, userID int) (*entity.ProjectResponse, error) {
	var project *entity.ProjectResponse

	err := p.tx.RunInTx(ctx, func(ctx context.Context) error {
		detail, err := p.projectRepository.GetDetail(ctx, projectID, userID)
		if err != nil {
			p.log.Errorf("failed to get project detail: %v", err)
			return err
		}

		project = converter.ToProjectResponse(detail)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectUsecase) Update(ctx context.Context, request *entity.ProjectUpdateRequest) (*entity.ProjectResponse, error) {
	var project *entity.ProjectResponse

	err := p.tx.RunInTx(ctx, func(ctx context.Context) error {

		payload := &model.Project{
			ID:          request.ID,
			UserID:      request.UserID,
			Title:       request.Title,
			Description: request.Description,
		}

		detail, err := p.projectRepository.Update(ctx, payload)
		if err != nil {
			p.log.Errorf("failed to get project detail: %v", err)
			return err
		}

		project = converter.ToProjectResponse(detail)
		return nil
	})

	if err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectUsecase) Delete(ctx context.Context, projectID, userID int) error {
	return p.tx.RunInTx(ctx, func(ctx context.Context) error {
		err := p.projectRepository.Delete(ctx, projectID, userID)
		if err != nil {
			p.log.Errorf("failed to delete project: %v", err)
			return err
		}
		return nil
	})
}
