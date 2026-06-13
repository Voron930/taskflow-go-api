package config

import (
	"fmt"
	"os"
)

type Config struct {
	AppPort string
	DBHost  string
	DBPort  string
	DBUser  string
	DBPass  string
	DBName  string
	DBSSL   string
}

func Load() (*Config, error) {
	cfg := &Config{
		AppPort: os.Getenv("APP_PORT"),
		DBHost:  os.Getenv("DB_HOST"),
		DBPort:  os.Getenv("DB_PORT"),
		DBUser:  os.Getenv("DB_USER"),
		DBPass:  os.Getenv("DB_PASSWORD"),
		DBName:  os.Getenv("DB_NAME"),
		DBSSL:   os.Getenv("DB_SSLMODE"),
	}

	if cfg.AppPort == "" {
		cfg.AppPort = "8080"
	}

	if cfg.DBHost == "" {
		cfg.DBHost = "localhost"
	}
	if cfg.DBPort == "" {
		cfg.DBPort = "5432"
	}
	if cfg.DBUser == "" {
		cfg.DBUser = "postgres"
	}
	if cfg.DBName == "" {
		cfg.DBName = "taskflow"
	}
	if cfg.DBSSL == "" {
		cfg.DBSSL = "disable"
	}
	if cfg.DBPass == "" {
		return nil, fmt.Errorf("DB_PASSWORD is required")
	}

	return cfg, nil
}

func (c *Config) DatabaseURL() string {
	return fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		c.DBUser,
		c.DBPass,
		c.DBHost,
		c.DBPort,
		c.DBName,
		c.DBSSL,
	)
}
