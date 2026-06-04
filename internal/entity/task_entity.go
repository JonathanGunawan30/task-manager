package entity

import "time"

type TaskResponse struct {
	ID          int        `json:"id"`
	ProjectID   int        `json:"project_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type TaskCreateRequest struct {
	UserID      int        `json:"user_id" validate:"required"`
	ProjectID   int        `json:"project_id" validate:"required"`
	Title       string     `json:"title" validate:"required,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	Status      string     `json:"status" validate:"required,oneof=todo in_progress done"`
	Priority    string     `json:"priority" validate:"required,oneof=low medium high"`
	Deadline    *time.Time `json:"deadline" validate:"omitempty"`
}

type TaskUpdateRequest struct {
	ID          int        `json:"id" validate:"required"`
	UserID      int        `json:"user_id" validate:"required"`
	ProjectID   int        `json:"project_id" validate:"required"`
	Title       string     `json:"title" validate:"required,max=255"`
	Description *string    `json:"description" validate:"omitempty,max=500"`
	Status      string     `json:"status" validate:"required,oneof=todo in_progress done"`
	Priority    string     `json:"priority" validate:"required,oneof=low medium high"`
	Deadline    *time.Time `json:"deadline" validate:"omitempty"`
}
