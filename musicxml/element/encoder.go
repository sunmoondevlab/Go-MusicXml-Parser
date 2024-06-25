package element

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
	CopyrightsType *enum.CopyrightsTypeEnum
	Encoder        string `xml:",chardata"`
}
type EncoderRawL []EncoderRaw
type EncoderRaw struct {
	XMLName xml.Name `xml:"encoder"`
	Type    *string  `xml:"type,attr,omitempty"`
	Encoder string   `xml:",chardata"`
}

func (eL *EncoderL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		eR := &EncoderRaw{}
		err := d.DecodeElement(eR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var crt *enum.CopyrightsTypeEnum
		if eR.Type != nil {
			crt, _ = enum.ToCopyrightsType(*eR.Type)
		}
		eE := Encoder{
			XMLName:        eR.XMLName,
			Type:           eR.Type,
			CopyrightsType: crt,
			Encoder:        eR.Encoder,
		}
		*eL = append(*eL, eE)
	}

	return nil
}

func (eL *EncoderL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	eRL := EncoderRawL{}

	for _, eE := range *eL {
		eR := EncoderRaw{
			XMLName: eE.XMLName,
			Type:    eE.Type,
			Encoder: eE.Encoder,
		}
		eRL = append(eRL, eR)
		start.Name = eE.XMLName
	}

	return e.EncodeElement(eRL, start)
}
