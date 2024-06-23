package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type YesNoEnum string

var YesNo = struct {
	Yes  YesNoEnum
	No   YesNoEnum
	opts map[string]map[string]interface{}
}{
	Yes: YesNoEnum("yes"),
	No:  YesNoEnum("no"),
	opts: map[string]map[string]interface{}{
		"yes": {
			"ordinal": "0",
		},
		"no": {
			"ordinal": "1",
		},
	},
}

func ToYesNo(t string, en string, at string) (*YesNoEnum, error) {
	switch t {
	case YesNo.Yes.String():
		return &YesNo.Yes, nil
	case YesNo.No.String():
		return &YesNo.No, nil
	}
	return nil, fmt.Errorf("invalid <%s/> attr. attr => %s yesno=%s", en, at, t)
}

func AllYesNoEnumValues() []YesNoEnum {
	values := []YesNoEnum{
		YesNo.Yes,
		YesNo.No,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *YesNoEnum) Equals(obj YesNoEnum) bool {
	return *e == obj
}

func (e *YesNoEnum) In(objs ...YesNoEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *YesNoEnum) Ordinal() string {
	return YesNo.opts[e.String()]["ordinal"].(string)
}

func (e *YesNoEnum) String() string {
	return string(*e)
}
