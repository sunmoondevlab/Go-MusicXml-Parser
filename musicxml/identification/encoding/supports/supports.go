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
	Attribute *string        `xml:"attribute,omitempty"`
	Value     *string        `xml:"value,omitempty"`
}

type SupportsL []Supports

func (sL *SupportsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		type serT struct {
			XMLName   xml.Name `xml:"supports"`
			Element   string   `xml:"element,attr"`
			Type      string   `xml:"type,attr"`
			Attribute *string  `xml:"attribute,attr,omitempty"`
			Value     *string  `xml:"value,attr,omitempty"`
		}

		ser := &serT{}
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
		(*sL) = append(*sL, ce)
	}

	return nil
}

func (sL *SupportsL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type serT struct {
		XMLName   xml.Name `xml:"supports"`
		Element   string   `xml:"element,attr"`
		Type      string   `xml:"type,attr"`
		Attribute *string  `xml:"attribute,attr,omitempty"`
		Value     *string  `xml:"value,attr,omitempty"`
	}
	type serL []serT
	serl := serL{}

	for _, se := range *sL {
		ser := serT{
			XMLName:   se.XMLName,
			Element:   se.Element,
			Type:      se.Type.String(),
			Attribute: se.Attribute,
			Value:     se.Value,
		}
		serl = append(serl, ser)
		start.Name = se.XMLName
	}

	return e.EncodeElement(serl, start)
}
