package config

import (
	"fmt"
	"os"
	"strings"
	"sync/atomic"

	"github.com/go-playground/validator/v10"
	"github.com/spf13/viper"
)

var config atomic.Value

func InitConfig() error {
	env := os.Getenv("APP_ENV")
	if env == "" {
		env = "dev"
	}

	v := viper.New()
	v.SetConfigName(
		fmt.Sprintf("config.%s", env),
	)
	v.SetConfigType("yaml")
	v.AddConfigPath("./configs")
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	v.AutomaticEnv()

	var cfg Config

	if err := v.ReadInConfig(); err != nil {
		return err
	}

	if err := v.Unmarshal(&cfg); err != nil {
		return err
	}

	config.Store(&cfg)

	validate := validator.New()

	if err := validate.Struct(cfg);err != nil {
		return err
	}
	
	WatchConfig(v)

	return nil
}

func GetConfig() *Config {
	return config.Load().(*Config)
}