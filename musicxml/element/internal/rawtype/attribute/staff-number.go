package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type StaffNumber string

func (sn *StaffNumber) UnmarshalXMLAttr(attr xml.Attr) error {
	_, err := datatype.NewStaffNumber(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}, invalid StaffNumber, %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*sn = (StaffNumber)(attr.Value)

	return nil
}

func (sn *StaffNumber) ToEntityData() *datatype.StaffNumber {
	if sn == nil {
		return nil
	}
	snV, _ := datatype.NewStaffNumber((string)(*sn))
	return snV
}
