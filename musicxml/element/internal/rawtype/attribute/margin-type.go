package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type MarginType enum.MarginTypeEnum

func (mt *MarginType) UnmarshalXMLAttr(attr xml.Attr) error {
	mtV, err := enum.ToMarginType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*mt = (MarginType)(*mtV)

	return nil
}
