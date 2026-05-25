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
	CreateAccount(CreateAccountParams) (Account, error)
	GetAccountByID(id int64) (Account, error)
}

type CreateAccountParams = account.CreateAccountParams

type Account = account.Account

type Server struct {
	accSvc AccountsService
}
