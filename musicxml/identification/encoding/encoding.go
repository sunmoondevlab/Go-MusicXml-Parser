package musicxml

import (
	"encoding/xml"
	"reflect"
	"time"

	tEncoder "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	tSupports "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
)

type Encoding struct {
	XMLName             xml.Name            `xml:"encoding"`
	Encoder             tEncoder.EncoderL   `xml:"encoder"`
	EncodingDate        time.Time           `xml:"encoding-date"`
	Software            string              `xml:"software"`
	EncodingDescription string              `xml:"encoding-description"`
	Supports            tSupports.SupportsL `xml:"supports"`
}

func (e *Encoding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	er := &struct {
		XMLName             xml.Name            `xml:"encoding"`
		Encoder             tEncoder.EncoderL   `xml:"encoder"`
		EncodingDate        string              `xml:"encoding-date"`
		Software            string              `xml:"software"`
		EncodingDescription string              `xml:"encoding-description"`
		Supports            tSupports.SupportsL `xml:"supports"`
	}{}
	err := d.DecodeElement(er, startP)
	if err != nil {
		return err
	}

	df := "2006-01-02"
	dt, err := time.Parse(df, er.EncodingDate)
	if err != nil {
		return err
	}
	*e = Encoding{
		XMLName:             er.XMLName,
		Encoder:             er.Encoder,
		EncodingDate:        dt,
		Software:            er.Software,
		EncodingDescription: er.EncodingDescription,
		Supports:            er.Supports,
	}

	return nil
}
