package database

import (
	"database/sql"
	"time"
)

func SetPool(db *sql.DB, maxOpen int, maxIdle int, maxLife int) {
	db.SetMaxOpenConns(
		maxOpen,
	)
	db.SetMaxIdleConns(
		maxIdle,
	)
	db.SetConnMaxLifetime(
		time.Duration(maxLife) * time.Second,
	)
}
