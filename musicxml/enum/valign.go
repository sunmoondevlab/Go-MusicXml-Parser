package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type ValignEnum string

var Valign = struct {
	Top      ValignEnum
	Middle   ValignEnum
	Bottom   ValignEnum
	Baseline ValignEnum
	opts     map[string]map[string]interface{}
}{
	Top:      ValignEnum("top"),
	Middle:   ValignEnum("middle"),
	Bottom:   ValignEnum("bottom"),
	Baseline: ValignEnum("baseline"),
	opts: map[string]map[string]interface{}{
		"top": {
			"ordinal": "0",
		},
		"middle": {
			"ordinal": "1",
		},
		"bottom": {
			"ordinal": "2",
		},
		"baseline": {
			"ordinal": "3",
		},
	},
}

func ToValign(t string) (*ValignEnum, error) {
	switch t {
	case Valign.Top.String():
		return &Valign.Top, nil
	case Valign.Middle.String():
		return &Valign.Middle, nil
	case Valign.Bottom.String():
		return &Valign.Bottom, nil
	case Valign.Baseline.String():
		return &Valign.Baseline, nil
	}
	return nil, fmt.Errorf("can not convert to Valign. t=%s", t)
}

func AllValignEnumValues() []ValignEnum {
	values := []ValignEnum{
		Valign.Top,
		Valign.Middle,
		Valign.Bottom,
		Valign.Baseline,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *ValignEnum) Equals(obj ValignEnum) bool {
	return *e == obj
}

func (e *ValignEnum) In(objs ...ValignEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *ValignEnum) Ordinal() string {
	return Valign.opts[e.String()]["ordinal"].(string)
}

func (e *ValignEnum) String() string {
	return string(*e)
}

func (e *ValignEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
