package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type NoteSizeTypeEnum string

var NoteSizeType = struct {
	Cue      NoteSizeTypeEnum
	Grace    NoteSizeTypeEnum
	GraceCue NoteSizeTypeEnum
	Large    NoteSizeTypeEnum
	opts     map[string]map[string]interface{}
}{
	Cue:      NoteSizeTypeEnum("cue"),
	Grace:    NoteSizeTypeEnum("grace"),
	GraceCue: NoteSizeTypeEnum("grace-cue"),
	Large:    NoteSizeTypeEnum("large"),
	opts: map[string]map[string]interface{}{
		"cue": {
			"ordinal": "0",
		},
		"grace": {
			"ordinal": "1",
		},
		"grace-cue": {
			"ordinal": "2",
		},
		"large": {
			"ordinal": "3",
		},
	},
}

func ToNoteSizeType(t string) (*NoteSizeTypeEnum, error) {
	switch t {
	case NoteSizeType.Cue.String():
		return &NoteSizeType.Cue, nil
	case NoteSizeType.Grace.String():
		return &NoteSizeType.Grace, nil
	case NoteSizeType.GraceCue.String():
		return &NoteSizeType.GraceCue, nil
	case NoteSizeType.Large.String():
		return &NoteSizeType.Large, nil
	}
	return nil, fmt.Errorf("can not convert to NoteSizeType. t=%s", t)
}

func AllNoteSizeTypeEnumValues() []NoteSizeTypeEnum {
	values := []NoteSizeTypeEnum{
		NoteSizeType.Cue,
		NoteSizeType.Grace,
		NoteSizeType.GraceCue,
		NoteSizeType.Large,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *NoteSizeTypeEnum) Equals(obj NoteSizeTypeEnum) bool {
	return *e == obj
}

func (e *NoteSizeTypeEnum) In(objs ...NoteSizeTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *NoteSizeTypeEnum) Ordinal() string {
	return NoteSizeType.opts[e.String()]["ordinal"].(string)
}

func (e *NoteSizeTypeEnum) String() string {
	return string(*e)
}

func (e *NoteSizeTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
