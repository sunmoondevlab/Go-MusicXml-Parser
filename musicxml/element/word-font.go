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

type WordFont struct {
	XMLName    xml.Name             `xml:"word-font"`
	FontFamily *datatype.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *datatype.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *enum.FontStyleEnum  `xml:"font-style,attr,omitempty"`
	FontWeight *enum.FontWeightEnum `xml:"font-weight,attr,omitempty"`
}

type WordFontRaw struct {
	XMLName    xml.Name              `xml:"word-font"`
	FontFamily *attribute.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *attribute.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *attribute.FontStyle  `xml:"font-style,attr,omitempty"`
	FontWeight *attribute.FontWeight `xml:"font-weight,attr,omitempty"`
}

func (wfO *WordFont) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	wfR := WordFontRaw{}
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	err := d.DecodeElement(&wfR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid attr:") {
			return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, wfR.XMLName.Space, wfR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*wfO = WordFont{
		XMLName:    wfR.XMLName,
		FontFamily: (*datatype.FontFamily)(wfR.FontFamily.ToEntityData()),
		FontSize:   (*datatype.FontSize)(wfR.FontSize.ToEntityData()),
		FontStyle:  (*enum.FontStyleEnum)(wfR.FontStyle),
		FontWeight: (*enum.FontWeightEnum)(wfR.FontWeight),
	}

	return nil
}

func (wfO *WordFont) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	wfR := &WordFontRaw{
		XMLName:    wfO.XMLName,
		FontFamily: (*attribute.FontFamily)(wfO.FontFamily.StringPtr()),
		FontSize:   (*attribute.FontSize)(wfO.FontSize.StringPtr()),
		FontStyle:  (*attribute.FontStyle)(wfO.FontStyle),
		FontWeight: (*attribute.FontWeight)(wfO.FontWeight),
	}

	start.Name = wfO.XMLName

	return e.EncodeElement(wfR, start)
}
