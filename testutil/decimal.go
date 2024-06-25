package testutil

import "github.com/shopspring/decimal"

func ToDecimal(ds string) decimal.Decimal {
	d, _ := decimal.NewFromString(ds)
	return d
}

func ToDecimalPtr(ds string) *decimal.Decimal {
	dp := ToDecimal(ds)
	return &dp
}

func DecimalPtr(d decimal.Decimal) *decimal.Decimal {
	return &d
}
