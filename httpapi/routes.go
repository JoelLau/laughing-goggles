package httpapi

import (
	_ "embed"
	"encoding/json"
	"laughing-goggles/gen/api"
	"log/slog"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	slogchi "github.com/samber/slog-chi"
)

//go:embed openapi.yaml
var openapiSpec []byte

func NewHandler(logr *slog.Logger, svc AccountsService) http.Handler {
	serverImpl := NewServer(svc)
	strictHandler := api.NewStrictHandler(serverImpl, nil)

	r := chi.NewRouter()

	r.Use(slogchi.New(logr))
	r.Use(middleware.Recoverer)

	r.Get("/openapi.yaml", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/yaml")
		w.Write(openapiSpec)
	})

	r.Get("/swagger", func(w http.ResponseWriter, r *http.Request) {
		http.Redirect(w, r, "/swagger/", http.StatusMovedPermanently)
	})

	r.Get("/swagger/", func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/html")
		w.Write([]byte(`<!doctype html><html><head><title>Transfers Service API</title><meta charset="utf-8"/><link rel="stylesheet" href="https://unpkg.com/swagger-ui-dist/swagger-ui.css"/></head><body><div id="swagger-ui"></div><script src="https://unpkg.com/swagger-ui-dist/swagger-ui-bundle.js"></script><script>SwaggerUIBundle({url:"/openapi.yaml",dom_id:"#swagger-ui"})</script></body></html>`))
	})

	api.HandlerWithOptions(strictHandler, api.ChiServerOptions{
		BaseRouter:       r,
		ErrorHandlerFunc: errorHandlerFunc,
	})

	return r
}

func errorHandlerFunc(w http.ResponseWriter, r *http.Request, err error) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusBadRequest)
	_ = json.NewEncoder(w).Encode(api.ErrorResponse{
		Type:   "https://github.com/JoelLau/laughing-goggles/errors/invalid-params",
		Title:  "Bad Request",
		Status: http.StatusBadRequest,
		Detail: new(err.Error()),
	})
}
