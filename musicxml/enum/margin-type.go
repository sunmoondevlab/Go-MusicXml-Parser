package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type MarginTypeEnum string

var MarginType = struct {
	Both MarginTypeEnum
	Even MarginTypeEnum
	Odd  MarginTypeEnum
	opts map[string]map[string]interface{}
}{
	Both: MarginTypeEnum("both"),
	Even: MarginTypeEnum("even"),
	Odd:  MarginTypeEnum("odd"),
	opts: map[string]map[string]interface{}{
		"both": {
			"ordinal": "0",
		},
		"even": {
			"ordinal": "1",
		},
		"odd": {
			"ordinal": "2",
		},
	},
}

func ToMarginType(t string) (*MarginTypeEnum, error) {
	switch t {
	case MarginType.Both.String():
		return &MarginType.Both, nil
	case MarginType.Even.String():
		return &MarginType.Even, nil
	case MarginType.Odd.String():
		return &MarginType.Odd, nil
	}
	return nil, fmt.Errorf("can not convert to MarginType. t=%s", t)
}

func AllMarginTypeEnumValues() []MarginTypeEnum {
	values := []MarginTypeEnum{
		MarginType.Both,
		MarginType.Even,
		MarginType.Odd,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *MarginTypeEnum) Equals(obj MarginTypeEnum) bool {
	return *e == obj
}

func (e *MarginTypeEnum) In(objs ...MarginTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *MarginTypeEnum) Ordinal() string {
	return MarginType.opts[e.String()]["ordinal"].(string)
}

func (e *MarginTypeEnum) String() string {
	return string(*e)
}

func (e *MarginTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
