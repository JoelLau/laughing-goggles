package httpapi

import (
	"context"
	"errors"
	"fmt"
	"laughing-goggles/account"
	"laughing-goggles/gen/api"
	"net/http"

	"github.com/shopspring/decimal"
)

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
	initialBalance, err := decimal.NewFromString(request.Body.InitialBalance)
	if err != nil {
		return api.CreateAccount400JSONResponse{
			Type:   "https://github.com/JoelLau/laughing-goggles/errors/invalid-initial-balance",
			Title:  "Bad Request",
			Status: http.StatusBadRequest,
			Detail: new(fmt.Sprintf("initial_balance must be a numeric string (e.g. \"10.00\"), got %q", request.Body.InitialBalance)),
		}, nil
	}

	acc, err := s.accSvc.CreateAccount(CreateAccountParams{
		AccountID:      request.Body.AccountId,
		InitialBalance: initialBalance,
	})
	if errors.Is(err, account.ErrInitialBalanceNotPositive) {
		return api.CreateAccount400JSONResponse{
			Type:   "https://github.com/JoelLau/laughing-goggles/errors/negative-initial-balance",
			Title:  "Bad Request",
			Status: http.StatusBadRequest,
			Detail: new("initial_balance must be a positive numeric string (e.g. \"10.00\")"),
		}, nil
	}
	if errors.Is(err, account.ErrAccountAlreadyExists) {
		return api.CreateAccount409JSONResponse{
			Type:     "https://github.com/JoelLau/laughing-goggles/errors/account-already-exists",
			Title:    "Conflict",
			Status:   http.StatusConflict,
			Detail:   new(fmt.Sprintf("account with id %d already exists", request.Body.AccountId)),
			Instance: new(fmt.Sprintf("/accounts/%d", request.Body.AccountId)),
		}, nil
	}
	if err != nil {
		return api.CreateAccount500JSONResponse{
			Type:     "https://github.com/JoelLau/laughing-goggles/errors/internal-server-error",
			Title:    "Internal Server Error",
			Status:   http.StatusInternalServerError,
			Detail:   new("failed to create account"),
			Instance: new("/accounts"),
		}, nil
	}

	return api.CreateAccount201JSONResponse{
		AccountId: acc.ID,
		Balance:   acc.Balance.String(),
	}, nil
}

// (GET /accounts/{account_id})
func (s *Server) GetAccountByID(ctx context.Context, request api.GetAccountByIDRequestObject) (api.GetAccountByIDResponseObject, error) {
	acc, err := s.accSvc.GetAccountByID(request.AccountId)
	if errors.Is(err, account.ErrAccountNotFound) {
		return api.GetAccountByID404JSONResponse{
			Type:   "https://github.com/JoelLau/laughing-goggles/errors/account-not-found",
			Title:  "Not Found",
			Status: http.StatusNotFound,
			Detail: new(fmt.Sprintf("no account with id %d", request.AccountId)),
		}, nil
	}
	if err != nil {
		return api.GetAccountByID500JSONResponse{
			Type:   "https://github.com/JoelLau/laughing-goggles/errors/internal-server-error",
			Title:  "Internal Server Error",
			Status: http.StatusInternalServerError,
			Detail: new("failed to get account"),
		}, nil
	}

	return api.GetAccountByID200JSONResponse{
		AccountId: acc.ID,
		Balance:   acc.Balance.String(),
	}, nil
}
