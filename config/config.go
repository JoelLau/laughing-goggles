package config

import (
	"fmt"
	"log/slog"

	"github.com/golobby/env/v2"
)

func Init() Config {
	c := Config{}
	if err := env.Feed(&c); err != nil {
		panic(fmt.Errorf("failed to feed env vars into Config struct: %w", err))
	}

	return c
}

type Config struct {
	Addr      string `env:"ADDR"`
	DebugMode bool   `env:"DEBUG"`
}

func (c *Config) Logger() *slog.Logger {
	return NewLogger(c.DebugMode)
}
