package account

import (
	"context"
	"laughing-goggles/db/sqlc"
)

type AccountRepository struct {
	queries *sqlc.Queries
}

func NewAccountRepository(q *sqlc.Queries) *AccountRepository {
	return &AccountRepository{
		queries: q,
	}
}

func (r *AccountRepository) Ping(ctx context.Context) error {
	return r.queries.Ping(ctx)
}
