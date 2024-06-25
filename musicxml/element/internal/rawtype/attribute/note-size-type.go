package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type NoteSizeType enum.NoteSizeTypeEnum

func (nst *NoteSizeType) UnmarshalXMLAttr(attr xml.Attr) error {
	nstV, err := enum.ToNoteSizeType(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*nst = (NoteSizeType)(*nstV)

	return nil
}
