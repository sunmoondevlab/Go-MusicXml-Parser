package datatype

import (
	"fmt"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype/internal/glyphnames"
)

type SmuflGlyphName struct {
	Val *string
	fmt.Stringer
}

func (sgn *SmuflGlyphName) String() string {
	if sgn == nil {
		return ""
	}
	if sgn.Val == nil {
		return ""
	}
	return *sgn.Val
}

func (sgn *SmuflGlyphName) StringPtr() *string {
	if sgn == nil {
		return nil
	}
	return sgn.Val
}

func NewSmuflGlyphName(value string) (*SmuflGlyphName, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, nil
	}
	_, fnd := glyphnames.SpecGlyphNames.BinarySearchByName(value)
	if !fnd {
		return nil, fmt.Errorf("can not convert to SmuflGlyphName. Please specify the SMuFL regular glyph name as the value. see berow: [https://w3c.github.io/smufl/latest/index.html]. value=%s", value)
	}

	return &SmuflGlyphName{
		Val: &value,
	}, nil
}
