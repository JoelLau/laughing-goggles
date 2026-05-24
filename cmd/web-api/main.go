package main

import (
	"context"
	"laughing-goggles/config"
	"laughing-goggles/handlers"
	"log/slog"
	"net/http"
)

func main() {
	ctx := context.Background()
	cfg := config.Init()

	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting ..")

	handler := handlers.NewHandler()
	server := &http.Server{
		Addr:    cfg.Addr,
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		logr.ErrorContext(ctx, "server error", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "stopping ..")
}
