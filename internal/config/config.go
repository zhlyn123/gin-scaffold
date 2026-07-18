package config

import (
	"github.com/spf13/viper"
)

type Config struct {
	App   AppConfig
	Mysql MysqlConfig
	Redis RedisConfig
	Log   LogConfig
}

type AppConfig struct {
	Name string `mapstructure:"name"`
	Env  string `mapstructure:"env"`
	Port int    `mapstructure:"port"`
}

type MysqlConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	User     string `mapstructure:"user"`
	Password string `mapstructure:"password"`
	DBName   string `mapstructure:"dbname"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host"`
	Port     int    `mapstructure:"port"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
}

type LogConfig struct {
	Level    string `mapstructure:"level"`
	FileName string `mapstructure:"filename"`
}

var Conf Config

func InitConfig(path string) error {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath(path)

	err := v.ReadInConfig()
	if err != nil {
		return err
	}

	err = v.Unmarshal(&Conf)
	if err != nil {
		return err
	}
	return nil
}
