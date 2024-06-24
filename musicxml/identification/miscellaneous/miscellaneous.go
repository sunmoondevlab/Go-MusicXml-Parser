package miscellaneous

import (
	"encoding/xml"
	"reflect"

	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
)

type Miscellaneous struct {
	XMLName            xml.Name                                 `xml:"miscellaneous"`
	MiscellaneousField *tMiscellaneousfield.MiscellaneousFieldL `xml:"miscellaneous-field,omitempty"`
}

func (mO *Miscellaneous) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	type mrt struct {
		XMLName            xml.Name                                 `xml:"miscellaneous"`
		MiscellaneousField *tMiscellaneousfield.MiscellaneousFieldL `xml:"miscellaneous-field,omitempty"`
	}

	mr := &mrt{}
	err := d.DecodeElement(mr, startP)
	if err != nil {
		return err
	}
	*mO = Miscellaneous(*mr)

	return nil
}

func (mO *Miscellaneous) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type mrt struct {
		XMLName            xml.Name                                 `xml:"miscellaneous"`
		MiscellaneousField *tMiscellaneousfield.MiscellaneousFieldL `xml:"miscellaneous-field,omitempty"`
	}
	mr := mrt(*mO)

	start.Name = mO.XMLName

	return e.EncodeElement(mr, start)
}
