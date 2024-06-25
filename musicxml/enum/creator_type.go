package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type CreatorTypeEnum string

var CreatorType = struct {
	Composer   CreatorTypeEnum
	Arranger   CreatorTypeEnum
	Lyricist   CreatorTypeEnum
	Poet       CreatorTypeEnum
	Translator CreatorTypeEnum
	opts       map[string]map[string]interface{}
}{
	Composer:   CreatorTypeEnum("composer"),
	Arranger:   CreatorTypeEnum("arranger"),
	Lyricist:   CreatorTypeEnum("lyricist"),
	Poet:       CreatorTypeEnum("poet"),
	Translator: CreatorTypeEnum("translator"),
	opts: map[string]map[string]interface{}{
		"composer": {
			"ordinal": "0",
		},
		"arranger": {
			"ordinal": "1",
		},
		"lyricist": {
			"ordinal": "2",
		},
		"poet": {
			"ordinal": "3",
		},
		"translator": {
			"ordinal": "4",
		},
	},
}

func ToCreatorType(t string) (*CreatorTypeEnum, error) {
	switch t {
	case CreatorType.Composer.String():
		return &CreatorType.Composer, nil
	case CreatorType.Arranger.String():
		return &CreatorType.Arranger, nil
	case CreatorType.Lyricist.String():
		return &CreatorType.Lyricist, nil
	case CreatorType.Poet.String():
		return &CreatorType.Poet, nil
	case CreatorType.Translator.String():
		return &CreatorType.Translator, nil
	}
	return nil, fmt.Errorf("invalid <creator/> attr. attr => xlink:actuate actuate=%s", t)
}

func AllCreatorTypeEnumValues() []CreatorTypeEnum {
	values := []CreatorTypeEnum{
		CreatorType.Composer,
		CreatorType.Arranger,
		CreatorType.Lyricist,
		CreatorType.Poet,
		CreatorType.Translator,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *CreatorTypeEnum) Equals(obj CreatorTypeEnum) bool {
	return *e == obj
}

func (e *CreatorTypeEnum) In(objs ...CreatorTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *CreatorTypeEnum) Ordinal() string {
	return CreatorType.opts[e.String()]["ordinal"].(string)
}

func (e *CreatorTypeEnum) String() string {
	return string(*e)
}

func (e *CreatorTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
