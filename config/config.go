package config

import (
	"os"
)

type Config struct {
	Env        string
	Local      string
	Port       string
	DBType     string
	DBUsername string
	DBPassword string
	DBHost     string
	DBPort     string
	DBName     string
	JWTSecret  string
}

var instance *Config

func Get() *Config {
	if instance == nil {
		config := newConfig()
		instance = &config
	}
	return instance
}

func newConfig() Config {
	return Config{
		Env:        GetEnv("ENV", ""),
		Local:      GetEnv("LOCAL", "false"),
		Port:       GetEnv("PORT", "5001"),
		DBType:     GetEnv("DB_TYPE", ""),
		DBUsername: GetEnv("DB_USERNAME", ""),
		DBPassword: GetEnv("DB_PASSWORD", ""),
		DBHost:     GetEnv("DB_HOST", ""),
		DBPort:     GetEnv("DB_PORT", ""),
		DBName:     GetEnv("DB_NAME", ""),
		JWTSecret:  GetEnv("JWT_SECRET", ""),
	}
}

func GetEnv(key, fallback string) string {
	if val, exists := os.LookupEnv(key); exists {
		return val
	}
	return fallback
}
