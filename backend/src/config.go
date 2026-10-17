package main

import (
	"fmt"
	"os"
	"strconv"
)

// Config - настройки приложения, читаются из переменных окружения.
type Config struct {
	Port           string
	DBHost         string
	DBPort         string
	DBUser         string
	DBPassword     string
	DBName         string
	UploadDir      string
	CORSOrigins    string
	GinMode        string
	MaxUploadBytes int64
	DefaultLimit   int
	MaxLimit       int
}

func LoadConfig() Config {
	return Config{
		Port:           env("PORT", "8080"),
		DBHost:         env("DB_HOST", "localhost"),
		DBPort:         env("DB_PORT", "5432"),
		DBUser:         env("DB_USER", "cafe"),
		DBPassword:     env("DB_PASSWORD", ""),
		DBName:         env("DB_NAME", "cafe"),
		UploadDir:      env("UPLOAD_DIR", "./uploads"),
		CORSOrigins:    env("CORS_ALLOWED_ORIGINS", "*"),
		GinMode:        env("GIN_MODE", "release"),
		MaxUploadBytes: int64(envInt("MAX_UPLOAD_MB", 5)) << 20,
		DefaultLimit:   envInt("DEFAULT_PAGE_LIMIT", 20),
		MaxLimit:       envInt("MAX_PAGE_LIMIT", 500),
	}
}

func (c Config) DSN() string {
	return fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		c.DBHost, c.DBPort, c.DBUser, c.DBPassword, c.DBName)
}

func env(key, def string) string {
	if v := os.Getenv(key); v != "" {
		return v
	}
	return def
}

func envInt(key string, def int) int {
	if v, err := strconv.Atoi(os.Getenv(key)); err == nil && v > 0 {
		return v
	}
	return def
}
