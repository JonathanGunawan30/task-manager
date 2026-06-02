package model

import "time"

type (
	Status   string
	Priority string
)

const (
	Todo       Status   = "todo"
	InProgress Status   = "in_progress"
	Done       Status   = "done"
	Low        Priority = "low"
	Medium     Priority = "medium"
	High       Priority = "high"
)

type Task struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement;<-:false"`
	ProjectID   int        `gorm:"column:project_id;not null"`
	Title       string     `gorm:"column:title;not null"`
	Description *string    `gorm:"column:description"`
	Status      Status     `gorm:"column:status;default:todo"`
	Priority    Priority   `gorm:"column:priority;default:medium"`
	Deadline    *time.Time `gorm:"column:deadline"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`

	Project Project `gorm:"foreignKey:ProjectID;references:ID"`
}

func (t *Task) TableName() string {
	return "tasks"
}
