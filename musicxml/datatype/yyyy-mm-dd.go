package datatype

import (
	"fmt"
	"strings"
	"time"
)

type YyyyMmDd struct {
	Val *time.Time
	fmt.Stringer
}

func (ymd *YyyyMmDd) String() string {
	if ymd == nil {
		return ""
	}
	df := "2006-01-02"
	var ds string
	if ymd.Val != nil {
		dsv := ymd.Val.Format(df)
		ds = dsv
	}
	return ds
}

func (ymd *YyyyMmDd) StringPtr() *string {
	if ymd == nil {
		return nil
	}
	if ymd.Val == nil {
		return nil
	}
	ds := ymd.String()
	return &ds
}

func (ymd *YyyyMmDd) Value() *time.Time {
	if ymd == nil {
		return nil
	}
	return ymd.Val
}

func NewYyyyMmDd(value string) (*YyyyMmDd, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, nil
	}
	df := "2006-01-02"
	dtv, err := time.Parse(df, value)
	if err != nil {
		return nil, err
	}

	return &YyyyMmDd{
		Val: &dtv,
	}, nil
}
