package main

import (
	"context"
	"laughing-goggles/config"
	"laughing-goggles/handlers"
	"log/slog"
	"net/http"
	"time"
)

// TODO: handle SIGTERM, SIGKILL
func main() {
	ctx := context.Background()
	cfg := config.Init()

	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting ..")

	handler := handlers.NewHandler(logr)
	server := &http.Server{
		Addr:              cfg.Address,
		Handler:           handler,
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := server.ListenAndServe(); err != nil {
		logr.ErrorContext(ctx, "server error", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "stopping ..")
}
