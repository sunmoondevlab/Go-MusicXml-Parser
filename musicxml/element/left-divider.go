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

type LeftDivider struct {
	XMLName     xml.Name                  `xml:"left-divider"`
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

type LeftDividerRaw struct {
	XMLName     xml.Name                   `xml:"left-divider"`
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

func (ldO *LeftDivider) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	ldR := LeftDividerRaw{}
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	err := d.DecodeElement(&ldR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid attr:") {
			return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, ldR.XMLName.Space, ldR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*ldO = LeftDivider{
		XMLName:     ldR.XMLName,
		Color:       (*datatype.Color)(ldR.Color.ToEntityData()),
		DefaultX:    (*datatype.Tenths)(ldR.DefaultX.ToEntityData()),
		DefaultY:    (*datatype.Tenths)(ldR.DefaultY.ToEntityData()),
		FontFamily:  (*datatype.FontFamily)(ldR.FontFamily.ToEntityData()),
		FontSize:    (*datatype.FontSize)(ldR.FontSize.ToEntityData()),
		FontStyle:   (*enum.FontStyleEnum)(ldR.FontStyle),
		FontWeight:  (*enum.FontWeightEnum)(ldR.FontWeight),
		Halign:      (*enum.LeftCenterRightEnum)(ldR.Halign),
		PrintObject: (*enum.YesNoEnum)(ldR.PrintObject),
		RelativeX:   (*datatype.Tenths)(ldR.RelativeX.ToEntityData()),
		RelativeY:   (*datatype.Tenths)(ldR.RelativeY.ToEntityData()),
		Valign:      (*enum.ValignEnum)(ldR.Valign),
	}

	return nil
}

func (ldO *LeftDivider) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	ldR := &LeftDividerRaw{
		XMLName:     ldO.XMLName,
		Color:       (*attribute.Color)(ldO.Color.StringPtr()),
		DefaultX:    (*attribute.Tenths)(ldO.DefaultX.StringPtr()),
		DefaultY:    (*attribute.Tenths)(ldO.DefaultY.StringPtr()),
		FontFamily:  (*attribute.FontFamily)(ldO.FontFamily.StringPtr()),
		FontSize:    (*attribute.FontSize)(ldO.FontSize.StringPtr()),
		FontStyle:   (*attribute.FontStyle)(ldO.FontStyle),
		FontWeight:  (*attribute.FontWeight)(ldO.FontWeight),
		Halign:      (*attribute.LeftCenterRight)(ldO.Halign),
		PrintObject: (*attribute.YesNo)(ldO.PrintObject),
		RelativeX:   (*attribute.Tenths)(ldO.RelativeX.StringPtr()),
		RelativeY:   (*attribute.Tenths)(ldO.RelativeY.StringPtr()),
		Valign:      (*attribute.Valign)(ldO.Valign),
	}

	start.Name = ldO.XMLName

	return e.EncodeElement(ldR, start)
}
