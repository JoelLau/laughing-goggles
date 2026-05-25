package httpapi

import (
	"laughing-goggles/account"
	"laughing-goggles/gen/api"
)

var _ api.StrictServerInterface = &Server{}

func NewServer(svc AccountsService) *Server {
	return &Server{
		accSvc: svc,
	}
}

type AccountsService interface {
	CreateAccount(account.CreateAccountParams) (account.Account, error)
	GetAccountByID(id int64) (account.Account, error)
}

type Server struct {
	accSvc AccountsService
}
