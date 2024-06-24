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
	Type           *string  `xml:"type,attr,omitempty"`
	CopyRightsType *enum.CopyRightsTypeEnum
	Relation       string `xml:",chardata"`
}

func (rL *RelationL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		type rerT struct {
			XMLName  xml.Name `xml:"relation"`
			Type     *string  `xml:"type,attr,omitempty"`
			Relation string   `xml:",chardata"`
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
		re := Relation{
			XMLName:        rer.XMLName,
			Type:           rer.Type,
			CopyRightsType: crt,
			Relation:       rer.Relation,
		}
		(*rL) = append(*rL, re)
	}

	return nil
}

func (rL *RelationL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type rerT struct {
		XMLName  xml.Name `xml:"relation"`
		Type     *string  `xml:"type,attr,omitempty"`
		Relation string   `xml:",chardata"`
	}
	type rerL []rerT
	rerl := rerL{}

	for _, re := range *rL {
		rer := rerT{
			XMLName:  re.XMLName,
			Type:     re.Type,
			Relation: re.Relation,
		}
		rerl = append(rerl, rer)
		start.Name = re.XMLName
	}

	return e.EncodeElement(rerl, start)
}
