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
	Type           string   `xml:"type,attr"`
	CopyRightsType enum.CopyRightsTypeEnum
	CopyRight      string `xml:",chardata"`
}

func (rl *RightsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		rer := &struct {
			XMLName   xml.Name `xml:"rights"`
			Type      string   `xml:"type,attr"`
			CopyRight string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(rer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		crt, _ := enum.ToCopyRightsType(rer.Type)
		re := Rights{
			XMLName:        rer.XMLName,
			Type:           rer.Type,
			CopyRightsType: *crt,
			CopyRight:      rer.CopyRight,
		}
		(*rl) = append(*rl, re)
	}

	return nil
}
