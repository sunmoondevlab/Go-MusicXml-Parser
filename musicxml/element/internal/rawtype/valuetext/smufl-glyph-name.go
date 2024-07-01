package valuetext

import (
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type SmuflGlyphName string

func (sgn *SmuflGlyphName) UnmarshalText(b []byte) error {
	s := string(b)

	_, err := datatype.NewSmuflGlyphName(s)
	if err != nil {
		return fmt.Errorf("invalid value: invalid smufl-glyph-name, %s", err.Error())
	}
	*sgn = (SmuflGlyphName)(s)

	return nil
}

func (sgn SmuflGlyphName) ToEntityData() datatype.SmuflGlyphName {
	if sgn == "" {
		return datatype.SmuflGlyphName{}
	}

	sgnV, err := datatype.NewSmuflGlyphName((string)(sgn))
	if err != nil {
		return datatype.SmuflGlyphName{}
	}
	return *sgnV
}
