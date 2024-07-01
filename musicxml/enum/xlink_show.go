package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type XlinkShowEnum string

var XlinkShow = struct {
	New     XlinkShowEnum
	Replace XlinkShowEnum
	Embed   XlinkShowEnum
	Other   XlinkShowEnum
	None    XlinkShowEnum
	opts    map[string]map[string]interface{}
}{
	New:     XlinkShowEnum("new"),
	Replace: XlinkShowEnum("replace"),
	Embed:   XlinkShowEnum("embed"),
	Other:   XlinkShowEnum("other"),
	None:    XlinkShowEnum("none"),
	opts: map[string]map[string]interface{}{
		"new": {
			"ordinal": "0",
		},
		"replace": {
			"ordinal": "1",
		},
		"embed": {
			"ordinal": "2",
		},
		"other": {
			"ordinal": "3",
		},
		"none": {
			"ordinal": "4",
		},
	},
}

func ToXlinkShow(t string) (*XlinkShowEnum, error) {
	switch t {
	case XlinkShow.New.String():
		return &XlinkShow.New, nil
	case XlinkShow.Replace.String():
		return &XlinkShow.Replace, nil
	case XlinkShow.Embed.String():
		return &XlinkShow.Embed, nil
	case XlinkShow.Other.String():
		return &XlinkShow.Other, nil
	case XlinkShow.None.String():
		return &XlinkShow.None, nil
	}
	return nil, fmt.Errorf("can not convert to XlinkShow. t=%s", t)
}

func AllXlinkShowEnumValues() []XlinkShowEnum {
	values := []XlinkShowEnum{
		XlinkShow.New,
		XlinkShow.Replace,
		XlinkShow.Embed,
		XlinkShow.Other,
		XlinkShow.None,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *XlinkShowEnum) Equals(obj XlinkShowEnum) bool {
	return *e == obj
}

func (e *XlinkShowEnum) In(objs ...XlinkShowEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *XlinkShowEnum) Ordinal() string {
	return XlinkShow.opts[e.String()]["ordinal"].(string)
}

func (e *XlinkShowEnum) String() string {
	return string(*e)
}

func (e *XlinkShowEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
