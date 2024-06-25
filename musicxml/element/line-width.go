package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/valuetext"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type LineWidthL []LineWidth
type LineWidth struct {
	XMLName   xml.Name                `xml:"line-width"`
	Type      *enum.LineWidthTypeEnum `xml:"type,attr,omitempty"`
	LineWidth datatype.Tenths         `xml:",chardata"`
}

type LineWidthRawL []LineWidthRaw
type LineWidthRaw struct {
	XMLName   xml.Name                 `xml:"line-width"`
	Type      *attribute.LineWidthType `xml:"type,attr,omitempty"`
	LineWidth valuetext.Tenths         `xml:",chardata"`
}

func (lwL *LineWidthL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		lwR := &LineWidthRaw{}
		err := d.DecodeElement(lwR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, lwR.XMLName.Space, lwR.XMLName.Local, err.Error())
			} else if strings.HasPrefix(err.Error(), "invalid value:") {
				return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, lwR.XMLName.Space, lwR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		if lwR.Type == nil {
			return fmt.Errorf("attribute type is required in <line-width/>")
		}

		lwE := LineWidth{
			XMLName:   lwR.XMLName,
			Type:      (*enum.LineWidthTypeEnum)(lwR.Type),
			LineWidth: lwR.LineWidth.ToEntityData(),
		}
		*lwL = append(*lwL, lwE)
	}

	return nil
}

func (lwL *LineWidthL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	lwRL := LineWidthRawL{}
	for _, lwE := range *lwL {
		if lwE.Type == nil {
			return fmt.Errorf("attribute type is required in <line-width/>")
		}

		lwR := LineWidthRaw{
			XMLName:   lwE.XMLName,
			Type:      (*attribute.LineWidthType)(lwE.Type),
			LineWidth: valuetext.Tenths(lwE.LineWidth.String()),
		}
		lwRL = append(lwRL, lwR)
		start.Name = lwE.XMLName
	}

	return e.EncodeElement(lwRL, start)
}
