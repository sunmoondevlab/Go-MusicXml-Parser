package miscellaneousfield

import (
	"encoding/xml"
	"io"
	"reflect"
)

type MiscellaneousFieldL []MiscellaneousField
type MiscellaneousField struct {
	XMLName            xml.Name `xml:"miscellaneous-field"`
	Name               string   `xml:"name,attr"`
	MiscellaneousField string   `xml:",chardata"`
}

func (mfL *MiscellaneousFieldL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		mfer := &struct {
			XMLName            xml.Name `xml:"miscellaneous-field"`
			Name               string   `xml:"name,attr"`
			MiscellaneousField string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(mfer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		mfe := MiscellaneousField{
			XMLName:            mfer.XMLName,
			Name:               mfer.Name,
			MiscellaneousField: mfer.MiscellaneousField,
		}
		(*mfL) = append(*mfL, mfe)
	}

	return nil
}

func (mfL *MiscellaneousFieldL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	type mferT struct {
		XMLName            xml.Name `xml:"miscellaneous-field"`
		Name               string   `xml:"name,attr"`
		MiscellaneousField string   `xml:",chardata"`
	}
	type mferL []mferT
	mferl := mferL{}

	for _, mfe := range *mfL {
		mfer := mferT(mfe)
		mferl = append(mferl, mfer)
		start.Name = mfe.XMLName
	}

	return e.EncodeElement(mferl, start)
}
