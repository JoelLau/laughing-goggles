package account

import (
	"errors"
	"fmt"

	"github.com/shopspring/decimal"
)

type AccountService struct {
	accountsByID map[int64]Account
}

func NewAccountService() *AccountService {
	return &AccountService{
		accountsByID: make(map[int64]Account),
	}
}

var (
	ErrAccountAlreadyExists      = errors.New("account already exists")
	ErrAccountNotFound           = errors.New("account not found")
	ErrInitialBalanceNotPositive = errors.New("initial balance must be positive")
	ErrAmountNotPositive         = errors.New("amount must be positive")
	ErrInsufficientBalance       = errors.New("insufficient balance")
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

type CreateAccountParams struct {
	AccountID      int64
	InitialBalance decimal.Decimal
}

func (s *AccountService) GetAccountByID(id int64) (Account, error) {
	acc, ok := s.accountsByID[id]
	if !ok {
		return Account{}, ErrAccountNotFound
	}
	return acc, nil
}

func (s *AccountService) CreateTransaction(params CreateTransactionParams) error {
	if !params.Amount.IsPositive() {
		return ErrAmountNotPositive
	}

	source, sourceOK := s.accountsByID[params.SourceAccountID]
	if !sourceOK {
		return fmt.Errorf("%w: %d", ErrAccountNotFound, params.SourceAccountID)
	}

	if source.Balance.LessThan(params.Amount) {
		return ErrInsufficientBalance
	}

	destination, destinationOK := s.accountsByID[params.DestinationAccountID]
	if !destinationOK {
		return fmt.Errorf("%w: %d", ErrAccountNotFound, params.DestinationAccountID)
	}

	source.Balance = source.Balance.Sub(params.Amount)
	destination.Balance = destination.Balance.Add(params.Amount)

	s.accountsByID[params.SourceAccountID] = source
	s.accountsByID[params.DestinationAccountID] = destination

	return nil
}

type CreateTransactionParams struct {
	SourceAccountID      int64
	DestinationAccountID int64
	Amount               decimal.Decimal
}
