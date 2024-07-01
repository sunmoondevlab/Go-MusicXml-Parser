package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type DistanceTypeEnum string

var DistanceType = struct {
	Beam   DistanceTypeEnum
	Hyphen DistanceTypeEnum
	opts   map[string]map[string]interface{}
}{
	Beam:   DistanceTypeEnum("beam"),
	Hyphen: DistanceTypeEnum("hyphen"),
	opts: map[string]map[string]interface{}{
		"beam": {
			"ordinal": "0",
		},
		"hyphen": {
			"ordinal": "1",
		},
	},
}

func ToDistanceType(t string) (*DistanceTypeEnum, error) {
	switch t {
	case DistanceType.Beam.String():
		return &DistanceType.Beam, nil
	case DistanceType.Hyphen.String():
		return &DistanceType.Hyphen, nil
	}
	return nil, fmt.Errorf("can not convert to DistanceType. t=%s", t)
}

func AllDistanceTypeEnumValues() []DistanceTypeEnum {
	values := []DistanceTypeEnum{
		DistanceType.Beam,
		DistanceType.Hyphen,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *DistanceTypeEnum) Equals(obj DistanceTypeEnum) bool {
	return *e == obj
}

func (e *DistanceTypeEnum) In(objs ...DistanceTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *DistanceTypeEnum) Ordinal() string {
	return DistanceType.opts[e.String()]["ordinal"].(string)
}

func (e *DistanceTypeEnum) String() string {
	return string(*e)
}

func (e *DistanceTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
