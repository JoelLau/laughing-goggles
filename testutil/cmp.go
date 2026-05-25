package testutil

import (
	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
)

var DefaultCmpOpts = cmp.Options{
	DecimalComparerOpt,
}

var DecimalComparerOpt = cmp.Comparer(
	func(a, b decimal.Decimal) bool { return a.Equal(b) },
)
