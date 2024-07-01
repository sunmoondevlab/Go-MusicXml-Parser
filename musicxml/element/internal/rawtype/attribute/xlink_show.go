package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type XlinkShow enum.XlinkShowEnum

func (xs *XlinkShow) UnmarshalXMLAttr(attr xml.Attr) error {
	xsV, err := enum.ToXlinkShow(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*xs = (XlinkShow)(*xsV)

	return nil
}
