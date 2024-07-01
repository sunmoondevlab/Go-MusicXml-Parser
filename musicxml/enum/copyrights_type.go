package enum

import (
	"sort"
	"strconv"
)

type CopyrightsTypeEnum string

var CopyrightsType = struct {
	Music       CopyrightsTypeEnum
	Arrangement CopyrightsTypeEnum
	Words       CopyrightsTypeEnum
	Translation CopyrightsTypeEnum
	Parody      CopyrightsTypeEnum
	Other       CopyrightsTypeEnum
	opts        map[string]map[string]interface{}
}{
	Music:       CopyrightsTypeEnum("music"),
	Arrangement: CopyrightsTypeEnum("arrangement"),
	Words:       CopyrightsTypeEnum("words"),
	Translation: CopyrightsTypeEnum("translation"),
	Parody:      CopyrightsTypeEnum("parody"),
	Other:       CopyrightsTypeEnum("other"),
	opts: map[string]map[string]interface{}{
		"music": {
			"ordinal": "0",
		},
		"arrangement": {
			"ordinal": "1",
		},
		"words": {
			"ordinal": "2",
		},
		"translation": {
			"ordinal": "3",
		},
		"parody": {
			"ordinal": "4",
		},
		"other": {
			"ordinal": "5",
		},
	},
}

func ToCopyrightsType(t string) (*CopyrightsTypeEnum, error) {
	switch t {
	case CopyrightsType.Music.String():
		return &CopyrightsType.Music, nil
	case CopyrightsType.Arrangement.String():
		return &CopyrightsType.Arrangement, nil
	case CopyrightsType.Words.String():
		return &CopyrightsType.Words, nil
	case CopyrightsType.Translation.String():
		return &CopyrightsType.Translation, nil
	case CopyrightsType.Parody.String():
		return &CopyrightsType.Parody, nil
	default:
		return &CopyrightsType.Other, nil
	}
}

func AllCopyrightsTypeEnumValues() []CopyrightsTypeEnum {
	values := []CopyrightsTypeEnum{
		CopyrightsType.Music,
		CopyrightsType.Arrangement,
		CopyrightsType.Words,
		CopyrightsType.Translation,
		CopyrightsType.Parody,
		CopyrightsType.Other,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *CopyrightsTypeEnum) Equals(obj CopyrightsTypeEnum) bool {
	return *e == obj
}

func (e *CopyrightsTypeEnum) In(objs ...CopyrightsTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *CopyrightsTypeEnum) Ordinal() string {
	return CopyrightsType.opts[e.String()]["ordinal"].(string)
}

func (e *CopyrightsTypeEnum) String() string {
	return string(*e)
}

func (e *CopyrightsTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
