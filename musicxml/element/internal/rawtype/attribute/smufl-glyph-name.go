package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type SmuflGlyphName string

func (sgn *SmuflGlyphName) UnmarshalXMLAttr(attr xml.Attr) error {
	_, err := datatype.NewSmuflGlyphName(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}: invalid smufl-glyph-name, %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*sgn = (SmuflGlyphName)(attr.Value)

	return nil
}

func (sgn *SmuflGlyphName) ToEntityData() *datatype.SmuflGlyphName {
	if sgn == nil {
		return nil
	}

	sgnV, err := datatype.NewSmuflGlyphName((string)(*sgn))
	if err != nil {
		return nil
	}
	return sgnV
}
