package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Miscellaneous struct {
	XMLName            xml.Name             `xml:"miscellaneous"`
	MiscellaneousField *MiscellaneousFieldL `xml:"miscellaneous-field,omitempty"`
}

type MiscellaneousRaw struct {
	XMLName            xml.Name             `xml:"miscellaneous"`
	MiscellaneousField *MiscellaneousFieldL `xml:"miscellaneous-field,omitempty"`
}

func (mO *Miscellaneous) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}

	mR := &MiscellaneousRaw{}
	err := d.DecodeElement(mR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	*mO = Miscellaneous(*mR)

	return nil
}

func (mO *Miscellaneous) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	mR := MiscellaneousRaw(*mO)

	start.Name = mO.XMLName

	return e.EncodeElement(mR, start)
}
