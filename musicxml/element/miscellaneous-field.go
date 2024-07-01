package element

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

type MiscellaneousFieldRawL []MiscellaneousFieldRaw
type MiscellaneousFieldRaw struct {
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
		mfR := &MiscellaneousFieldRaw{}
		err := d.DecodeElement(mfR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		mfE := MiscellaneousField(*mfR)
		*mfL = append(*mfL, mfE)
	}

	return nil
}

func (mfL *MiscellaneousFieldL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	mfRL := MiscellaneousFieldRawL{}

	for _, mfE := range *mfL {
		mfR := MiscellaneousFieldRaw(mfE)
		mfRL = append(mfRL, mfR)
		start.Name = mfE.XMLName
	}

	return e.EncodeElement(mfRL, start)
}
