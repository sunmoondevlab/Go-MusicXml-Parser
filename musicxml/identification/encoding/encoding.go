package musicxml

import (
	"encoding/xml"
	"reflect"
	"time"

	tEncoder "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	tSupports "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
)

type Encoding struct {
	XMLName             xml.Name             `xml:"encoding"`
	Encoder             *tEncoder.EncoderL   `xml:"encoder"`
	EncodingDate        *time.Time           `xml:"encoding-date"`
	Software            *string              `xml:"software"`
	EncodingDescription *string              `xml:"encoding-description"`
	Supports            *tSupports.SupportsL `xml:"supports"`
}

func (eO *Encoding) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	er := &struct {
		XMLName             xml.Name             `xml:"encoding"`
		Encoder             *tEncoder.EncoderL   `xml:"encoder"`
		EncodingDate        *string              `xml:"encoding-date"`
		Software            *string              `xml:"software"`
		EncodingDescription *string              `xml:"encoding-description"`
		Supports            *tSupports.SupportsL `xml:"supports"`
	}{}
	err := d.DecodeElement(er, startP)
	if err != nil {
		return err
	}

	df := "2006-01-02"
	var dt *time.Time
	if er.EncodingDate != nil {
		dtv, err := time.Parse(df, *er.EncodingDate)
		if err != nil {
			return err
		}
		dt = &dtv
	}
	*eO = Encoding{
		XMLName:             er.XMLName,
		Encoder:             er.Encoder,
		EncodingDate:        dt,
		Software:            er.Software,
		EncodingDescription: er.EncodingDescription,
		Supports:            er.Supports,
	}

	return nil
}

func (eO *Encoding) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	df := "2006-01-02"
	var ds *string
	if eO.EncodingDate != nil {
		dsv := eO.EncodingDate.Format(df)
		ds = &dsv
	}
	type erT struct {
		XMLName             xml.Name             `xml:"encoding"`
		Encoder             *tEncoder.EncoderL   `xml:"encoder"`
		EncodingDate        *string              `xml:"encoding-date"`
		Software            *string              `xml:"software"`
		EncodingDescription *string              `xml:"encoding-description"`
		Supports            *tSupports.SupportsL `xml:"supports"`
	}

	er := erT{
		XMLName:             eO.XMLName,
		Encoder:             eO.Encoder,
		EncodingDate:        ds,
		Software:            eO.Software,
		EncodingDescription: eO.EncodingDescription,
		Supports:            eO.Supports,
	}

	start.Name = eO.XMLName

	return e.EncodeElement(er, start)
}
