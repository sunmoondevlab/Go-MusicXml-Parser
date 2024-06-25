package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type FontStyle enum.FontStyleEnum

func (fs *FontStyle) UnmarshalXMLAttr(attr xml.Attr) error {
	fsV, err := enum.NewFontStyle(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*fs = (FontStyle)(*fsV)

	return nil
}
