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
	XMLName       xml.Name                      `xml:"identification"`
	Creator       *tCreator.Creator             `xml:"creator,omitempty"`
	Rights        *tRights.RightsL              `xml:"rights,omitempty"`
	Encoding      *tEncoding.Encoding           `xml:"encoding,omitempty"`
	Source        *string                       `xml:"source,omitempty"`
	Relation      *tRelation.RelationL          `xml:"relation,omitempty"`
	Miscellaneous *tMiscellaneous.Miscellaneous `xml:"miscellaneous,omitempty"`
}

func (iO *Identification) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	type irT struct {
		XMLName       xml.Name                      `xml:"identification"`
		Creator       *tCreator.Creator             `xml:"creator,omitempty"`
		Rights        *tRights.RightsL              `xml:"rights,omitempty"`
		Encoding      *tEncoding.Encoding           `xml:"encoding,omitempty"`
		Source        *string                       `xml:"source,omitempty"`
		Relation      *tRelation.RelationL          `xml:"relation,omitempty"`
		Miscellaneous *tMiscellaneous.Miscellaneous `xml:"miscellaneous,omitempty"`
	}
	ir := &irT{}
	err := d.DecodeElement(ir, startP)
	if err != nil {
		return err
	}

	*iO = Identification{
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

func (iO *Identification) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type irT struct {
		XMLName       xml.Name                      `xml:"identification"`
		Creator       *tCreator.Creator             `xml:"creator,omitempty"`
		Rights        *tRights.RightsL              `xml:"rights,omitempty"`
		Encoding      *tEncoding.Encoding           `xml:"encoding,omitempty"`
		Source        *string                       `xml:"source,omitempty"`
		Relation      *tRelation.RelationL          `xml:"relation,omitempty"`
		Miscellaneous *tMiscellaneous.Miscellaneous `xml:"miscellaneous,omitempty"`
	}
	ir := irT(*iO)

	start.Name = iO.XMLName

	return e.EncodeElement(ir, start)
}
