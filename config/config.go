package config

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/spf13/cast"
)

type Config struct {
	Host     string
	Port     int
	LogLevel string

	PostgresHost       string
	PostgresPort       int
	PostgresDBName     string
	PostgresPassword   string
	PostgresUser       string
	PoolMaxConnections int
}

func GetConfig() (*Config, error) {
	if err := godotenv.Load("./.env"); err != nil {
		return nil, err
	}

	cfg := new(Config)

	cfg.Host = cast.ToString(getEnvOrReturnDefaultValue("Host", "localhost"))
	cfg.Port = cast.ToInt(getEnvOrReturnDefaultValue("Port", 9000))
	cfg.LogLevel = cast.ToString(getEnvOrReturnDefaultValue("LogLevel", "debug"))

	cfg.PostgresHost = cast.ToString(getEnvOrReturnDefaultValue("PostgresHost", "localhost"))
	cfg.PostgresPort = cast.ToInt(getEnvOrReturnDefaultValue("PostgresPort", 5432))
	cfg.PostgresDBName = cast.ToString(getEnvOrReturnDefaultValue("PostgresDBName", "postgres"))
	cfg.PostgresPassword = cast.ToString(getEnvOrReturnDefaultValue("PostgresPassword", "postgres"))
	cfg.PostgresUser = cast.ToString(getEnvOrReturnDefaultValue("PostgresUser", "postgres"))
	cfg.PoolMaxConnections = cast.ToInt(getEnvOrReturnDefaultValue("PoolMaxConnections", 60))

	return cfg, nil
}

func getEnvOrReturnDefaultValue(name string, default_value interface{}) interface{} {
	a, exists := os.LookupEnv(name)
	if exists {
		return a
	}

	return default_value
}
