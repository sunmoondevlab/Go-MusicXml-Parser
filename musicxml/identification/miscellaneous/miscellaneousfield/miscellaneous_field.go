package miscellaneousfield

import (
	"encoding/xml"
	"io"
	"reflect"
)

type MiscellaneousFieldL []MiscellaneousField
type MiscellaneousField struct {
	XMLName xml.Name `xml:"miscellaneous-field"`
	Name    string   `xml:"name,attr"`
	Value   string   `xml:",chardata"`
}

func (mfl *MiscellaneousFieldL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		mfer := &struct {
			XMLName xml.Name `xml:"miscellaneous-field"`
			Name    string   `xml:"name,attr"`
			Value   string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(mfer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		mfe := MiscellaneousField{
			XMLName: mfer.XMLName,
			Name:    mfer.Name,
			Value:   mfer.Value,
		}
		(*mfl) = append(*mfl, mfe)
	}

	return nil
}
