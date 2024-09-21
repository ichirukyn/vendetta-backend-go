package config

import (
	"github.com/caarlos0/env/v6"
	"github.com/joho/godotenv"
	"log"
)

const (
	AppModeDevelopment = "DEV"
	AppModeDebug       = "DEBUG"
	AppModeProduction  = "PROD"
)

// Config структура env переменных
type Config struct {
	AppMode string `env:"APP_MODE"`

	BindAddress        string `env:"BIND_ADDRESS,required"`
	MetricsBindAddress string `env:"METRICS_BIND_ADDRESS,required"`

	PostgresHost     string `env:"POSTGRES_HOST"`
	PostgresPort     int    `env:"POSTGRES_PORT"`
	PostgresUser     string `env:"POSTGRES_USER"`
	PostgresPassword string `env:"POSTGRES_PASSWORD"`
	PostgresDatabase string `env:"POSTGRES_DATABASE"`

	RabbitMQPort     int    `env:"RABBITMQ_PORT,required"`
	RabbitMQHost     string `env:"RABBITMQ_HOST,required"`
	RabbitMQPath     string `env:"RABBITMQ_PATH,required"`
	RabbitMQUser     string `env:"RABBITMQ_USER,required"`
	RabbitMQPassword string `env:"RABBITMQ_PASSWORD,required"`
}

// NewConfig инициализатор Config
func NewConfig(filename ...string) *Config {
	if err := godotenv.Load(filename...); err != nil {
		log.Fatal("Error loading the .env file ", err.Error())
	}

	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		log.Fatalf("unable to parse environment variables: %e", err)
	}

	if len(cfg.AppMode) == 0 {
		cfg.AppMode = AppModeDevelopment
	}

	return cfg

}
