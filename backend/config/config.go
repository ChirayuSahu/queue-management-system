package config

import (
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port        string
	AppName     string
	Env         string
	DatabaseUrl string
	JWTSecret   string
}

func LoadConfig() Config {
	_ = godotenv.Load()

	return Config{
		Port:        getEnv("PORT", "8000"),
		AppName:     getEnv("APP_NAME", "LineUp"),
		Env:         getEnv("ENV", "development"),
		DatabaseUrl: getEnv("DATABASE_URL", ""),
		JWTSecret:   getEnv("JWT_SECRET_KEY", ""),
	}
}

func getEnv(key, fallback string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return fallback
}
