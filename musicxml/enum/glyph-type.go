package enum

import (
	"fmt"
	"sort"
	"strconv"
)

type GlyphTypeEnum string

var GlyphType = struct {
	QuarterRest           GlyphTypeEnum
	GClefOttavaBassa      GlyphTypeEnum
	CClef                 GlyphTypeEnum
	FClef                 GlyphTypeEnum
	PercussionClef        GlyphTypeEnum
	OctaveShiftUp8        GlyphTypeEnum
	OctaveShiftDown8      GlyphTypeEnum
	OctaveShiftContinue8  GlyphTypeEnum
	OctaveShiftDown15     GlyphTypeEnum
	OctaveShiftUp15       GlyphTypeEnum
	OctaveShiftContinue15 GlyphTypeEnum
	OctaveShiftDown22     GlyphTypeEnum
	OctaveShiftUp22       GlyphTypeEnum
	OctaveShiftContinue22 GlyphTypeEnum
	opts                  map[string]map[string]interface{}
}{
	QuarterRest:           GlyphTypeEnum("quarter-rest"),
	GClefOttavaBassa:      GlyphTypeEnum("g-clef-ottava-bassa"),
	CClef:                 GlyphTypeEnum("c-clef"),
	FClef:                 GlyphTypeEnum("f-clef"),
	PercussionClef:        GlyphTypeEnum("percussion-clef"),
	OctaveShiftUp8:        GlyphTypeEnum("octave-shift-up-8"),
	OctaveShiftDown8:      GlyphTypeEnum("octave-shift-down-8"),
	OctaveShiftContinue8:  GlyphTypeEnum("octave-shift-continue-8"),
	OctaveShiftDown15:     GlyphTypeEnum("octave-shift-down-15"),
	OctaveShiftUp15:       GlyphTypeEnum("octave-shift-up-15"),
	OctaveShiftContinue15: GlyphTypeEnum("octave-shift-continue-15"),
	OctaveShiftDown22:     GlyphTypeEnum("octave-shift-down-22"),
	OctaveShiftUp22:       GlyphTypeEnum("octave-shift-up-22"),
	OctaveShiftContinue22: GlyphTypeEnum("octave-shift-continue-22"),
	opts: map[string]map[string]interface{}{
		"quarter-rest": {
			"ordinal": "0",
		},
		"g-clef-ottava-bassa": {
			"ordinal": "1",
		},
		"c-clef": {
			"ordinal": "2",
		},
		"f-clef": {
			"ordinal": "3",
		},
		"percussion-clef": {
			"ordinal": "4",
		},
		"octave-shift-up-8": {
			"ordinal": "5",
		},
		"octave-shift-down-8": {
			"ordinal": "6",
		},
		"octave-shift-continue-8": {
			"ordinal": "7",
		},
		"octave-shift-up-15": {
			"ordinal": "8",
		},
		"octave-shift-down-15": {
			"ordinal": "9",
		},
		"octave-shift-continue-15": {
			"ordinal": "10",
		},
		"octave-shift-up-22": {
			"ordinal": "11",
		},
		"octave-shift-down-22": {
			"ordinal": "12",
		},
		"octave-shift-continue-22": {
			"ordinal": "13",
		},
	},
}

func ToGlyphType(t string) (*GlyphTypeEnum, error) {
	switch t {
	case GlyphType.QuarterRest.String():
		return &GlyphType.QuarterRest, nil
	case GlyphType.GClefOttavaBassa.String():
		return &GlyphType.GClefOttavaBassa, nil
	case GlyphType.CClef.String():
		return &GlyphType.CClef, nil
	case GlyphType.FClef.String():
		return &GlyphType.FClef, nil
	case GlyphType.PercussionClef.String():
		return &GlyphType.PercussionClef, nil
	case GlyphType.OctaveShiftUp8.String():
		return &GlyphType.OctaveShiftUp8, nil
	case GlyphType.OctaveShiftDown8.String():
		return &GlyphType.OctaveShiftDown8, nil
	case GlyphType.OctaveShiftContinue8.String():
		return &GlyphType.OctaveShiftContinue8, nil
	case GlyphType.OctaveShiftDown15.String():
		return &GlyphType.OctaveShiftDown15, nil
	case GlyphType.OctaveShiftUp15.String():
		return &GlyphType.OctaveShiftUp15, nil
	case GlyphType.OctaveShiftContinue15.String():
		return &GlyphType.OctaveShiftContinue15, nil
	case GlyphType.OctaveShiftDown22.String():
		return &GlyphType.OctaveShiftDown22, nil
	case GlyphType.OctaveShiftUp22.String():
		return &GlyphType.OctaveShiftUp22, nil
	case GlyphType.OctaveShiftContinue22.String():
		return &GlyphType.OctaveShiftContinue22, nil
	}
	return nil, fmt.Errorf("can not convert to GlyphType. t=%s", t)
}

func AllGlyphTypeEnumValues() []GlyphTypeEnum {
	values := []GlyphTypeEnum{
		GlyphType.QuarterRest,
		GlyphType.GClefOttavaBassa,
		GlyphType.CClef,
		GlyphType.FClef,
		GlyphType.PercussionClef,
		GlyphType.OctaveShiftUp8,
		GlyphType.OctaveShiftDown8,
		GlyphType.OctaveShiftContinue8,
		GlyphType.OctaveShiftDown15,
		GlyphType.OctaveShiftUp15,
		GlyphType.OctaveShiftContinue15,
		GlyphType.OctaveShiftDown22,
		GlyphType.OctaveShiftUp22,
		GlyphType.OctaveShiftContinue22,
	}

	sort.Slice(values, func(i, j int) bool {
		ordinalI, _ := strconv.Atoi(values[i].Ordinal())
		ordinalJ, _ := strconv.Atoi(values[j].Ordinal())
		return ordinalI < ordinalJ
	})

	return values
}

func (e *GlyphTypeEnum) Equals(obj GlyphTypeEnum) bool {
	return *e == obj
}

func (e *GlyphTypeEnum) In(objs ...GlyphTypeEnum) bool {
	for _, obj := range objs {
		if e.Equals(obj) {
			return true
		}
	}

	return false
}

func (e *GlyphTypeEnum) Ordinal() string {
	return GlyphType.opts[e.String()]["ordinal"].(string)
}

func (e *GlyphTypeEnum) String() string {
	return string(*e)
}

func (e *GlyphTypeEnum) StringPtr() *string {
	s := string(*e)
	return &s
}
