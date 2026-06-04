package converter

import (
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/model"
)

func ToTaskResponse(task *model.Task) *entity.TaskResponse {
	return &entity.TaskResponse{
		ID:          task.ID,
		ProjectID:   task.ProjectID,
		Title:       task.Title,
		Description: task.Description,
		Status:      string(task.Status),
		Priority:    string(task.Priority),
		Deadline:    task.Deadline,
		CreatedAt:   task.CreatedAt,
		UpdatedAt:   task.UpdatedAt,
	}
}

func ToTaskResponses(tasks []*model.Task) []*entity.TaskResponse {
	var responses []*entity.TaskResponse
	for _, task := range tasks {
		responses = append(responses, ToTaskResponse(task))
	}
	return responses
}
