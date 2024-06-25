package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type XlinkActuateEnum string

var XlinkActuate = struct {
	OnRequest XlinkActuateEnum
	OnLoad    XlinkActuateEnum
	Other     XlinkActuateEnum
	None      XlinkActuateEnum
	opts      map[string]map[string]interface{}
}{
	OnRequest: XlinkActuateEnum("onRequest"),
	OnLoad:    XlinkActuateEnum("onLoad"),
	Other:     XlinkActuateEnum("other"),
	None:      XlinkActuateEnum("none"),
	opts: map[string]map[string]interface{}{
		"onRequest": {
			"ordinal": "0",
		},
		"onLoad": {
			"ordinal": "1",
		},
		"other": {
			"ordinal": "2",
		},
		"none": {
			"ordinal": "3",
		},
	},
}

func ToXlinkActuate(t string) (*XlinkActuateEnum, error) {
	switch t {
	case XlinkActuate.OnRequest.String():
		return &XlinkActuate.OnRequest, nil
	case XlinkActuate.OnLoad.String():
		return &XlinkActuate.OnLoad, nil
	case XlinkActuate.Other.String():
		return &XlinkActuate.Other, nil
	case XlinkActuate.None.String():
		return &XlinkActuate.None, nil
	}
	return nil, fmt.Errorf("can not convert to XlinkActuate. t=%s", t)
}

func AllXlinkActuateEnumValues() []XlinkActuateEnum {
	values := []XlinkActuateEnum{
		XlinkActuate.OnRequest,
		XlinkActuate.OnLoad,
		XlinkActuate.Other,
		XlinkActuate.None,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *XlinkActuateEnum) Equals(obj XlinkActuateEnum) bool {
	return *e == obj
}

func (e *XlinkActuateEnum) In(objs ...XlinkActuateEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *XlinkActuateEnum) Ordinal() string {
	return XlinkActuate.opts[e.String()]["ordinal"].(string)
}

func (e *XlinkActuateEnum) String() string {
	return string(*e)
}

func (e *XlinkActuateEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
