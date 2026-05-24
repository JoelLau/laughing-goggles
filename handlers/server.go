package handlers

import (
	"context"
	"laughing-goggles/gen/server"
)

var _ server.StrictServerInterface = &Server{}

func NewServer() *Server {
	return &Server{}
}

type Server struct{}

// (GET /api/livez)
func (s *Server) Livez(ctx context.Context, request server.LivezRequestObject) (server.LivezResponseObject, error) {
	return server.Livez200JSONResponse{Data: "ok"}, nil
}

// (GET /api/readyz)
func (s *Server) Readyz(ctx context.Context, request server.ReadyzRequestObject) (server.ReadyzResponseObject, error) {
	// TODO: ping live database
	return server.Readyz200JSONResponse{Data: "ok"}, nil
}

// (POST /accounts)
func (s *Server) CreateAccount(ctx context.Context, request server.CreateAccountRequestObject) (server.CreateAccountResponseObject, error) {
	return server.CreateAccount201JSONResponse{}, nil
}
