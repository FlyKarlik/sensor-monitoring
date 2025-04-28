package config

import (
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/joho/godotenv/autoload"
)

type Config struct {
	SensorServiceConfig
	MigratorConfig
	Infrastructure
}

type SensorServiceConfig struct {
	AuthKey     string `env:"APP__SENSOR_MONITORING__AUTH_KEY"`
	LogLevel    string `env:"APP__SENSOR_MONITORING__LOG_LEVEL"`
	AppMode     string `env:"APP__SENSOR_MONITORING__MODE"`
	ServiceName string `env:"APP__SENSOR_MONITORING__NAME"`
	ServiceHost string `env:"APP__SENSOR_MONITORING__HOST"`
	ServicePort string `env:"APP__SENSOR_MONITORING__PORT"`
}

type MigratorConfig struct {
	LogLevel       string `env:"APP__MIGRATOR__LOG_LEVEL"`
	AppMode        string `env:"APP__MIGRATOR__MODE"`
	ServiceName    string `env:"APP__MIGRATOR__NAME"`
	MigrationsPath string `env:"APP__MIGRATOR__MIGRATIONS_PATH"`
}

type Infrastructure struct {
	Postgres PostgreSQL
}

type PostgreSQL struct {
	Host     string `env:"INFRA__POSTGRES__HOST"`
	Port     string `env:"INFRA__POSTGRES__PORT"`
	User     string `env:"INFRA__POSTGRES__USER"`
	Password string `env:"INFRA__POSTGRES__PASSWORD"`
	Database string `env:"INFRA__POSTGRES__DATABASE"`
	ConnStr  string `env:"INFRA__POSTGRES__CONN_STR"`
}

func New() (*Config, error) {
	cfg := &Config{}
	if err := cleanenv.ReadEnv(cfg); err != nil {
		return nil, err
	}
	return cfg, nil
}
