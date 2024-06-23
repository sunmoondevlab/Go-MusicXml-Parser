package miscellaneous

import (
	"encoding/xml"
	"reflect"

	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
)

type Miscellaneous struct {
	XMLName            xml.Name                                `xml:"miscellaneous"`
	MiscellaneousField tMiscellaneousfield.MiscellaneousFieldL `xml:"miscellaneous-field"`
}

func (m *Miscellaneous) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	mr := &struct {
		XMLName            xml.Name                                `xml:"miscellaneous"`
		MiscellaneousField tMiscellaneousfield.MiscellaneousFieldL `xml:"miscellaneous-field"`
	}{}
	err := d.DecodeElement(mr, startP)
	if err != nil {
		return err
	}
	*m = Miscellaneous{
		XMLName:            mr.XMLName,
		MiscellaneousField: mr.MiscellaneousField,
	}

	return nil
}
