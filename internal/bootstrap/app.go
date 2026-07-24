package bootstrap

import (
	"context"
	"gin-scaffold/internal/config"
	"gin-scaffold/internal/infrastructure/database"
	"gin-scaffold/internal/infrastructure/migration"
	mysqlrepo "gin-scaffold/internal/infrastructure/persistence/mysql"
	"gin-scaffold/internal/infrastructure/redis"
	"gin-scaffold/internal/logger"
	"gin-scaffold/internal/repository"

	"go.uber.org/zap"
)

type App struct {
	DB       *database.Mysql
	Redis    *redis.Client
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

	// 初始化数据库
	mysql := database.NewMySQL(
		config.GetConfig().Mysql,
	)

	if err := migration.AutoMigrate(mysql); err != nil {
		logger.Log.Fatal("mysql migrate failed", zap.Error(err))
	}


	if err := mysql.Health(); err != nil {
		logger.Log.Fatal("mysql unavailable", zap.Error(err))
	}

	userRepo := mysqlrepo.NewUserRepository(
		mysql.DB,
	)

	RegisterShutdown(mysql.Close)

	// 初始化Redis
	redisClient := redis.NewClient(
		&config.GetConfig().Redis,
	)

	if err := redisClient.Health(context.Background()); err != nil {
		logger.Log.Fatal("redis unavailable", zap.Error(err))
	}

	RegisterShutdown(redisClient.Close)

	return &App{
		DB:       mysql,
		Redis:    redisClient,
		UserRepo: userRepo,
	}
}

func (a *App) Run() {
	println("server running")
}
