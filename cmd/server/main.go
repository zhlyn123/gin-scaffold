package main

import(
	"fmt"
	"gin-scaffold/internal/config"
)

func main() {
	err := config.InitConfig("./configs")
	if err != nil {
		panic(err)
	}

	fmt.Println(config.Conf.App.Name)

	fmt.Println(config.Conf.Mysql.Host)

	fmt.Println(config.Conf.Redis.Host)
}