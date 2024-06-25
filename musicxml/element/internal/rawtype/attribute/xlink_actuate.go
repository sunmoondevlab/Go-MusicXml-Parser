package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type XlinkActuate enum.XlinkActuateEnum

func (xa *XlinkActuate) UnmarshalXMLAttr(attr xml.Attr) error {
	xaV, err := enum.ToXlinkActuate(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*xa = (XlinkActuate)(*xaV)

	return nil
}
