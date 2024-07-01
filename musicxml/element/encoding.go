package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Encoding struct {
	XMLName             xml.Name      `xml:"encoding"`
	Encoder             *EncoderL     `xml:"encoder,omitempty"`
	EncodingDate        *EncodingDate `xml:"encoding-date,omitempty"`
	Software            *string       `xml:"software,omitempty"`
	EncodingDescription *string       `xml:"encoding-description,omitempty"`
	Supports            *SupportsL    `xml:"supports,omitempty"`
}

type EncodingRaw struct {
	XMLName             xml.Name      `xml:"encoding"`
	Encoder             *EncoderL     `xml:"encoder,omitempty"`
	EncodingDate        *EncodingDate `xml:"encoding-date,omitempty"`
	Software            *string       `xml:"software,omitempty"`
	EncodingDescription *string       `xml:"encoding-description,omitempty"`
	Supports            *SupportsL    `xml:"supports,omitempty"`
}

func (eO *Encoding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	eR := &EncodingRaw{}
	err := d.DecodeElement(eR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*eO = Encoding{
		XMLName:             eR.XMLName,
		Encoder:             eR.Encoder,
		EncodingDate:        eR.EncodingDate,
		Software:            eR.Software,
		EncodingDescription: eR.EncodingDescription,
		Supports:            eR.Supports,
	}

	return nil
}

func (eO *Encoding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	eR := EncodingRaw{
		XMLName:             eO.XMLName,
		Encoder:             eO.Encoder,
		EncodingDate:        eO.EncodingDate,
		Software:            eO.Software,
		EncodingDescription: eO.EncodingDescription,
		Supports:            eO.Supports,
	}

	start.Name = eO.XMLName

	return e.EncodeElement(eR, start)
}
