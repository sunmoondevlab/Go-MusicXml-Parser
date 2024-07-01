package element

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
)

type StaffLayoutL []StaffLayout
type StaffLayout struct {
	XMLName       xml.Name              `xml:"staff-layout"`
	Number        *datatype.StaffNumber `xml:"number,attr,omitempty"`
	StaffDistance *StaffDistance        `xml:"staff-distance"`
}
type StaffLayoutRawL []StaffLayoutRaw
type StaffLayoutRaw struct {
	XMLName       xml.Name               `xml:"staff-layout"`
	Number        *attribute.StaffNumber `xml:"number,attr,omitempty"`
	StaffDistance *StaffDistance         `xml:"staff-distance"`
}

func (slL *StaffLayoutL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		slR := &StaffLayoutRaw{}
		err := d.DecodeElement(slR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		slE := StaffLayout{
			XMLName:       slR.XMLName,
			Number:        slR.Number.ToEntityData(),
			StaffDistance: slR.StaffDistance,
		}
		*slL = append(*slL, slE)
	}

	return nil
}

func (slL *StaffLayoutL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	slRL := StaffLayoutRawL{}

	for _, slE := range *slL {
		slR := StaffLayoutRaw{
			XMLName:       slE.XMLName,
			Number:        (*attribute.StaffNumber)(slE.Number.StringPtr()),
			StaffDistance: slE.StaffDistance,
		}
		slRL = append(slRL, slR)
		start.Name = slE.XMLName
	}

	return e.EncodeElement(slRL, start)
}
