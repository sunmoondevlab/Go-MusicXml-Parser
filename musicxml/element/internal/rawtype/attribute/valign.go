package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type Valign enum.ValignEnum

func (v *Valign) UnmarshalXMLAttr(attr xml.Attr) error {
	vV, err := enum.ToValign(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*v = (Valign)(*vV)

	return nil
}
