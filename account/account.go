package account

import "github.com/shopspring/decimal"

type Account struct {
	ID      string
	Balance decimal.Decimal
}
