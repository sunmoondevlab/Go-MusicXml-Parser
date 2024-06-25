package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Identification struct {
	XMLName       xml.Name       `xml:"identification"`
	Creator       *Creator       `xml:"creator,omitempty"`
	Rights        *RightsL       `xml:"rights,omitempty"`
	Encoding      *Encoding      `xml:"encoding,omitempty"`
	Source        *string        `xml:"source,omitempty"`
	Relation      *RelationL     `xml:"relation,omitempty"`
	Miscellaneous *Miscellaneous `xml:"miscellaneous,omitempty"`
}

type IdentificationRaw struct {
	XMLName       xml.Name       `xml:"identification"`
	Creator       *Creator       `xml:"creator,omitempty"`
	Rights        *RightsL       `xml:"rights,omitempty"`
	Encoding      *Encoding      `xml:"encoding,omitempty"`
	Source        *string        `xml:"source,omitempty"`
	Relation      *RelationL     `xml:"relation,omitempty"`
	Miscellaneous *Miscellaneous `xml:"miscellaneous,omitempty"`
}

func (iO *Identification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	iR := &IdentificationRaw{}
	err := d.DecodeElement(iR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*iO = Identification(*iR)

	return nil
}

func (iO *Identification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	iR := Identification(*iO)

	start.Name = iO.XMLName

	return e.EncodeElement(iR, start)
}
