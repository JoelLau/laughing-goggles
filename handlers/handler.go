package handlers

import (
	"fmt"
	"log/slog"
	"net/http"
)

func NewHandler(logr *slog.Logger) http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/api/livez", func(w http.ResponseWriter, r *http.Request) {
		logr.InfoContext(r.Context(), fmt.Sprintf("%s %s", r.Method, r.RequestURI))

		w.WriteHeader(http.StatusOK)
		_, _ = w.Write([]byte("{\"data\": \"ok\"}"))
	})

	return r
}
