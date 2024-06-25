package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type RightDivider struct {
	XMLName     xml.Name                  `xml:"right-divider"`
	Color       *datatype.Color           `xml:"color,attr,omitempty"`
	DefaultX    *datatype.Tenths          `xml:"default-x,attr,omitempty"`
	DefaultY    *datatype.Tenths          `xml:"default-y,attr,omitempty"`
	FontFamily  *datatype.FontFamily      `xml:"font-family,attr,omitempty"`
	FontSize    *datatype.FontSize        `xml:"font-size,attr,omitempty"`
	FontStyle   *enum.FontStyleEnum       `xml:"font-style,attr,omitempty"`
	FontWeight  *enum.FontWeightEnum      `xml:"font-weight,attr,omitempty"`
	Halign      *enum.LeftCenterRightEnum `xml:"halign,attr,omitempty"`
	PrintObject *enum.YesNoEnum           `xml:"print-object,attr,omitempty"`
	RelativeX   *datatype.Tenths          `xml:"relative-x,attr,omitempty"`
	RelativeY   *datatype.Tenths          `xml:"relative-y,attr,omitempty"`
	Valign      *enum.ValignEnum          `xml:"valign,attr,omitempty"`
}

type RightDividerRaw struct {
	XMLName     xml.Name                   `xml:"right-divider"`
	Color       *attribute.Color           `xml:"color,attr,omitempty"`
	DefaultX    *attribute.Tenths          `xml:"default-x,attr,omitempty"`
	DefaultY    *attribute.Tenths          `xml:"default-y,attr,omitempty"`
	FontFamily  *attribute.FontFamily      `xml:"font-family,attr,omitempty"`
	FontSize    *attribute.FontSize        `xml:"font-size,attr,omitempty"`
	FontStyle   *attribute.FontStyle       `xml:"font-style,attr,omitempty"`
	FontWeight  *attribute.FontWeight      `xml:"font-weight,attr,omitempty"`
	Halign      *attribute.LeftCenterRight `xml:"halign,attr,omitempty"`
	PrintObject *attribute.YesNo           `xml:"print-object,attr,omitempty"`
	RelativeX   *attribute.Tenths          `xml:"relative-x,attr,omitempty"`
	RelativeY   *attribute.Tenths          `xml:"relative-y,attr,omitempty"`
	Valign      *attribute.Valign          `xml:"valign,attr,omitempty"`
}

func (rdO *RightDivider) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	rdR := RightDividerRaw{}
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	err := d.DecodeElement(&rdR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid attr:") {
			return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, rdR.XMLName.Space, rdR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*rdO = RightDivider{
		XMLName:     rdR.XMLName,
		Color:       (*datatype.Color)(rdR.Color.ToEntityData()),
		DefaultX:    (*datatype.Tenths)(rdR.DefaultX.ToEntityData()),
		DefaultY:    (*datatype.Tenths)(rdR.DefaultY.ToEntityData()),
		FontFamily:  (*datatype.FontFamily)(rdR.FontFamily.ToEntityData()),
		FontSize:    (*datatype.FontSize)(rdR.FontSize.ToEntityData()),
		FontStyle:   (*enum.FontStyleEnum)(rdR.FontStyle),
		FontWeight:  (*enum.FontWeightEnum)(rdR.FontWeight),
		Halign:      (*enum.LeftCenterRightEnum)(rdR.Halign),
		PrintObject: (*enum.YesNoEnum)(rdR.PrintObject),
		RelativeX:   (*datatype.Tenths)(rdR.RelativeX.ToEntityData()),
		RelativeY:   (*datatype.Tenths)(rdR.RelativeY.ToEntityData()),
		Valign:      (*enum.ValignEnum)(rdR.Valign),
	}

	return nil
}

func (rdO *RightDivider) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	rdR := &RightDividerRaw{
		XMLName:     rdO.XMLName,
		Color:       (*attribute.Color)(rdO.Color.StringPtr()),
		DefaultX:    (*attribute.Tenths)(rdO.DefaultX.StringPtr()),
		DefaultY:    (*attribute.Tenths)(rdO.DefaultY.StringPtr()),
		FontFamily:  (*attribute.FontFamily)(rdO.FontFamily.StringPtr()),
		FontSize:    (*attribute.FontSize)(rdO.FontSize.StringPtr()),
		FontStyle:   (*attribute.FontStyle)(rdO.FontStyle),
		FontWeight:  (*attribute.FontWeight)(rdO.FontWeight),
		Halign:      (*attribute.LeftCenterRight)(rdO.Halign),
		PrintObject: (*attribute.YesNo)(rdO.PrintObject),
		RelativeX:   (*attribute.Tenths)(rdO.RelativeX.StringPtr()),
		RelativeY:   (*attribute.Tenths)(rdO.RelativeY.StringPtr()),
		Valign:      (*attribute.Valign)(rdO.Valign),
	}

	start.Name = rdO.XMLName

	return e.EncodeElement(rdR, start)
}
