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

type GlyphL []Glyph
type Glyph struct {
	XMLName xml.Name                `xml:"glyph"`
	Type    *enum.GlyphTypeEnum     `xml:"type,attr,omitempty"`
	Glyph   datatype.SmuflGlyphName `xml:",chardata"`
}

type GlyphRawL []GlyphRaw
type GlyphRaw struct {
	XMLName xml.Name                 `xml:"glyph"`
	Type    *attribute.GlyphType     `xml:"type,attr,omitempty"`
	Glyph   valuetext.SmuflGlyphName `xml:",chardata"`
}

func (gL *GlyphL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		gR := &GlyphRaw{}
		err := d.DecodeElement(gR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, gR.XMLName.Space, gR.XMLName.Local, err.Error())
			} else if strings.HasPrefix(err.Error(), "invalid value:") {
				return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, gR.XMLName.Space, gR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		if gR.Type == nil {
			return fmt.Errorf("attribute type is required in <glyph/>")
		}

		gE := Glyph{
			XMLName: gR.XMLName,
			Type:    (*enum.GlyphTypeEnum)(gR.Type),
			Glyph:   gR.Glyph.ToEntityData(),
		}
		*gL = append(*gL, gE)
	}

	return nil
}

func (gL *GlyphL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	gRL := GlyphRawL{}
	for _, gE := range *gL {
		if gE.Type == nil {
			return fmt.Errorf("attribute type is required in <glyph/>")
		}

		gR := GlyphRaw{
			XMLName: gE.XMLName,
			Type:    (*attribute.GlyphType)(gE.Type),
			Glyph:   valuetext.SmuflGlyphName(gE.Glyph.String()),
		}
		gRL = append(gRL, gR)
		start.Name = gE.XMLName
	}

	return e.EncodeElement(gRL, start)
}
