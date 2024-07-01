package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type DistanceType enum.DistanceTypeEnum

func (dt *DistanceType) UnmarshalXMLAttr(attr xml.Attr) error {
	dtV, err := enum.ToDistanceType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*dt = (DistanceType)(*dtV)

	return nil
}
