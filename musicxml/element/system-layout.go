package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type SystemLayout struct {
	XMLName           xml.Name           `xml:"system-layout"`
	SystemMargins     *SystemMargins     `xml:"system-margins,omitempty"`
	SystemDistance    *SystemDistance    `xml:"system-distance,omitempty"`
	TopSystemDistance *TopSystemDistance `xml:"top-system-distance,omitempty"`
	SystemDividers    *SystemDividers    `xml:"system-dividers"`
}

type SystemLayoutRaw struct {
	XMLName           xml.Name           `xml:"system-layout"`
	SystemMargins     *SystemMargins     `xml:"system-margins,omitempty"`
	SystemDistance    *SystemDistance    `xml:"system-distance,omitempty"`
	TopSystemDistance *TopSystemDistance `xml:"top-system-distance,omitempty"`
	SystemDividers    *SystemDividers    `xml:"system-dividers"`
}

func (slO *SystemLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	slR := &SystemLayoutRaw{}

	err := d.DecodeElement(slR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	*slO = SystemLayout(*slR)

	return err
}

func (slO *SystemLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = slO.XMLName

	slR := SystemLayoutRaw(*slO)

	return e.EncodeElement(slR, start)
}
