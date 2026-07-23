package database

import (
	"fmt"
	"gin-scaffold/internal/config"
	"time"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Mysql struct {
	DB *gorm.DB
}

func NewMySQL(cfg config.MysqlConfig) *Mysql {
	dsn := fmt.Sprintf(

		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",

		cfg.User,

		cfg.Password,

		cfg.Host,

		cfg.Port,

		cfg.DBName,
	)

	GormLogger := NewGormLogger(
		ConvertGormLogLevel(
			config.GetConfig().Log.Level,
		),
		500*time.Millisecond,
	)

	db, err := gorm.Open(
		mysql.Open(dsn), 
		&gorm.Config{
			Logger: GormLogger,
		},
	)
	if err != nil {
		panic(err)
	}

	sqlDB, err := db.DB()
	if err != nil {
		panic(err)
	}

	SetPool(
		sqlDB,
		cfg.MaxOpenConn,
		cfg.MaxIdleConn,
		cfg.MaxLifetime,
	)

	return &Mysql{
		DB: db,
	}
}
