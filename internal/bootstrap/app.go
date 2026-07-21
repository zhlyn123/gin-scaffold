package bootstrap

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/logger"
	"gin-scaffold/internal/infrastructure/database"
)

type App struct {
	DB     *database.Mysql
}

func NewApp() *App {

	// 初始化配置
	if err := config.InitConfig(); err != nil {
		panic(err)
	}

	// 初始化日志
	log := logger.NewLogger()
	logger.InitLogger(log)
	logger.Log.Info("logger initialized")

	mysql := database.NewMySQL(
		config.GetConfig().Mysql,
	)

	return &App{
		DB: mysql,
	}
}

func (a *App) Run() {
	println("server running")
}
