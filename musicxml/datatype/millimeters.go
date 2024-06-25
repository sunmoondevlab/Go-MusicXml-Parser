package datatype

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
)

type Millimeters struct {
	Val *decimal.Decimal
	fmt.Stringer
}

func (m *Millimeters) String() string {
	if m == nil {
		return ""
	}
	if m.Val == nil {
		return ""
	}
	return m.Val.String()
}

func (m *Millimeters) StringPtr() *string {
	if m == nil {
		return nil
	}
	if m.Val == nil {
		return nil
	}
	s := m.Val.String()
	return &s
}

func (m *Millimeters) Value() *decimal.Decimal {
	if m == nil {
		return nil
	}
	return m.Val
}

func NewMillimetersFromString(value string) (*Millimeters, error) {
	if strings.TrimSpace(value) == "" {
		return nil, nil
	}
	d, err := decimal.NewFromString(value)
	if err != nil {
		return nil, err
	}
	return &Millimeters{Val: &d}, nil
}

func NewMillimetersFixedPoint(value int64, exp int32) *Millimeters {
	d := decimal.New(value, exp)
	return &Millimeters{Val: &d}
}

func NewMillimetersFromInt(value int64) *Millimeters {
	d := decimal.NewFromInt(value)
	return &Millimeters{Val: &d}
}

func NewMillimetersFromUint64(value uint64) *Millimeters {
	d := decimal.NewFromUint64(value)
	return &Millimeters{Val: &d}
}

func NewMillimetersFromFloat(value float64) *Millimeters {
	d := decimal.NewFromFloat(value)
	return &Millimeters{Val: &d}
}
