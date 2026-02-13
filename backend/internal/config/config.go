package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	Host        string
	DatabaseURL string
	JWTSecret   string
	AppEnv      string
}

func Load() *Config {
	_ = godotenv.Load()

	return &Config{
		Port:      os.Getenv("PORT"),
		Host:      os.Getenv("HOST"),
		JWTSecret: os.Getenv("JWT_SECRET"),
		AppEnv:    os.Getenv("APP_ENV"),
	}
}
