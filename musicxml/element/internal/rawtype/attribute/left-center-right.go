package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type LeftCenterRight enum.LeftCenterRightEnum

func (lcr *LeftCenterRight) UnmarshalXMLAttr(attr xml.Attr) error {
	lcrV, err := enum.ToLeftCenterRight(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*lcr = (LeftCenterRight)(*lcrV)

	return nil
}
