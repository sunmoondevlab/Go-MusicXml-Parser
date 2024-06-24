package rights

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
	CopyRightsType *enum.CopyRightsTypeEnum
	CopyRight      string `xml:",chardata"`
}

func (rL *RightsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		type rerT struct {
			XMLName   xml.Name `xml:"rights"`
			Type      *string  `xml:"type,attr,omitempty"`
			CopyRight string   `xml:",chardata"`
		}
		rer := &rerT{}
		err := d.DecodeElement(rer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		var crt *enum.CopyRightsTypeEnum
		if rer.Type != nil {
			crt, _ = enum.ToCopyRightsType(*rer.Type)
		}
		re := Rights{
			XMLName:        rer.XMLName,
			Type:           rer.Type,
			CopyRightsType: crt,
			CopyRight:      rer.CopyRight,
		}
		(*rL) = append(*rL, re)
	}

	return nil
}

func (rL *RightsL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type rerT struct {
		XMLName   xml.Name `xml:"rights"`
		Type      *string  `xml:"type,attr,omitempty"`
		CopyRight string   `xml:",chardata"`
	}
	type rerL []rerT
	rerl := rerL{}

	for _, re := range *rL {
		rer := rerT{
			XMLName:   re.XMLName,
			Type:      re.Type,
			CopyRight: re.CopyRight,
		}
		rerl = append(rerl, rer)
		start.Name = re.XMLName
	}

	return e.EncodeElement(rerl, start)
}
