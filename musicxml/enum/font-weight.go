package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type FontWeightEnum string

var FontWeight = struct {
	Normal FontWeightEnum
	Bold   FontWeightEnum
	opts   map[string]map[string]interface{}
}{
	Normal: FontWeightEnum("normal"),
	Bold:   FontWeightEnum("bold"),
	opts: map[string]map[string]interface{}{
		"normal": {
			"ordinal": "0",
		},
		"bold": {
			"ordinal": "1",
		},
	},
}

func NewFontWeight(t string) (*FontWeightEnum, error) {
	switch t {
	case FontWeight.Normal.String():
		return &FontWeight.Normal, nil
	case FontWeight.Bold.String():
		return &FontWeight.Bold, nil
	}
	return nil, fmt.Errorf("can not convert to FontWeight. t=%s", t)
}

func AllFontWeightEnumValues() []FontWeightEnum {
	values := []FontWeightEnum{
		FontWeight.Normal,
		FontWeight.Bold,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *FontWeightEnum) Equals(obj FontWeightEnum) bool {
	return *e == obj
}

func (e *FontWeightEnum) In(objs ...FontWeightEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *FontWeightEnum) Ordinal() string {
	return FontWeight.opts[e.String()]["ordinal"].(string)
}

func (e *FontWeightEnum) String() string {
	return string(*e)
}

func (e *FontWeightEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
