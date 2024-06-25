package element

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type RightsL []Rights
type Rights struct {
	XMLName        xml.Name `xml:"rights"`
	Type           *string  `xml:"type,attr,omitempty"`
	CopyrightsType *enum.CopyrightsTypeEnum
	CopyRight      string `xml:",chardata"`
}
type RightsRawL []RightsRaw
type RightsRaw struct {
	XMLName   xml.Name `xml:"rights"`
	Type      *string  `xml:"type,attr,omitempty"`
	CopyRight string   `xml:",chardata"`
}

func (rL *RightsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		rR := &RightsRaw{}
		err := d.DecodeElement(rR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var crt *enum.CopyrightsTypeEnum
		if rR.Type != nil {
			crt, _ = enum.ToCopyrightsType(*rR.Type)
		}
		rE := Rights{
			XMLName:        rR.XMLName,
			Type:           rR.Type,
			CopyrightsType: crt,
			CopyRight:      rR.CopyRight,
		}
		*rL = append(*rL, rE)
	}

	return nil
}

func (rL *RightsL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	rRL := RightsRawL{}

	for _, rE := range *rL {
		rR := RightsRaw{
			XMLName:   rE.XMLName,
			Type:      rE.Type,
			CopyRight: rE.CopyRight,
		}
		rRL = append(rRL, rR)
		start.Name = rE.XMLName
	}

	return e.EncodeElement(rRL, start)
}
