package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type CssFontSizeEnum string

var CssFontSize = struct {
	XXSmall CssFontSizeEnum
	XSmall  CssFontSizeEnum
	Small   CssFontSizeEnum
	Medium  CssFontSizeEnum
	Large   CssFontSizeEnum
	XLarge  CssFontSizeEnum
	XXLarge CssFontSizeEnum
	opts    map[string]map[string]interface{}
}{
	XXSmall: CssFontSizeEnum("xx-small"),
	XSmall:  CssFontSizeEnum("x-small"),
	Small:   CssFontSizeEnum("small"),
	Medium:  CssFontSizeEnum("medium"),
	Large:   CssFontSizeEnum("large"),
	XLarge:  CssFontSizeEnum("x-large"),
	XXLarge: CssFontSizeEnum("xx-large"),
	opts: map[string]map[string]interface{}{
		"xx-small": {
			"ordinal": "0",
		},
		"x-small": {
			"ordinal": "1",
		},
		"small": {
			"ordinal": "2",
		},
		"medium": {
			"ordinal": "3",
		},
		"large": {
			"ordinal": "4",
		},
		"x-large": {
			"ordinal": "5",
		},
		"xx-large": {
			"ordinal": "6",
		},
	},
}

func ToCssFontSize(t string) (*CssFontSizeEnum, error) {
	switch t {
	case CssFontSize.XXSmall.String():
		return &CssFontSize.XXSmall, nil
	case CssFontSize.XSmall.String():
		return &CssFontSize.XSmall, nil
	case CssFontSize.Small.String():
		return &CssFontSize.Small, nil
	case CssFontSize.Medium.String():
		return &CssFontSize.Medium, nil
	case CssFontSize.Large.String():
		return &CssFontSize.Large, nil
	case CssFontSize.XLarge.String():
		return &CssFontSize.XLarge, nil
	case CssFontSize.XXLarge.String():
		return &CssFontSize.XXLarge, nil
	}
	return nil, fmt.Errorf("can not convert to CssFontSize. t=%s", t)
}

func AllCssFontSizeEnumValues() []CssFontSizeEnum {
	values := []CssFontSizeEnum{
		CssFontSize.XXSmall,
		CssFontSize.XSmall,
		CssFontSize.Small,
		CssFontSize.Medium,
		CssFontSize.Large,
		CssFontSize.XLarge,
		CssFontSize.XXLarge,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *CssFontSizeEnum) Equals(obj CssFontSizeEnum) bool {
	return *e == obj
}

func (e *CssFontSizeEnum) In(objs ...CssFontSizeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *CssFontSizeEnum) Ordinal() string {
	return CssFontSize.opts[e.String()]["ordinal"].(string)
}

func (e *CssFontSizeEnum) String() string {
	return string(*e)
}

func (e *CssFontSizeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
