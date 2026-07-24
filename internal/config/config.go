package config

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
	Host        string `mapstructure:"host" validate:"required"`
	Port        int    `mapstructure:"port" validate:"required"`
	User        string `mapstructure:"user" validate:"required"`
	Password    string `mapstructure:"password"`
	DBName      string `mapstructure:"dbname" validate:"required"`
	MaxOpenConn int    `mapstructure:"max_open_conn"`
	MaxIdleConn int    `mapstructure:"max_idle_conn"`
	MaxLifetime int    `mapstructure:"max_lifetime"`
}

type RedisConfig struct {
	Host     string `mapstructure:"host" validate:"required"`
	Port     int    `mapstructure:"port" validate:"required"`
	Password string `mapstructure:"password"`
	DB       int    `mapstructure:"db"`
	PoolSize int `mapstructure:"pool_size"`
	MinIdleConn int `mapstructure:"min_idle_conn"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	FileName   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"maxsize"`
	MaxBackups int    `mapstructure:"maxbackups"`
	MaxAge     int    `mapstructure:"maxage"`
	Compress   bool   `mapstructure:"compress"`
}
