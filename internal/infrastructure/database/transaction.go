package database

import (
	"context"

	"gorm.io/gorm"
)

func (m *Mysql) Transaction(
	ctx context.Context,
	fn func(ctx context.Context) error,
) error {
	return m.DB.WithContext(ctx).
		Transaction(
			func(tx *gorm.DB) error {
				return fn(ctx)
			},
		)
}
