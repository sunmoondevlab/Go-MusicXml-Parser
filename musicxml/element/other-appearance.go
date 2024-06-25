package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type OtherAppearanceL []OtherAppearance
type OtherAppearance struct {
	XMLName         xml.Name `xml:"other-appearance"`
	Type            *string  `xml:"type,attr,omitempty"`
	OtherAppearance string   `xml:",chardata"`
}

type OtherAppearanceRawL []OtherAppearanceRaw
type OtherAppearanceRaw struct {
	XMLName         xml.Name `xml:"other-appearance"`
	Type            *string  `xml:"type,attr,omitempty"`
	OtherAppearance string   `xml:",chardata"`
}

func (gL *OtherAppearanceL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		gR := &OtherAppearanceRaw{}
		err := d.DecodeElement(gR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		if gR.Type == nil {
			return fmt.Errorf("attribute type is required in <other-appearance/>")
		}

		gE := OtherAppearance(*gR)
		*gL = append(*gL, gE)
	}

	return nil
}

func (gL *OtherAppearanceL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	gRL := OtherAppearanceRawL{}
	for _, gE := range *gL {
		if gE.Type == nil {
			return fmt.Errorf("attribute type is required in <other-appearance/>")
		}

		gR := OtherAppearanceRaw(gE)
		gRL = append(gRL, gR)
		start.Name = gE.XMLName
	}

	return e.EncodeElement(gRL, start)
}
