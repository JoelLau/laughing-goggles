package account

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"laughing-goggles/gen/sqlc"

	"github.com/jackc/pgerrcode"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/shopspring/decimal"
)

type AccountService struct {
	pool    *pgxpool.Pool
	queries *sqlc.Queries
}

func NewAccountService(pool *pgxpool.Pool) *AccountService {
	return &AccountService{
		pool:    pool,
		queries: sqlc.New(pool),
	}
}

var (
	ErrAccountAlreadyExists      = errors.New("account already exists")
	ErrAccountNotFound           = errors.New("account not found")
	ErrInitialBalanceNotPositive = errors.New("initial balance must be positive")
	ErrAmountNotPositive         = errors.New("amount must be positive")
	ErrInsufficientBalance       = errors.New("insufficient balance")
)

var microsPerUnit = decimal.NewFromInt(1_000_000)

func fromMicros(micros int64) decimal.Decimal {
	return decimal.NewFromInt(micros).Div(microsPerUnit)
}

func toMicros(d decimal.Decimal) int64 {
	return d.Mul(microsPerUnit).BigInt().Int64()
}

func (s *AccountService) Ping(ctx context.Context) error {
	return s.pool.Ping(ctx)
}

func (s *AccountService) CreateAccount(ctx context.Context, params CreateAccountParams) (Account, error) {
	if !params.InitialBalance.IsPositive() {
		return Account{}, ErrInitialBalanceNotPositive
	}

	eventBody, err := json.Marshal(params)
	if err != nil {
		return Account{}, err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return Account{}, err
	}
	defer tx.Rollback(ctx)

	qtx := s.queries.WithTx(tx)

	event, err := qtx.CreateEvent(ctx, sqlc.CreateEventParams{
		Type: "CreateAccount",
		Data: []byte(eventBody),
	})
	if err != nil {
		return Account{}, err
	}

	accountID, err := qtx.CreateAccount(ctx, params.AccountID)
	if err != nil {
		var pgErr *pgconn.PgError
		if errors.As(err, &pgErr) && pgErr.Code == pgerrcode.UniqueViolation {
			return Account{}, ErrAccountAlreadyExists
		}

		return Account{}, err
	}

	_, err = qtx.CreateLedgerEntry(ctx, sqlc.CreateLedgerEntryParams{
		EventID:     event.ID,
		AccountID:   accountID,
		AmountMicro: toMicros(params.InitialBalance),
	})
	if err != nil {
		return Account{}, err
	}

	err = tx.Commit(ctx)
	if err != nil {
		return Account{}, err
	}

	return Account{
		ID:      accountID,
		Balance: params.InitialBalance,
	}, nil
}

type CreateAccountParams struct {
	AccountID      int64
	InitialBalance decimal.Decimal
}

func (s *AccountService) GetAccountByID(ctx context.Context, accountID int64) (Account, error) {
	row, err := s.queries.GetAccountByID(ctx, accountID)
	if errors.Is(err, pgx.ErrNoRows) {
		return Account{}, ErrAccountNotFound
	}
	if err != nil {
		return Account{}, fmt.Errorf("error fetching account by id: %w", err)
	}

	return Account{
		ID:      row.AccountID,
		Balance: fromMicros(row.BalanceMicros),
	}, nil
}

func (s *AccountService) CreateTransaction(ctx context.Context, params CreateTransactionParams) error {
	if !params.Amount.IsPositive() {
		return ErrAmountNotPositive
	}

	eventBody, err := json.Marshal(params)
	if err != nil {
		return err
	}

	tx, err := s.pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qtx := s.queries.WithTx(tx)

	event, err := qtx.CreateEvent(ctx, sqlc.CreateEventParams{
		Type: "CreateTransaction",
		Data: []byte(eventBody),
	})
	if err != nil {
		return err
	}

	sourceAccount, err := qtx.GetAccountByID(ctx, params.SourceAccountID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	destinationAccount, err := qtx.GetAccountByID(ctx, params.DestinationAccountID)
	if errors.Is(err, pgx.ErrNoRows) {
		return ErrAccountNotFound
	}
	if err != nil {
		return err
	}

	if fromMicros(sourceAccount.BalanceMicros).LessThan(params.Amount) {
		return ErrInsufficientBalance
	}

	_, err = qtx.CreateLedgerEntry(ctx, sqlc.CreateLedgerEntryParams{
		EventID:     event.ID,
		AccountID:   sourceAccount.AccountID,
		AmountMicro: -toMicros(params.Amount),
	})
	if err != nil {
		return err
	}

	_, err = qtx.CreateLedgerEntry(ctx, sqlc.CreateLedgerEntryParams{
		EventID:     event.ID,
		AccountID:   destinationAccount.AccountID,
		AmountMicro: toMicros(params.Amount),
	})
	if err != nil {
		return err
	}

	if err := tx.Commit(ctx); err != nil {
		return err
	}

	return nil
}

type CreateTransactionParams struct {
	SourceAccountID      int64
	DestinationAccountID int64
	Amount               decimal.Decimal
}
