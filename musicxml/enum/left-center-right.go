package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type LeftCenterRightEnum string

var LeftCenterRight = struct {
	Left   LeftCenterRightEnum
	Center LeftCenterRightEnum
	Right  LeftCenterRightEnum
	opts   map[string]map[string]interface{}
}{
	Left:   LeftCenterRightEnum("left"),
	Center: LeftCenterRightEnum("center"),
	Right:  LeftCenterRightEnum("right"),
	opts: map[string]map[string]interface{}{
		"left": {
			"ordinal": "0",
		},
		"center": {
			"ordinal": "1",
		},
		"right": {
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

func ToLeftCenterRight(t string) (*LeftCenterRightEnum, error) {
	switch t {
	case LeftCenterRight.Left.String():
		return &LeftCenterRight.Left, nil
	case LeftCenterRight.Center.String():
		return &LeftCenterRight.Center, nil
	case LeftCenterRight.Right.String():
		return &LeftCenterRight.Right, nil
	}
	return nil, fmt.Errorf("can not convert to LeftCenterRight. t=%s", t)
}

func AllLeftCenterRightEnumValues() []LeftCenterRightEnum {
	values := []LeftCenterRightEnum{
		LeftCenterRight.Left,
		LeftCenterRight.Center,
		LeftCenterRight.Right,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *LeftCenterRightEnum) Equals(obj LeftCenterRightEnum) bool {
	return *e == obj
}

func (e *LeftCenterRightEnum) In(objs ...LeftCenterRightEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *LeftCenterRightEnum) Ordinal() string {
	return LeftCenterRight.opts[e.String()]["ordinal"].(string)
}

func (e *LeftCenterRightEnum) String() string {
	return string(*e)
}

func (e *LeftCenterRightEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
