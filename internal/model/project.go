package model

import "time"

type Project struct {
	ID          int        `gorm:"column:id;primaryKey;autoIncrement;<-:false"`
	UserID      int        `gorm:"column:user_id"`
	Title       string     `gorm:"column:title;not null"`
	Description *string    `gorm:"column:description"`
	CreatedAt   *time.Time `gorm:"column:created_at"`
	UpdatedAt   *time.Time `gorm:"column:updated_at"`

	User User   `gorm:"foreignKey:UserID;references:ID"`
	Task []Task `gorm:"foreignKey:ProjectID"`
}

func (p *Project) TableName() string {
	return "projects"
}
