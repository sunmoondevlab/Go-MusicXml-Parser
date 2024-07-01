package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type XlinkType enum.XlinkTypeEnum

func (xt *XlinkType) UnmarshalXMLAttr(attr xml.Attr) error {
	xtV, err := enum.ToXlinkType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*xt = (XlinkType)(*xtV)

	return nil
}
