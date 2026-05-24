package config

import (
	"log/slog"
	"os"
)

func NewLogger(debug bool) *slog.Logger {
	level := slog.LevelInfo
	if debug {
		level = slog.LevelDebug
	}

	return slog.New(
		slog.NewJSONHandler(
			os.Stderr,
			&slog.HandlerOptions{Level: level, AddSource: true},
		),
	)
}
