package enum

import (
	"sort"
	"strconv"
)

type CopyRightsTypeEnum string

var CopyRightsType = struct {
	All         CopyRightsTypeEnum
	Music       CopyRightsTypeEnum
	Arrangement CopyRightsTypeEnum
	Words       CopyRightsTypeEnum
	Translation CopyRightsTypeEnum
	Parody      CopyRightsTypeEnum
	Other       CopyRightsTypeEnum
	opts        map[string]map[string]interface{}
}{
	All:         CopyRightsTypeEnum(""),
	Music:       CopyRightsTypeEnum("music"),
	Arrangement: CopyRightsTypeEnum("arrangement"),
	Words:       CopyRightsTypeEnum("words"),
	Translation: CopyRightsTypeEnum("translation"),
	Parody:      CopyRightsTypeEnum("parody"),
	Other:       CopyRightsTypeEnum("other"),
	opts: map[string]map[string]interface{}{
		"": {
			"ordinal": "0",
		},
		"music": {
			"ordinal": "1",
		},
		"arrangement": {
			"ordinal": "2",
		},
		"words": {
			"ordinal": "3",
		},
		"translation": {
			"ordinal": "4",
		},
		"parody": {
			"ordinal": "5",
		},
		"other": {
			"ordinal": "6",
		},
	},
}

func ToCopyRightsType(t string) (*CopyRightsTypeEnum, error) {
	switch t {
	case CopyRightsType.All.String():
		return &CopyRightsType.All, nil
	case CopyRightsType.Music.String():
		return &CopyRightsType.Music, nil
	case CopyRightsType.Arrangement.String():
		return &CopyRightsType.Arrangement, nil
	case CopyRightsType.Words.String():
		return &CopyRightsType.Words, nil
	case CopyRightsType.Translation.String():
		return &CopyRightsType.Translation, nil
	case CopyRightsType.Parody.String():
		return &CopyRightsType.Parody, nil
	default:
		return &CopyRightsType.Other, nil
	}
}

func AllCopyRightsTypeEnumValues() []CopyRightsTypeEnum {
	values := []CopyRightsTypeEnum{
		CopyRightsType.All,
		CopyRightsType.Music,
		CopyRightsType.Arrangement,
		CopyRightsType.Words,
		CopyRightsType.Translation,
		CopyRightsType.Parody,
		CopyRightsType.Other,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *CopyRightsTypeEnum) Equals(obj CopyRightsTypeEnum) bool {
	return *e == obj
}

func (e *CopyRightsTypeEnum) In(objs ...CopyRightsTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *CopyRightsTypeEnum) Ordinal() string {
	return CopyRightsType.opts[e.String()]["ordinal"].(string)
}

func (e *CopyRightsTypeEnum) String() string {
	return string(*e)
}

func (e *CopyRightsTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
