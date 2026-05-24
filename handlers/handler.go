package handlers

import (
	"net/http"
)

func NewHandler() http.Handler {
	r := http.NewServeMux()

	r.HandleFunc("/api/livez", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	return r
}
