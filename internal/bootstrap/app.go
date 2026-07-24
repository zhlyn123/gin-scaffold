package bootstrap

import (
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/infrastructure/database"
	mysqlrepo "gin-scaffold/internal/infrastructure/persistence/mysql"
	"gin-scaffold/internal/logger"
	"gin-scaffold/internal/repository"
	"gin-scaffold/internal/infrastructure/migration"
	"go.uber.org/zap"
)

type App struct {
	DB       *database.Mysql
	UserRepo repository.UserRepository
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
	RegisterShutdown(logger.Close)

	mysql := database.NewMySQL(
		config.GetConfig().Mysql,
	)

	if err := migration.AutoMigrate(mysql); err != nil {
		logger.Log.Fatal("mysql migrate failed", zap.Error(err))
	}

	userRepo := mysqlrepo.NewUserRepository(
		mysql.DB,
	)

	if err := mysql.Health(); err != nil {
		logger.Log.Fatal("mysql unavailable", zap.Error(err))
	}

	RegisterShutdown(mysql.Close)

	return &App{
		DB:       mysql,
		UserRepo: userRepo,
	}
}

func (a *App) Run() {
	println("server running")
}
