package repository

import (
	"context"

	"gorm.io/gorm"
)

type BaseRepository struct {
	db *gorm.DB
}

func (b *BaseRepository) GetDB(ctx context.Context) *gorm.DB {
	if tx, ok := ctx.Value("tx").(*gorm.DB); ok {
		return tx
	}
	return b.db
}
