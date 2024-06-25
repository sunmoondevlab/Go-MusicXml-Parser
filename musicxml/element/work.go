package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Work struct {
	XMLName    xml.Name `xml:"work"`
	WorkTitle  *string  `xml:"work-title,omitempty"`
	WorkNumber *string  `xml:"work-number,omitempty"`
	Opus       *Opus    `xml:"opus,omitempty"`
}

type WorkRaw struct {
	XMLName    xml.Name `xml:"work"`
	WorkTitle  *string  `xml:"work-title,omitempty"`
	WorkNumber *string  `xml:"work-number,omitempty"`
	Opus       *Opus    `xml:"opus,omitempty"`
}

func (wO *Work) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	wR := &WorkRaw{}

	err := d.DecodeElement(wR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	*wO = Work(*wR)

	return err
}

func (wO *Work) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = wO.XMLName

	wR := WorkRaw(*wO)

	return e.EncodeElement(wR, start)
}
