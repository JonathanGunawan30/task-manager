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

type ProjectRepository struct {
	log  *logrus.Logger
	base *repository.BaseRepository
}

func NewProjectRepository(log *logrus.Logger, base *repository.BaseRepository) repository.ProjectRepository {
	return &ProjectRepository{log: log, base: base}
}

func (p *ProjectRepository) Save(ctx context.Context, project *model.Project) (*model.Project, error) {
	if err := p.base.GetDB(ctx).WithContext(ctx).Create(project).Error; err != nil {
		return nil, err
	}

	return project, nil
}

func (p *ProjectRepository) GetAll(ctx context.Context, userID int) ([]*model.Project, error) {
	var projects []*model.Project
	if err := p.base.GetDB(ctx).WithContext(ctx).Preload("Task").Where("user_id = ?", userID).Find(&projects).Error; err != nil {
		return nil, err
	}
	return projects, nil
}

func (p *ProjectRepository) GetDetail(ctx context.Context, projectID, userID int) (*model.Project, error) {
	project := &model.Project{}
	if err := p.base.GetDB(ctx).WithContext(ctx).Preload("Task").Where("id = ? AND user_id = ?", projectID, userID).First(project).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, errDto.ProjectNotFound
		}
		return nil, err
	}
	return project, nil
}

func (p *ProjectRepository) Update(ctx context.Context, project *model.Project) (*model.Project, error) {
	result := p.base.GetDB(ctx).WithContext(ctx).Where("id = ? AND user_id = ?", project.ID, project.UserID).Updates(project)
	if result.RowsAffected == 0 {
		return nil, errDto.ProjectNotFound
	}
	if result.Error != nil {
		return nil, result.Error
	}

	detail, err := p.GetDetail(ctx, project.ID, project.UserID)
	if err != nil {
		return nil, err
	}

	return detail, nil
}

func (p *ProjectRepository) Delete(ctx context.Context, projectID, userID int) error {
	result := p.base.GetDB(ctx).WithContext(ctx).Where("id = ? AND user_id = ?", projectID, userID).Delete(&model.Project{})
	if result.RowsAffected == 0 {
		return errDto.ProjectNotFound
	}

	return result.Error
}
