package work

import (
	"encoding/xml"
	"reflect"

	tWorkOpus "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work/workopus"
)

type Work struct {
	XMLName    xml.Name           `xml:"work"`
	WorkTitle  string             `xml:"work-title"`
	WorkNumber string             `xml:"work-number"`
	WorkOpus   tWorkOpus.WorkOpus `xml:"opus"`
}

func (w *Work) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	wr := &struct {
		XMLName    xml.Name           `xml:"work"`
		WorkTitle  string             `xml:"work-title"`
		WorkNumber string             `xml:"work-number"`
		WorkOpus   tWorkOpus.WorkOpus `xml:"opus"`
	}{}

	err := d.DecodeElement(wr, startP)
	if err != nil {
		return err
	}
	*w = Work{
		XMLName:    wr.XMLName,
		WorkTitle:  wr.WorkTitle,
		WorkNumber: wr.WorkNumber,
		WorkOpus:   wr.WorkOpus,
	}

	return err
}
