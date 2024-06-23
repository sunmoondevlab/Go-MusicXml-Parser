package identification

import (
	"encoding/xml"
	"reflect"

	tCreator "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/creator"
	tEncoding "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding"
	tMiscellaneous "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous"
	tRelation "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/relation"
	tRights "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/rights"
)

type Identification struct {
	XMLName       xml.Name                     `xml:"identification"`
	Creator       tCreator.Creator             `xml:"creator"`
	Rights        tRights.RightsL              `xml:"rights"`
	Encoding      tEncoding.Encoding           `xml:"encoding"`
	Source        string                       `xml:"source"`
	Relation      tRelation.RelationL          `xml:"relation"`
	Miscellaneous tMiscellaneous.Miscellaneous `xml:"miscellaneous"`
}

func (i *Identification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	ir := &struct {
		XMLName       xml.Name                     `xml:"identification"`
		Creator       tCreator.Creator             `xml:"creator"`
		Rights        tRights.RightsL              `xml:"rights"`
		Encoding      tEncoding.Encoding           `xml:"encoding"`
		Source        string                       `xml:"source"`
		Relation      tRelation.RelationL          `xml:"relation"`
		Miscellaneous tMiscellaneous.Miscellaneous `xml:"miscellaneous"`
	}{}
	err := d.DecodeElement(ir, startP)
	if err != nil {
		return err
	}

	*i = Identification{
		XMLName:       ir.XMLName,
		Creator:       ir.Creator,
		Rights:        ir.Rights,
		Encoding:      ir.Encoding,
		Source:        ir.Source,
		Relation:      ir.Relation,
		Miscellaneous: ir.Miscellaneous,
	}

	return nil
}
