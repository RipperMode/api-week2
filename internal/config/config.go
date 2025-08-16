package config

import (
	"os"
)

type Config struct {
	Port string
	Env  string
	DB   string
}

func LoadConfig() *Config {
	return &Config{
		Port: getEnv("PORT", "8080"),
		Env:  getEnv("APP_ENV", "dev"),
		DB:   getEnv("DB_URL", "postgres://user:pass@localhost:5432/app"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
