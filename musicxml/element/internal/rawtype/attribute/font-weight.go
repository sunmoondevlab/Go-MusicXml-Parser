package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type FontWeight enum.FontWeightEnum

func (fw *FontWeight) UnmarshalXMLAttr(attr xml.Attr) error {
	fwV, err := enum.NewFontWeight(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*fw = (FontWeight)(*fwV)

	return nil
}
