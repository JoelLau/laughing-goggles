package main

import (
	"context"
	"errors"
	"laughing-goggles/account"
	"laughing-goggles/config"
	"laughing-goggles/httpapi"
	"log/slog"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), os.Interrupt, syscall.SIGTERM)
	defer stop()

	cfg := config.Init()
	logr := cfg.Logger()
	logr.InfoContext(ctx, "starting..")

	pool, err := pgxpool.New(ctx, cfg.PostgresDSN())
	if err != nil {
		logr.ErrorContext(ctx, "failed to create db pool", slog.Any("error", err))
		return
	}
	defer pool.Close()

	srv := &http.Server{
		Handler:     httpapi.NewHandler(logr, account.NewAccountService(pool)),
		Addr:        cfg.Address,
		ReadTimeout: 10 * time.Second,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && !errors.Is(err, http.ErrServerClosed) {
			logr.ErrorContext(ctx, "unexpected server shutdown", slog.Any("error", err))
		}
	}()

	<-ctx.Done()
	stop()

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	if err := srv.Shutdown(shutdownCtx); err != nil {
		logr.ErrorContext(shutdownCtx, "server shutdown error", slog.Any("error", err))
	}

	logr.InfoContext(shutdownCtx, "exiting..")
}
