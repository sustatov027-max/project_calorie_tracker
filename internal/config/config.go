package config

import (
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/joho/godotenv"
)

type Config struct {
	DBHost           string
	DBPort           string
	PostgresUser     string
	PostgresPassword string
	PostgresDB       string
	Secret           string
	Cost             int
	Port             string
}

var loadedConfig *Config

func Load() (*Config, error) {
	_ = godotenv.Load(".env")

	cfg := &Config{
		DBHost:           strings.TrimSpace(os.Getenv("DB_HOST")),
		DBPort:           strings.TrimSpace(os.Getenv("DB_PORT")),
		PostgresUser:     strings.TrimSpace(os.Getenv("POSTGRES_USER")),
		PostgresPassword: strings.TrimSpace(os.Getenv("POSTGRES_PASSWORD")),
		PostgresDB:       strings.TrimSpace(os.Getenv("POSTGRES_DB")),
		Secret:           strings.TrimSpace(os.Getenv("SECRET")),
		Port:             strings.TrimSpace(os.Getenv("PORT")),
	}

	var errs []string
	costRaw := strings.TrimSpace(os.Getenv("COST"))
	required := []struct {
		key   string
		value string
	}{
		{key: "DB_HOST", value: cfg.DBHost},
		{key: "DB_PORT", value: cfg.DBPort},
		{key: "POSTGRES_USER", value: cfg.PostgresUser},
		{key: "POSTGRES_PASSWORD", value: cfg.PostgresPassword},
		{key: "POSTGRES_DB", value: cfg.PostgresDB},
		{key: "SECRET", value: cfg.Secret},
		{key: "COST", value: costRaw},
		{key: "PORT", value: cfg.Port},
	}

	for _, field := range required {
		if field.value == "" {
			errs = append(errs, fmt.Sprintf("%s is required", field.key))
		}
	}

	if costRaw != "" {
		cost, err := strconv.Atoi(costRaw)
		if err != nil {
			errs = append(errs, "COST must be a valid integer")
		} else {
			cfg.Cost = cost
		}
	}

	if cfg.DBPort != "" {
		if _, err := strconv.Atoi(cfg.DBPort); err != nil {
			errs = append(errs, "DB_PORT must be a valid integer")
		}
	}

	if cfg.Port != "" {
		if _, err := strconv.Atoi(cfg.Port); err != nil {
			errs = append(errs, "PORT must be a valid integer")
		}
	}

	if len(errs) > 0 {
		return nil, fmt.Errorf("invalid configuration: %s", strings.Join(errs, "; "))
	}

	loadedConfig = cfg
	return loadedConfig, nil
}

func MustGet() *Config {
	if loadedConfig == nil {
		cfg, err := Load()
		if err != nil {
			panic(err)
		}
		loadedConfig = cfg
	}

	return loadedConfig
}
