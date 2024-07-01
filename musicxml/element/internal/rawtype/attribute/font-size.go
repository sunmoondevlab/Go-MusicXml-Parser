package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type FontSize string

func (fs *FontSize) UnmarshalXMLAttr(attr xml.Attr) error {
	_, err := datatype.NewFontSize(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*fs = (FontSize)(attr.Value)

	return nil
}

func (fs *FontSize) ToEntityData() *datatype.FontSize {
	if fs == nil {
		return nil
	}
	fsV, _ := datatype.NewFontSize((string)(*fs))
	return fsV
}
