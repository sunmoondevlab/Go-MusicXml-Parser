package element

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type RelationL []Relation
type Relation struct {
	XMLName        xml.Name `xml:"relation"`
	Type           *string  `xml:"type,attr,omitempty"`
	CopyrightsType *enum.CopyrightsTypeEnum
	Relation       string `xml:",chardata"`
}

type RelationRawL []RelationRaw
type RelationRaw struct {
	XMLName  xml.Name `xml:"relation"`
	Type     *string  `xml:"type,attr,omitempty"`
	Relation string   `xml:",chardata"`
}

func (rL *RelationL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		rR := &RelationRaw{}
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
		rE := Relation{
			XMLName:        rR.XMLName,
			Type:           rR.Type,
			CopyrightsType: crt,
			Relation:       rR.Relation,
		}
		*rL = append(*rL, rE)
	}

	return nil
}

func (rL *RelationL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	rRL := RelationRawL{}

	for _, rE := range *rL {
		rR := RelationRaw{
			XMLName:  rE.XMLName,
			Type:     rE.Type,
			Relation: rE.Relation,
		}
		rRL = append(rRL, rR)
		start.Name = rE.XMLName
	}

	return e.EncodeElement(rRL, start)
}
