package element

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/valuetext"
)

type TopSystemDistance struct {
	XMLName           xml.Name        `xml:"top-system-distance"`
	TopSystemDistance datatype.Tenths `xml:",chardata"`
}

type TopSystemDistanceRaw struct {
	XMLName           xml.Name         `xml:"top-system-distance"`
	TopSystemDistance valuetext.Tenths `xml:",chardata"`
}

func (tsdO *TopSystemDistance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	tsdR := &TopSystemDistanceRaw{}
	err := d.DecodeElement(tsdR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*tsdO = TopSystemDistance{
		XMLName:           tsdR.XMLName,
		TopSystemDistance: tsdR.TopSystemDistance.ToEntityData(),
	}

	return nil
}

func (tsdO *TopSystemDistance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tsdR := TopSystemDistanceRaw{
		XMLName:           tsdO.XMLName,
		TopSystemDistance: valuetext.Tenths(tsdO.TopSystemDistance.String()),
	}

	start.Name = tsdO.XMLName

	return e.EncodeElement(tsdR, start)
}
