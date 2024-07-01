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

type MusicFont struct {
	XMLName    xml.Name             `xml:"music-font"`
	FontFamily *datatype.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *datatype.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *enum.FontStyleEnum  `xml:"font-style,attr,omitempty"`
	FontWeight *enum.FontWeightEnum `xml:"font-weight,attr,omitempty"`
}

type MusicFontRaw struct {
	XMLName    xml.Name              `xml:"music-font"`
	FontFamily *attribute.FontFamily `xml:"font-family,attr,omitempty"`
	FontSize   *attribute.FontSize   `xml:"font-size,attr,omitempty"`
	FontStyle  *attribute.FontStyle  `xml:"font-style,attr,omitempty"`
	FontWeight *attribute.FontWeight `xml:"font-weight,attr,omitempty"`
}

func (mfO *MusicFont) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	mfR := MusicFontRaw{}
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	err := d.DecodeElement(&mfR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid attr:") {
			return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, mfR.XMLName.Space, mfR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*mfO = MusicFont{
		XMLName:    mfR.XMLName,
		FontFamily: (*datatype.FontFamily)(mfR.FontFamily.ToEntityData()),
		FontSize:   (*datatype.FontSize)(mfR.FontSize.ToEntityData()),
		FontStyle:  (*enum.FontStyleEnum)(mfR.FontStyle),
		FontWeight: (*enum.FontWeightEnum)(mfR.FontWeight),
	}

	return nil
}

func (mfO *MusicFont) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	mfR := &MusicFontRaw{
		XMLName:    mfO.XMLName,
		FontFamily: (*attribute.FontFamily)(mfO.FontFamily.StringPtr()),
		FontSize:   (*attribute.FontSize)(mfO.FontSize.StringPtr()),
		FontStyle:  (*attribute.FontStyle)(mfO.FontStyle),
		FontWeight: (*attribute.FontWeight)(mfO.FontWeight),
	}

	start.Name = mfO.XMLName

	return e.EncodeElement(mfR, start)
}
