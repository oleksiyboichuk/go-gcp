package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

type Config struct {
	Port      string
	JWTSecret string
}

func LoadConfig() *Config {
	err := godotenv.Load()
	if err != nil {
		fmt.Println("WARNING: No config found!")
	}

	return &Config{
		Port:      getEnv("PORT", ":8080"),
		JWTSecret: getEnv("JWT_SECRET", "dummy"),
	}
}

func getEnv(key, defaultValue string) string {
	value := os.Getenv(key)
	if value == "" {
		return defaultValue
	}
	return value
}
