package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type LineWidthTypeEnum string

var LineWidthType = struct {
	Beam          LineWidthTypeEnum
	Bracket       LineWidthTypeEnum
	Dashes        LineWidthTypeEnum
	Enclosure     LineWidthTypeEnum
	Ending        LineWidthTypeEnum
	Extend        LineWidthTypeEnum
	HeavyBarline  LineWidthTypeEnum
	Leger         LineWidthTypeEnum
	LightBarline  LineWidthTypeEnum
	OctaveShift   LineWidthTypeEnum
	Pedal         LineWidthTypeEnum
	SlurMiddle    LineWidthTypeEnum
	SlurTip       LineWidthTypeEnum
	Staff         LineWidthTypeEnum
	Stem          LineWidthTypeEnum
	TieMiddle     LineWidthTypeEnum
	TieTip        LineWidthTypeEnum
	TupletBracket LineWidthTypeEnum
	Wedge         LineWidthTypeEnum
	opts          map[string]map[string]interface{}
}{
	Beam:          LineWidthTypeEnum("beam"),
	Bracket:       LineWidthTypeEnum("bracket"),
	Dashes:        LineWidthTypeEnum("dashes"),
	Enclosure:     LineWidthTypeEnum("enclosure"),
	Ending:        LineWidthTypeEnum("ending"),
	Extend:        LineWidthTypeEnum("extend"),
	HeavyBarline:  LineWidthTypeEnum("heavy barline"),
	Leger:         LineWidthTypeEnum("leger"),
	LightBarline:  LineWidthTypeEnum("light barline"),
	OctaveShift:   LineWidthTypeEnum("octave shift"),
	Pedal:         LineWidthTypeEnum("pedal"),
	SlurMiddle:    LineWidthTypeEnum("slur middle"),
	SlurTip:       LineWidthTypeEnum("slur tip"),
	Staff:         LineWidthTypeEnum("staff"),
	Stem:          LineWidthTypeEnum("stem"),
	TieMiddle:     LineWidthTypeEnum("tie middle"),
	TieTip:        LineWidthTypeEnum("tie tip"),
	TupletBracket: LineWidthTypeEnum("tuplet bracket"),
	Wedge:         LineWidthTypeEnum("wedge"),
	opts: map[string]map[string]interface{}{
		"beam": {
			"ordinal": "0",
		},
		"bracket": {
			"ordinal": "1",
		},
		"dashes": {
			"ordinal": "2",
		},
		"enclosure": {
			"ordinal": "3",
		},
		"ending": {
			"ordinal": "4",
		},
		"extend": {
			"ordinal": "5",
		},
		"heavy barline": {
			"ordinal": "6",
		},
		"leger": {
			"ordinal": "7",
		},
		"light barline": {
			"ordinal": "8",
		},
		"octave shift": {
			"ordinal": "9",
		},
		"pedal": {
			"ordinal": "10",
		},
		"slur middle": {
			"ordinal": "11",
		},
		"slur tip": {
			"ordinal": "12",
		},
		"staff": {
			"ordinal": "13",
		},
		"stem": {
			"ordinal": "14",
		},
		"tie middle": {
			"ordinal": "15",
		},
		"tie tip": {
			"ordinal": "16",
		},
		"tuplet bracket": {
			"ordinal": "17",
		},
		"wedge": {
			"ordinal": "18",
		},
	},
}

func ToLineWidthType(t string) (*LineWidthTypeEnum, error) {
	switch t {
	case LineWidthType.Beam.String():
		return &LineWidthType.Beam, nil
	case LineWidthType.Bracket.String():
		return &LineWidthType.Bracket, nil
	case LineWidthType.Dashes.String():
		return &LineWidthType.Dashes, nil
	case LineWidthType.Enclosure.String():
		return &LineWidthType.Enclosure, nil
	case LineWidthType.Ending.String():
		return &LineWidthType.Ending, nil
	case LineWidthType.Extend.String():
		return &LineWidthType.Extend, nil
	case LineWidthType.HeavyBarline.String():
		return &LineWidthType.HeavyBarline, nil
	case LineWidthType.Leger.String():
		return &LineWidthType.Leger, nil
	case LineWidthType.LightBarline.String():
		return &LineWidthType.LightBarline, nil
	case LineWidthType.OctaveShift.String():
		return &LineWidthType.OctaveShift, nil
	case LineWidthType.Pedal.String():
		return &LineWidthType.Pedal, nil
	case LineWidthType.SlurMiddle.String():
		return &LineWidthType.SlurMiddle, nil
	case LineWidthType.SlurTip.String():
		return &LineWidthType.SlurTip, nil
	case LineWidthType.Staff.String():
		return &LineWidthType.Staff, nil
	case LineWidthType.Stem.String():
		return &LineWidthType.Stem, nil
	case LineWidthType.TieMiddle.String():
		return &LineWidthType.TieMiddle, nil
	case LineWidthType.TieTip.String():
		return &LineWidthType.TieTip, nil
	case LineWidthType.TupletBracket.String():
		return &LineWidthType.TupletBracket, nil
	case LineWidthType.Wedge.String():
		return &LineWidthType.Wedge, nil
	}
	return nil, fmt.Errorf("can not convert to LineWidthType. t=%s", t)
}

func AllLineWidthTypeEnumValues() []LineWidthTypeEnum {
	values := []LineWidthTypeEnum{
		LineWidthType.Beam,
		LineWidthType.Bracket,
		LineWidthType.Dashes,
		LineWidthType.Enclosure,
		LineWidthType.Ending,
		LineWidthType.Extend,
		LineWidthType.HeavyBarline,
		LineWidthType.Leger,
		LineWidthType.LightBarline,
		LineWidthType.OctaveShift,
		LineWidthType.Pedal,
		LineWidthType.SlurMiddle,
		LineWidthType.SlurTip,
		LineWidthType.Staff,
		LineWidthType.Stem,
		LineWidthType.TieMiddle,
		LineWidthType.TieTip,
		LineWidthType.TupletBracket,
		LineWidthType.Wedge,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *LineWidthTypeEnum) Equals(obj LineWidthTypeEnum) bool {
	return *e == obj
}

func (e *LineWidthTypeEnum) In(objs ...LineWidthTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *LineWidthTypeEnum) Ordinal() string {
	return LineWidthType.opts[e.String()]["ordinal"].(string)
}

func (e *LineWidthTypeEnum) String() string {
	return string(*e)
}

func (e *LineWidthTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
