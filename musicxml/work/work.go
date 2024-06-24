package work

import (
	"encoding/xml"
	"reflect"

	tWorkOpus "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work/workopus"
)

type Work struct {
	XMLName    xml.Name            `xml:"work"`
	WorkTitle  *string             `xml:"work-title,omitempty"`
	WorkNumber *string             `xml:"work-number,omitempty"`
	WorkOpus   *tWorkOpus.WorkOpus `xml:"opus,omitempty"`
}

func (wo *Work) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	wr := &struct {
		XMLName    xml.Name            `xml:"work"`
		WorkTitle  *string             `xml:"work-title,omitempty"`
		WorkNumber *string             `xml:"work-number,omitempty"`
		WorkOpus   *tWorkOpus.WorkOpus `xml:"opus,omitempty"`
	}{}

	err := d.DecodeElement(wr, startP)
	if err != nil {
		return err
	}
	*wo = Work{
		XMLName:    wr.XMLName,
		WorkTitle:  wr.WorkTitle,
		WorkNumber: wr.WorkNumber,
		WorkOpus:   wr.WorkOpus,
	}

	return err
}

func (wo *Work) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = wo.XMLName

	wr := &struct {
		XMLName    xml.Name            `xml:"work"`
		WorkTitle  *string             `xml:"work-title,omitempty"`
		WorkNumber *string             `xml:"work-number,omitempty"`
		WorkOpus   *tWorkOpus.WorkOpus `xml:"opus,omitempty"`
	}{
		XMLName:    wo.XMLName,
		WorkTitle:  wo.WorkTitle,
		WorkNumber: wo.WorkNumber,
		WorkOpus:   wo.WorkOpus,
	}

	return e.EncodeElement(wr, start)
}
