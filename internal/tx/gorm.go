package tx

import (
	"context"

	"gorm.io/gorm"
)

type gormTxManager struct {
	db *gorm.DB
}

func NewTx(db *gorm.DB) TxManager {
	return &gormTxManager{db: db}
}

func (t *gormTxManager) RunInTx(ctx context.Context, fn func(ctx context.Context) error) error {
	return t.db.WithContext(ctx).Transaction(func(tx *gorm.DB) error {
		ctx = context.WithValue(ctx, "tx", tx)
		return fn(ctx)
	})
}
