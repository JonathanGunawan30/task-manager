package entity

import "time"

type ProjectResponse struct {
	ID          int        `json:"id"`
	UserID      int        `json:"user_id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Tasks       []*Task    `json:"tasks"`
	CreatedAt   *time.Time `json:"created_at"`
	UpdatedAt   *time.Time `json:"updated_at"`
}

type Task struct {
	ID          int        `json:"id"`
	Title       string     `json:"title"`
	Description *string    `json:"description"`
	Status      string     `json:"status"`
	Priority    string     `json:"priority"`
	Deadline    *time.Time `json:"deadline"`
}

type ProjectCreateRequest struct {
	UserID      int     `json:"user_id" validate:"required"`
	Title       string  `json:"title" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}

type ProjectUpdateRequest struct {
	ID          int     `json:"id" validate:"required"`
	UserID      int     `json:"user_id" validate:"required"`
	Title       string  `json:"title" validate:"required,max=255"`
	Description *string `json:"description" validate:"omitempty,max=500"`
}
