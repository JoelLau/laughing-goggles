package main

import (
	"context"
	"laughing-goggles/config"
	"laughing-goggles/db/sqlc"
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

	pool, err := pgxpool.New(ctx, "")
	if err != nil {
		panic(err)
	}
	defer pool.Close()

	srv := &http.Server{
		Addr:              cfg.Address,
		Handler:           httpapi.NewHandler(logr, sqlc.New(pool)),
		ReadHeaderTimeout: 10 * time.Second,
	}

	if err := srv.ListenAndServe(); err != nil {
		logr.ErrorContext(ctx, "server error", slog.Any("error", err))
	}

	logr.InfoContext(ctx, "stopping ..")
}
