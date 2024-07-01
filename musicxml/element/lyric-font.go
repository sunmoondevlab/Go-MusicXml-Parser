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

type LyricFontL []LyricFont
type LyricFont struct {
	XMLName    xml.Name             `xml:"lyric-font"`
	FontFamily *datatype.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *datatype.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *enum.FontStyleEnum  `xml:"font-style,attr,omitempty"`
	FontWeight *enum.FontWeightEnum `xml:"font-weight,attr,omitempty"`
	Name       *string              `xml:"name,attr,omitempty"`
	Number     *string              `xml:"number,attr,omitempty"`
}

type LyricFontRawL []LyricFontRaw
type LyricFontRaw struct {
	XMLName    xml.Name              `xml:"lyric-font"`
	FontFamily *attribute.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *attribute.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *attribute.FontStyle  `xml:"font-style,attr,omitempty"`
	FontWeight *attribute.FontWeight `xml:"font-weight,attr,omitempty"`
	Name       *string               `xml:"name,attr,omitempty"`
	Number     *string               `xml:"number,attr,omitempty"`
}

func (lfL *LyricFontL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		lfR := LyricFontRaw{}
		err := d.DecodeElement(&lfR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, lfR.XMLName.Space, lfR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}

		lfE := LyricFont{
			XMLName:    lfR.XMLName,
			FontFamily: (*datatype.FontFamily)(lfR.FontFamily.ToEntityData()),
			FontSize:   (*datatype.FontSize)(lfR.FontSize.ToEntityData()),
			FontStyle:  (*enum.FontStyleEnum)(lfR.FontStyle),
			FontWeight: (*enum.FontWeightEnum)(lfR.FontWeight),
			Name:       lfR.Name,
			Number:     lfR.Number,
		}
		*lfL = append(*lfL, lfE)
	}

	return nil
}

func (lfL *LyricFontL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	lfRL := LyricFontRawL{}
	for _, lfE := range *lfL {
		lfR := LyricFontRaw{
			XMLName:    lfE.XMLName,
			FontFamily: (*attribute.FontFamily)(lfE.FontFamily.StringPtr()),
			FontSize:   (*attribute.FontSize)(lfE.FontSize.StringPtr()),
			FontStyle:  (*attribute.FontStyle)(lfE.FontStyle),
			FontWeight: (*attribute.FontWeight)(lfE.FontWeight),
			Name:       lfE.Name,
			Number:     lfE.Number,
		}
		lfRL = append(lfRL, lfR)
		start.Name = lfE.XMLName
	}
	return e.EncodeElement(lfRL, start)
}
