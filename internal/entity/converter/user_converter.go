package converter

import (
	"jonathangunawan30/task-manager/internal/entity"
	"jonathangunawan30/task-manager/internal/model"
)

func ToUserResponse(user *model.User) *entity.UserResponse {
	return &entity.UserResponse{
		ID:        user.ID,
		Name:      user.Name,
		Email:     user.Email,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
	}
}

func ToUserResponses(users []*model.User) []*entity.UserResponse {
	var responses []*entity.UserResponse
	for _, user := range users {
		responses = append(responses, ToUserResponse(user))
	}
	return responses
}
