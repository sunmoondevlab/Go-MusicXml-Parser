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

type Opus struct {
	XMLName xml.Name               `xml:"opus"`
	Xlink   *string                `xml:"xlink,attr,omitempty"`
	Href    *string                `xml:"http://www.w3.org/1999/xlink href,attr"`
	Type    *enum.XlinkTypeEnum    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role    *string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title   *string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show    *enum.XlinkShowEnum    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate *enum.XlinkActuateEnum `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

type OpusRaw struct {
	XMLName xml.Name                `xml:"opus"`
	Xlink   *string                 `xml:"xlink,attr,omitempty"`
	Href    *string                 `xml:"http://www.w3.org/1999/xlink href,attr"`
	Type    *attribute.XlinkType    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role    *string                 `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title   *string                 `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show    *attribute.XlinkShow    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate *attribute.XlinkActuate `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
}

func (oO *Opus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	oR := OpusRaw{}
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	err := d.DecodeElement(&oR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid attr:") {
			return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, oR.XMLName.Space, oR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}
	if oR.Xlink == nil {
		return fmt.Errorf("attribute xmlns:xlink is required in <opus/>")
	}
	if oR.Href == nil {
		return fmt.Errorf("attribute xlink:href is required in <opus/>")
	}

	*oO = Opus{
		XMLName: oR.XMLName,
		Xlink:   oR.Xlink,
		Href:    oR.Href,
		Type:    (*enum.XlinkTypeEnum)(oR.Type),
		Role:    oR.Role,
		Title:   oR.Title,
		Show:    (*enum.XlinkShowEnum)(oR.Show),
		Actuate: (*enum.XlinkActuateEnum)(oR.Actuate),
	}

	return nil
}

func (oO *Opus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if oO.Href == nil {
		return fmt.Errorf("attribute xlink:href is required in <opus/>")
	}

	oR := &OpusRaw{
		XMLName: oO.XMLName,
		Xlink:   nil,
		Href:    oO.Href,
		Type:    (*attribute.XlinkType)(oO.Type),
		Role:    oO.Role,
		Title:   oO.Title,
		Show:    (*attribute.XlinkShow)(oO.Show),
		Actuate: (*attribute.XlinkActuate)(oO.Actuate),
	}

	start.Name = oO.XMLName

	return e.EncodeElement(oR, start)
}
