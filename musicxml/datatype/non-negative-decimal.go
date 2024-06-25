package datatype

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type NonNegativeDecimal struct {
	Val *decimal.Decimal
	fmt.Stringer
}

func (nnd *NonNegativeDecimal) String() string {
	if nnd == nil {
		return ""
	}
	if nnd.Val == nil {
		return ""
	}
	return nnd.Val.String()
}

func (nnd *NonNegativeDecimal) StringPtr() *string {
	if nnd == nil {
		return nil
	}
	if nnd.Val == nil {
		return nil
	}
	ns := nnd.String()
	return &ns
}

func (nnd *NonNegativeDecimal) Value() *decimal.Decimal {
	if nnd == nil {
		return nil
	}
	return nnd.Val
}

func NewNonNegativeDecimalFromString(value string) (*NonNegativeDecimal, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}
	d, err := decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}
	if decimal.Zero.Compare(d) > 0 {
		return nil, fmt.Errorf("value must be greater than or equal to 0. value=%s", value)
	}
	return &NonNegativeDecimal{Val: &d}, nil
}

func NewNonNegativeDecimalFixedPoint(value int64, exp int32) *NonNegativeDecimal {
	d := decimal.New(value, exp)
	if decimal.Zero.Compare(d) > 0 {
		return nil
	}
	return &NonNegativeDecimal{Val: &d}
}

func NewNonNegativeDecimalFromInt(value int64) *NonNegativeDecimal {
	d := decimal.NewFromInt(value)
	if decimal.Zero.Compare(d) > 0 {
		return nil
	}
	return &NonNegativeDecimal{Val: &d}
}

func NewNonNegativeDecimalFromUint64(value uint64) *NonNegativeDecimal {
	d := decimal.NewFromUint64(value)
	return &NonNegativeDecimal{Val: &d}
}

func NewNonNegativeDecimalFromFloat(value float64) *NonNegativeDecimal {
	d := decimal.NewFromFloat(value)
	if decimal.Zero.Compare(d) > 0 {
		return nil
	}
	return &NonNegativeDecimal{Val: &d}
}
