package bootstrap

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/logger"
	"net/http"

	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

type App struct {
	server *http.Server
	DB     *gorm.DB
	Redis  *redis.Client
	Logger *zap.Logger
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
	println(config.GetConfig().Mysql.Port)
}
