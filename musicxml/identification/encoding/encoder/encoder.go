package encoder

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type EncoderL []Encoder
type Encoder struct {
	XMLName        xml.Name `xml:"encoder"`
	Type           *string  `xml:"type,attr,omitempty"`
	CopyRightsType *enum.CopyRightsTypeEnum
	Encoder        string `xml:",chardata"`
}

func (eL *EncoderL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		type eerT struct {
			XMLName xml.Name `xml:"encoder"`
			Type    *string  `xml:"type,attr,omitempty"`
			Encoder string   `xml:",chardata"`
		}

		eer := &eerT{}
		err := d.DecodeElement(eer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var crt *enum.CopyRightsTypeEnum
		if eer.Type != nil {
			crt, _ = enum.ToCopyRightsType(*eer.Type)
		}
		ee := Encoder{
			XMLName:        eer.XMLName,
			Type:           eer.Type,
			CopyRightsType: crt,
			Encoder:        eer.Encoder,
		}
		(*eL) = append(*eL, ee)
	}

	return nil
}

func (eL *EncoderL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type eerT struct {
		XMLName xml.Name `xml:"encoder"`
		Type    *string  `xml:"type,attr,omitempty"`
		Encoder string   `xml:",chardata"`
	}
	type eerL []eerT
	eerl := eerL{}

	for _, ee := range *eL {
		eer := eerT{
			XMLName: ee.XMLName,
			Type:    ee.Type,
			Encoder: ee.Encoder,
		}
		eerl = append(eerl, eer)
		start.Name = ee.XMLName
	}

	return e.EncodeElement(eerl, start)
}
