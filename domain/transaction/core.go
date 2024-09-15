package transaction

import (
	"context"
	
	"gorm.io/gorm"
	"zoo/constants"
)

type (
	transaction struct {
		trx *gorm.DB
	}
)

func Begin(ctx context.Context, db *gorm.DB) (*transaction, context.Context, error) {
	trx := db.WithContext(ctx).Begin()
	if trx.Error != nil {
		return nil, ctx, trx.Error
	}
	
	ctx = context.WithValue(ctx, constants.KeyTransaction, trx)
	return &transaction{trx: trx}, ctx, nil
}

func (t *transaction) Commit() error {
	return t.trx.Commit().Error
}

func (t *transaction) Rollback() error {
	return t.trx.Rollback().Error
}

func (t *transaction) Close(ctx context.Context) error {
	ctx = context.WithValue(ctx, constants.KeyTransaction, nil)
	return nil
}
