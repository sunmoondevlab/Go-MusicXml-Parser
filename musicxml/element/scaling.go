package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type Scaling struct {
	XMLName     xml.Name     `xml:"scaling"`
	Millimeters *Millimeters `xml:"millimeters,omitempty"`
	Tenths      *Tenths      `xml:"tenths,omitempty"`
}
type ScalingRaw struct {
	XMLName     xml.Name     `xml:"scaling"`
	Millimeters *Millimeters `xml:"millimeters,omitempty"`
	Tenths      *Tenths      `xml:"tenths,omitempty"`
}

func (sO *Scaling) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	sR := &ScalingRaw{}
	err := d.DecodeElement(sR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	if sR.Millimeters == nil {
		return fmt.Errorf("element <millimeters/> is required in <scaling/>")
	}
	if sR.Tenths == nil {
		return fmt.Errorf("element <tenths/> is required in <scaling/>")
	}

	*sO = Scaling(*sR)

	return nil
}

func (sO *Scaling) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if sO.Millimeters == nil {
		return fmt.Errorf("element <millimeters/> is require in <scaling/>d")
	}
	if sO.Tenths == nil {
		return fmt.Errorf("element <tenths/> is required in <scaling/>")
	}

	sR := ScalingRaw(*sO)

	start.Name = sO.XMLName

	return e.EncodeElement(sR, start)
}
