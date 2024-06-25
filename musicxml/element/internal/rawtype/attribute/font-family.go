package attribute

import (
	"encoding/xml"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type FontFamily string

func (ff *FontFamily) UnmarshalXMLAttr(attr xml.Attr) error {
	*ff = (FontFamily)(attr.Value)

	return nil
}

func (ff *FontFamily) ToEntityData() *datatype.FontFamily {
	if ff == nil {
		return nil
	}
	ffV := datatype.NewFontFamilyL((string)(*ff))
	return ffV
}
