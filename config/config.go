package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

// Config struct holds the configuration settings.
type Config struct {
	HTTPPort string

	// PostgreSQL Configuration
	PostgresHost     string
	PostgresPort     int
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	KafkaBrokers     []string
	LOG_PATH         string
}

// Load loads the configuration from environment variables.
func Load() Config {
	if err := godotenv.Load(); err != nil {
		fmt.Println("No .env file found")
	}

	config := Config{}

	config.HTTPPort = cast.ToString(coalesce("HTTP_PORT", ":9090"))

	// PostgreSQL Configuration
	config.PostgresHost = cast.ToString(coalesce("POSTGRES_HOST", "postgres_dock"))
	config.PostgresPort = cast.ToInt(coalesce("POSTGRES_PORT", 5432))
	config.PostgresUser = cast.ToString(coalesce("POSTGRES_USER", "postgres"))
	config.PostgresPassword = cast.ToString(coalesce("POSTGRES_PASSWORD", "example"))
	config.PostgresDB = cast.ToString(coalesce("POSTGRES_DB", "memory"))

	config.KafkaBrokers = cast.ToStringSlice(coalesce("KAFKA_BROKERS", []string{"kafka:9092"}))

	config.LOG_PATH = cast.ToString(coalesce("LOG_PATH", "logs/info.log"))

	return config
}

func coalesce(key string, defaultValue interface{}) interface{} {
	val, exists := os.LookupEnv(key)

	if exists {
		return val
	}

	return defaultValue
}
