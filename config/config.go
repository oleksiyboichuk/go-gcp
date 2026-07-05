package config

import (
	"fmt"
	"os"
	"strings"

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

	port := getEnv("PORT", ":8080")

	if !strings.HasPrefix(port, ":") {
		port = ":" + port
	}

	return &Config{
		Port:      port,
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
