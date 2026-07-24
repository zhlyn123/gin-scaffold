package migration

import (
	"gin-scaffold/internal/infrastructure/database"
)

func AutoMigrate(db *database.Mysql) error {

	return db.DB.AutoMigrate(
		Models()...,
	)

}
