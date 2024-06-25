package enum

import (
	"fmt"
)

type XlinkTypeEnum string

var XlinkType = struct {
	Simple XlinkTypeEnum
	opts   map[string]map[string]interface{}
}{
	Simple: XlinkTypeEnum("simple"),
	opts: map[string]map[string]interface{}{
		"simple": {
			"ordinal": "0",
		},
	},
}

func ToXlinkType(t string, en string) (*XlinkTypeEnum, error) {
	switch t {
	case XlinkType.Simple.String():
		return &XlinkType.Simple, nil
	}
	return nil, fmt.Errorf("invalid <%s/> attr. attr => xlink:type type=%s", en, t)
}

func AllXlinkTypeEnumValues() []XlinkTypeEnum {
	values := []XlinkTypeEnum{
		XlinkType.Simple,
	}

	// sort.Slice(values, func(i, j int) bool {
	// 	ordinalI, _ := strconv.Atoi(values[i].Ordinal())
	// 	ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
	// 	return ordinalI < ordinalJ
	// })

	return values
}

func (e *XlinkTypeEnum) Equals(obj XlinkTypeEnum) bool {
	return *e == obj
}

func (e *XlinkTypeEnum) In(objs ...XlinkTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *XlinkTypeEnum) Ordinal() string {
	return XlinkType.opts[e.String()]["ordinal"].(string)
}

func (e *XlinkTypeEnum) String() string {
	return string(*e)
}

func (e *XlinkTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
