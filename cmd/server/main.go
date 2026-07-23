package main

import "gin-scaffold/internal/bootstrap"

func main() {
	app := bootstrap.NewApp()

	go app.Run()
	
	bootstrap.Shutdown()
}
