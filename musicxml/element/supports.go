package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type SupportsL []Supports
type Supports struct {
	XMLName   xml.Name       `xml:"supports"`
	Element   string         `xml:"element"`
	Type      enum.YesNoEnum `xml:"type"`
	Attribute *string        `xml:"attribute,omitempty"`
	Value     *string        `xml:"value,omitempty"`
}

type SupportsRawL []SupportsRaw
type SupportsRaw struct {
	XMLName   xml.Name        `xml:"supports"`
	Element   string          `xml:"element,attr"`
	Type      attribute.YesNo `xml:"type,attr"`
	Attribute *string         `xml:"attribute,attr,omitempty"`
	Value     *string         `xml:"value,attr,omitempty"`
}

func (sL *SupportsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {

		sR := &SupportsRaw{}
		err := d.DecodeElement(sR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, sR.XMLName.Space, sR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		sE := Supports{
			XMLName:   sR.XMLName,
			Element:   sR.Element,
			Type:      enum.YesNoEnum(sR.Type),
			Attribute: sR.Attribute,
			Value:     sR.Value,
		}
		(*sL) = append(*sL, sE)
	}

	return nil
}

func (sL *SupportsL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	sRL := SupportsRawL{}

	for _, sE := range *sL {
		sR := SupportsRaw{
			XMLName:   sE.XMLName,
			Element:   sE.Element,
			Type:      attribute.YesNo(sE.Type),
			Attribute: sE.Attribute,
			Value:     sE.Value,
		}
		sRL = append(sRL, sR)
		start.Name = sE.XMLName
	}

	return e.EncodeElement(sRL, start)
}
