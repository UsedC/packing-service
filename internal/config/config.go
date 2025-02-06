package config

import (
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"strings"
)

type Config struct {
	Address   string
	LogLevel  slog.Level
	PackSizes []int
}

func NewConfig() (*Config, error) {
	cfg := &Config{}

	address := os.Getenv("ADDRESS")
	if address == "" {
		address = ":8080"
	}
	cfg.Address = address

	logLevelStr := os.Getenv("LOG_LEVEL")
	if logLevelStr == "" {
		logLevelStr = "info"
	}
	switch strings.ToLower(logLevelStr) {
	case "debug":
		cfg.LogLevel = slog.LevelDebug
	case "info":
		cfg.LogLevel = slog.LevelInfo
	case "warn":
		cfg.LogLevel = slog.LevelWarn
	case "error":
		cfg.LogLevel = slog.LevelError
	default:
		return nil, fmt.Errorf("invalid LOG_LEVEL: %s", logLevelStr)
	}

	packSizesStr := os.Getenv("PACK_SIZES")
	if packSizesStr != "" {
		parts := strings.Split(packSizesStr, ",")
		for _, part := range parts {
			part = strings.TrimSpace(part)
			if part == "" {
				continue
			}
			size, err := strconv.Atoi(part)
			if err != nil {
				return nil, fmt.Errorf("invalid PACK_SIZES value %q: %w", part, err)
			}
			cfg.PackSizes = append(cfg.PackSizes, size)
		}
	} else {
		cfg.PackSizes = []int{250, 500, 1000, 2000, 5000}
	}

	return cfg, nil
}
