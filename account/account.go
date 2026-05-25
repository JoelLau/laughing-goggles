package account

import "github.com/shopspring/decimal"

type Account struct {
	ID      int64
	Balance decimal.Decimal
}
