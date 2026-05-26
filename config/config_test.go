package config

import (
	"testing"
)

func TestConnectionString(t *testing.T) {
	cfg := Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "user",
		DBPassword: "pass",
		DBName:     "mydb",
	}

	want := "postgres://user:pass@localhost:5432/mydb?sslmode=disable"
	if got := cfg.PostgresDSN(); got != want {
		t.Errorf("got %q, want %q", got, want)
	}
}
