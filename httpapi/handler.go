package httpapi

import (
	"context"
	"laughing-goggles/api"
	"net/http"
)

// (GET /livez)
func (s *Server) Livez(ctx context.Context, request api.LivezRequestObject) (api.LivezResponseObject, error) {
	return api.Livez200JSONResponse{Data: "ok"}, nil
}

// (GET /readyz)
func (s *Server) Readyz(ctx context.Context, request api.ReadyzRequestObject) (api.ReadyzResponseObject, error) {
	if err := s.store.Ping(ctx); err != nil {
		return api.Readyz500JSONResponse{
			Type:   "https://github.com/JoelLau/laughing-goggles/errors/database-unavailable",
			Title:  "Service Unavailable",
			Status: http.StatusInternalServerError,
			Detail: new("database is not reachable"),
		}, nil
	}

	return api.Readyz200JSONResponse{Data: "ok"}, nil
}

// (POST /accounts)
//
// TODO: implement
func (s *Server) CreateAccount(ctx context.Context, request api.CreateAccountRequestObject) (api.CreateAccountResponseObject, error) {
	return api.CreateAccount201Response{}, nil
}
