package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type LineWidthType enum.LineWidthTypeEnum

func (lwt *LineWidthType) UnmarshalXMLAttr(attr xml.Attr) error {
	lwtV, err := enum.ToLineWidthType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*lwt = (LineWidthType)(*lwtV)

	return nil
}
