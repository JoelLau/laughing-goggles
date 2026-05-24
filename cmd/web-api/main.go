package main

import (
	"laughing-goggles/handlers"
	"log/slog"
	"net/http"
)

func main() {
	slog.Info("starting ..")

	handler := handlers.NewHandler()
	server := &http.Server{
		Addr:    ":8080", // TODO: move this to env var
		Handler: handler,
	}

	if err := server.ListenAndServe(); err != nil {
		slog.Error("server error", slog.Any("error", err))
	}

	slog.Info("stopping ..")
}
