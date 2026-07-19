package config

import (
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

func WatchConfig(v *viper.Viper) {
	v.WatchConfig()

	v.OnConfigChange(
		func(e fsnotify.Event) {
			println("config changed:", e.Name)

			var cfg Config

			if err := v.Unmarshal(&cfg); err != nil {
				println("unmarshal config failed:", err.Error())
			}
			println("App:",cfg.App.Name)
			config.Store(&cfg)
		},
	)
}