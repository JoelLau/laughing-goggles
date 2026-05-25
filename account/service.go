package account

import (
	"errors"

	"github.com/shopspring/decimal"
)

type AccountService struct {
	// key: id, value: accountsByID object
	accountsByID map[int64]Account
}

func NewAccountService() *AccountService {
	return &AccountService{
		accountsByID: make(map[int64]Account),
	}
}

var (
	ErrAccountAlreadyExists    = errors.New("account already exists")
	ErrAccountNotFound         = errors.New("account not found")
	ErrInitialBalanceNotPositive = errors.New("initial balance must be positive")
)

func (s *AccountService) CreateAccount(params CreateAccountParams) (Account, error) {
	if !params.InitialBalance.IsPositive() {
		return Account{}, ErrInitialBalanceNotPositive
	}

	if _, ok := s.accountsByID[params.AccountID]; ok {
		return Account{}, ErrAccountAlreadyExists
	}

	s.accountsByID[params.AccountID] = Account{
		ID:      params.AccountID,
		Balance: params.InitialBalance,
	}

	return s.accountsByID[params.AccountID], nil

}

func (s *AccountService) GetAccountByID(id int64) (Account, error) {
	a, ok := s.accountsByID[id]
	if !ok {
		return Account{}, ErrAccountNotFound
	}
	return a, nil
}

type CreateAccountParams struct {
	AccountID      int64
	InitialBalance decimal.Decimal
}
