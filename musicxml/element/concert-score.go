package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type ConcertScore struct {
	XMLName xml.Name `xml:"concert-score"`
}

type ConcertScoreRaw struct {
	XMLName xml.Name `xml:"concert-score"`
}

func (csO *ConcertScore) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	csR := &ConcertScoreRaw{}
	err := d.DecodeElement(csR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*csO = ConcertScore(*csR)

	return nil
}

func (csO *ConcertScore) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	csR := ConcertScoreRaw(*csO)

	start.Name = csO.XMLName

	return e.EncodeElement(csR, start)
}
