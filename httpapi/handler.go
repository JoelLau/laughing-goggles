package httpapi

import (
	"context"
	"laughing-goggles/gen/api"
)

var _ api.StrictServerInterface = &Server{}

func NewServer() *Server {
	return &Server{}
}

type Server struct{}

// (GET /livez)
func (s *Server) Livez(ctx context.Context, request api.LivezRequestObject) (api.LivezResponseObject, error) {
	return api.Livez200JSONResponse{Data: "ok"}, nil
}

// (GET /readyz)
//
// TODO: ping live database
func (s *Server) Readyz(ctx context.Context, request api.ReadyzRequestObject) (api.ReadyzResponseObject, error) {
	return api.Readyz200JSONResponse{Data: "ok"}, nil
}

// (POST /accounts)
func (s *Server) CreateAccount(ctx context.Context, request api.CreateAccountRequestObject) (api.CreateAccountResponseObject, error) {
	return api.CreateAccount201Response{}, nil
}
