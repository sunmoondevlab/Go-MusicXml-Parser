package supports

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type Supports struct {
	XMLName   xml.Name       `xml:"supports"`
	Element   string         `xml:"element"`
	Type      enum.YesNoEnum `xml:"type"`
	Attribute string         `xml:"attribute"`
	Value     string         `xml:"value"`
}

type SupportsL []Supports

func (sl *SupportsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		ser := &struct {
			XMLName   xml.Name `xml:"supports"`
			Element   string   `xml:"element,attr"`
			Type      string   `xml:"type,attr"`
			Attribute string   `xml:"attribute,attr"`
			Value     string   `xml:"value,attr"`
		}{}
		err := d.DecodeElement(ser, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		yn, err := enum.ToYesNo(ser.Type, ser.XMLName.Local, "type")
		if err != nil {
			return err
		}
		ce := Supports{
			XMLName:   ser.XMLName,
			Element:   ser.Element,
			Type:      *yn,
			Attribute: ser.Attribute,
			Value:     ser.Value,
		}
		(*sl) = append(*sl, ce)
	}

	return nil
}
