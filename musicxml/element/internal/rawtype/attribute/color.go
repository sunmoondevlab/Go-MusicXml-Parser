package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type Color string

func (c *Color) UnmarshalXMLAttr(attr xml.Attr) error {
	err := datatype.ValidateColorCodeRequireUpper(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid Color, %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*c = (Color)(attr.Value)

	return nil
}

func (c *Color) ToEntityData() *datatype.Color {
	if c == nil {
		return nil
	}
	cV, _ := datatype.NewColorRequireUpper((string)(*c))
	return cV
}
