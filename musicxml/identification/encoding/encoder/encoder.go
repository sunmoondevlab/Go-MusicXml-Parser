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
	Type           string   `xml:"type,attr"`
	CopyRightsType enum.CopyRightsTypeEnum
	Encoder        string `xml:",chardata"`
}

func (el *EncoderL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		eer := &struct {
			XMLName xml.Name `xml:"encoder"`
			Type    string   `xml:"type,attr"`
			Encoder string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(eer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		crt, _ := enum.ToCopyRightsType(eer.Type)
		ee := Encoder{
			XMLName:        eer.XMLName,
			Type:           eer.Type,
			CopyRightsType: *crt,
			Encoder:        eer.Encoder,
		}
		(*el) = append(*el, ee)
	}

	return nil
}
