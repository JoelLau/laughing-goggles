package httpapi

import (
	"context"
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
	Ping(ctx context.Context) error

	CreateAccount(ctx context.Context, params CreateAccountParams) (Account, error)
	GetAccountByID(ctx context.Context, accountID int64) (Account, error)
	CreateTransaction(ctx context.Context, params CreateTransactionParams) error
}

type CreateAccountParams = account.CreateAccountParams
type CreateTransactionParams = account.CreateTransactionParams

type Account = account.Account

type Server struct {
	accSvc AccountsService
}
