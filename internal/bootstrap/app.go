package bootstrap

import "gin-scaffold/internal/config"

type App struct {
	Config *config.Config
}

func NewApp() *App {
	if err := config.InitConfig(); err != nil {
		panic(err)
	}
	return &App{
		Config: config.GetConfig(),
	}
}

func (a *App) Run() {
	println("server running")
	println(a.Config.App.Name)
}