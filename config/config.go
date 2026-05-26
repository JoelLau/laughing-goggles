package config

import (
	"fmt"
	"log/slog"
	"net/url"

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
	Address   string `env:"ADDRESS"`
	DebugMode bool   `env:"DEBUG"`

	DBHost     string `env:"POSTGRES_HOST"`
	DBPort     string `env:"POSTGRES_PORT"`
	DBUser     string `env:"POSTGRES_USER"`
	DBPassword string `env:"POSTGRES_PASSWORD"`
	DBName     string `env:"POSTGRES_DB"`
}

func (c *Config) PostgresDSN() string {
	u := &url.URL{
		Scheme:   "postgres",
		User:     url.UserPassword(c.DBUser, c.DBPassword),
		Host:     c.DBHost + ":" + c.DBPort,
		Path:     "/" + c.DBName,
		RawQuery: "sslmode=disable",
	}
	return u.String()
}

func (c *Config) Logger() *slog.Logger {
	return NewLogger(c.DebugMode)
}
