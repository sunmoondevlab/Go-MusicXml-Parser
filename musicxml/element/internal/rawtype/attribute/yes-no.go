package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type YesNo enum.YesNoEnum

func (yn *YesNo) UnmarshalXMLAttr(attr xml.Attr) error {
	ynV, err := enum.ToYesNo(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*yn = (YesNo)(*ynV)

	return nil
}
