package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type CreatorType enum.CreatorTypeEnum

func (ct *CreatorType) UnmarshalXMLAttr(attr xml.Attr) error {
	ctV, err := enum.ToCreatorType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*ct = (CreatorType)(*ctV)

	return nil
}
