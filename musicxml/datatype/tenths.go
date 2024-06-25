package datatype

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Tenths struct {
	Val *decimal.Decimal
	fmt.Stringer
}

func (t *Tenths) String() string {
	if t == nil {
		return ""
	}
	if t.Val == nil {
		return ""
	}
	return t.Val.String()
}

func (t *Tenths) StringPtr() *string {
	if t == nil {
		return nil
	}
	if t.Val == nil {
		return nil
	}
	s := t.Val.String()
	return &s
}

func (t *Tenths) Value() *decimal.Decimal {
	if t == nil {
		return nil
	}
	return t.Val
}

func NewTenthsFromString(value string) (*Tenths, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}
	d, err := decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}
	return &Tenths{Val: &d}, nil
}

func NewTenthsFixedPoint(value int64, exp int32) *Tenths {
	d := decimal.New(value, exp)
	return &Tenths{Val: &d}
}

func NewTenthsFromInt(value int64) *Tenths {
	d := decimal.NewFromInt(value)
	return &Tenths{Val: &d}
}

func NewTenthsFromUint64(value uint64) *Tenths {
	d := decimal.NewFromUint64(value)
	return &Tenths{Val: &d}
}

func NewTenthsFromFloat(value float64) *Tenths {
	d := decimal.NewFromFloat(value)
	return &Tenths{Val: &d}
}
