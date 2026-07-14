package config

import (
	"os"
	"time"
)

type Config struct {
	Addr            string
	ShutdownTimeout time.Duration
}

func FromEnv() Config {
	addr := os.Getenv("API_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	shutdownTimeout := 10 * time.Second
	if raw := os.Getenv("API_SHUTDOWN_TIMEOUT"); raw != "" {
		if parsed, err := time.ParseDuration(raw); err == nil {
			shutdownTimeout = parsed
		}
	}

	return Config{
		Addr:            addr,
		ShutdownTimeout: shutdownTimeout,
	}
}
