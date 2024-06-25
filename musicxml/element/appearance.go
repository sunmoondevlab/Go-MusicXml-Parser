package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Appearance struct {
	XMLName         xml.Name          `xml:"appearance"`
	LineWidth       *LineWidthL       `xml:"line-width,omitempty"`
	NoteSize        *NoteSizeL        `xml:"note-size,omitempty"`
	Distance        *DistanceL        `xml:"distance,omitempty"`
	Glyph           *GlyphL           `xml:"glyph,omitempty"`
	OtherAppearance *OtherAppearanceL `xml:"other-appearance,omitempty"`
}

type AppearanceRaw struct {
	XMLName         xml.Name          `xml:"appearance"`
	LineWidth       *LineWidthL       `xml:"line-width,omitempty"`
	NoteSize        *NoteSizeL        `xml:"note-size,omitempty"`
	Distance        *DistanceL        `xml:"distance,omitempty"`
	Glyph           *GlyphL           `xml:"glyph,omitempty"`
	OtherAppearance *OtherAppearanceL `xml:"other-appearance,omitempty"`
}

func (aO *Appearance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	aR := &AppearanceRaw{}
	err := d.DecodeElement(aR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*aO = Appearance(*aR)

	return nil
}

func (aO *Appearance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	aR := AppearanceRaw(*aO)

	start.Name = aO.XMLName

	return e.EncodeElement(aR, start)
}
