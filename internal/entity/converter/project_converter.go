package converter

import (
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/model"
)

func ToProjectResponse(project *model.Project) *entity.ProjectResponse {
	return &entity.ProjectResponse{
		ID:          project.ID,
		UserID:      project.UserID,
		Title:       project.Title,
		Description: project.Description,
		Tasks:       ToTaskResponsesFromProject(project.Task),
		CreatedAt:   project.CreatedAt,
		UpdatedAt:   project.UpdatedAt,
	}
}

func ToProjectResponses(projects []*model.Project) []*entity.ProjectResponse {
	var responses []*entity.ProjectResponse
	for _, project := range projects {
		responses = append(responses, ToProjectResponse(project))
	}
	return responses
}

func ToTaskResponseFromProject(task *model.Task) *entity.Task {
	return &entity.Task{
		ID:          task.ID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		Priority:    string(task.Priority),
		Deadline:    task.Deadline,
	}
}

func ToTaskResponsesFromProject(tasks []*model.Task) []*entity.Task {
	var responses []*entity.Task
	for _, task := range tasks {
		responses = append(responses, ToTaskResponseFromProject(task))
	}
	return responses
}
