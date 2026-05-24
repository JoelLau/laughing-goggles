package handlers

import (
	"encoding/json"
	"laughing-goggles/gen/server"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

func NewHandler(logr *slog.Logger) http.Handler {
	serverImpl := NewServer()
	strictHandler := server.NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()

	r.Use(slogchi.New(logr))
	r.Use(middleware.Recoverer)

	server.HandlerWithOptions(strictHandler, server.ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: errorHandlerFunc,
	})

	return r
}

func errorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(server.ErrorResponse{
		Type:   "https://github.com/JoelLau/laughing-goggles/errors/invalid-params",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: new(err.Error()),
	})
}
