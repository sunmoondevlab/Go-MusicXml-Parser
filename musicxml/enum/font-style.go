package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type FontStyleEnum string

var FontStyle = struct {
	Normal FontStyleEnum
	Italic FontStyleEnum
	opts   map[string]map[string]interface{}
}{
	Normal: FontStyleEnum("normal"),
	Italic: FontStyleEnum("italic"),
	opts: map[string]map[string]interface{}{
		"normal": {
			"ordinal": "0",
		},
		"italic": {
			"ordinal": "1",
		},
	},
}

func NewFontStyle(t string) (*FontStyleEnum, error) {
	switch t {
	case FontStyle.Normal.String():
		return &FontStyle.Normal, nil
	case FontStyle.Italic.String():
		return &FontStyle.Italic, nil
	}
	return nil, fmt.Errorf("can not convert to FontStyle. t=%s", t)
}

func AllFontStyleEnumValues() []FontStyleEnum {
	values := []FontStyleEnum{
		FontStyle.Normal,
		FontStyle.Italic,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *FontStyleEnum) Equals(obj FontStyleEnum) bool {
	return *e == obj
}

func (e *FontStyleEnum) In(objs ...FontStyleEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *FontStyleEnum) Ordinal() string {
	return FontStyle.opts[e.String()]["ordinal"].(string)
}

func (e *FontStyleEnum) String() string {
	return string(*e)
}

func (e *FontStyleEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
