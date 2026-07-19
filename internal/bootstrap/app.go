package bootstrap

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/logger"
)

type App struct {
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

	return &App{}
}

func (a *App) Run() {
	println("server running")
}
