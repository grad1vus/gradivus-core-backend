package config

import "os"

type Config struct {
	Port        string
	DatabaseURL string
	Env         string
	JWTSecret   string
}

func Load() *Config {
	return &Config{
		Port:        getEnv("PORT", "8081"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/gradivus_core?sslmode=disable"),
		Env:         getEnv("APP_ENV", "development"),
		JWTSecret:   getEnv("JWT_SECRET", "change-this-secret-in-production"),
	}
}

func getEnv(key, fallback string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return fallback
}
