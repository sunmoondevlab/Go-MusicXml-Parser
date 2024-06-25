package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type Tenths string

func (t *Tenths) UnmarshalXMLAttr(attr xml.Attr) error {
	_, err := datatype.NewTenthsFromString(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid tenths, %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*t = (Tenths)(attr.Value)

	return nil
}

func (t *Tenths) ToEntityData() *datatype.Tenths {
	if t == nil {
		return nil
	}

	tV, _ := datatype.NewTenthsFromString((string)(*t))
	return tV
}
