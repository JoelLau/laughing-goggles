package httpapi

import (
	"context"
	"laughing-goggles/api"
)

var _ api.StrictServerInterface = &Server{}

type Server struct {
	store AccountStore
}

type AccountStore interface {
	Ping(context.Context) error
}

func NewServer(store AccountStore) *Server {
	return &Server{store: store}
}
