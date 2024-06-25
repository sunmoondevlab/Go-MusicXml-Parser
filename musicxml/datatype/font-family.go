package datatype

import (
	"strings"
)

type FontFamily []string

func NewFontFamilyL(ffR string) *FontFamily {
	ff := FontFamily{}
	if strings.TrimSpace(ffR) == "" {
		return nil
	}
	ffLT := strings.Split(ffR, ",")
	for _, ffE := range ffLT {
		ffE = strings.TrimSpace(ffE)
		ff = append(ff, ffE)
	}
	return &ff
}

func (ff *FontFamily) String() string {
	if ff == nil {
		return ""
	}
	return strings.Join(*ff, ",")
}

func (ff *FontFamily) StringPtr() *string {
	if ff == nil {
		return nil
	}
	s := strings.Join(*ff, ",")
	return &s
}
