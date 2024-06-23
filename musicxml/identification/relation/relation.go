package relation

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type RelationL []Relation
type Relation struct {
	XMLName        xml.Name `xml:"relation"`
	Type           string   `xml:"type,attr"`
	CopyRightsType enum.CopyRightsTypeEnum
	Relation       string `xml:",chardata"`
}

func (rl *RelationL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		eer := &struct {
			XMLName  xml.Name `xml:"relation"`
			Type     string   `xml:"type,attr"`
			Relation string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(eer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		crt, _ := enum.ToCopyRightsType(eer.Type)
		re := Relation{
			XMLName:        eer.XMLName,
			Type:           eer.Type,
			CopyRightsType: *crt,
			Relation:       eer.Relation,
		}
		(*rl) = append(*rl, re)
	}

	return nil
}
