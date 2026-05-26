package main

import (
	"context"
	"laughing-goggles/account"
	"laughing-goggles/config"
	"laughing-goggles/httpapi"
	"log/slog"
	"net/http"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

// TODO: handle SIGTERM, SIGKILL
func main() {
	ctx := context.Background()
	cfg := config.Init()

	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting ..")

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		logr.ErrorContext(ctx, "failed to connect to database", slog.Any("error", err))
		return
	}

	svc := account.NewAccountService(pool)
	handler := httpapi.NewHandler(logr, svc)
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
