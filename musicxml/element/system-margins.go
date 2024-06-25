package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type SystemMargins struct {
	XMLName     xml.Name     `xml:"system-margins"`
	LeftMargin  *LeftMargin  `xml:"left-margin,omitempty"`
	RightMargin *RightMargin `xml:"right-margin,omitempty"`
}

type SystemMarginsRaw struct {
	XMLName     xml.Name     `xml:"system-margins"`
	LeftMargin  *LeftMargin  `xml:"left-margin,omitempty"`
	RightMargin *RightMargin `xml:"right-margin,omitempty"`
}

func (smO *SystemMargins) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	smR := &SystemMarginsRaw{}

	err := d.DecodeElement(smR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	if smR.LeftMargin == nil {
		return fmt.Errorf("element <left-margin/> is required in <system-margins/>")
	}
	if smR.RightMargin == nil {
		return fmt.Errorf("element <right-margin/> is required in <system-margins/>")
	}

	*smO = SystemMargins(*smR)

	return err
}

func (smO *SystemMargins) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if smO.LeftMargin == nil {
		return fmt.Errorf("element <left-margin/> is required in <system-margins/>")
	}
	if smO.RightMargin == nil {
		return fmt.Errorf("element <right-margin/> is required in <system-margins/>")
	}

	start.Name = smO.XMLName

	smR := SystemMarginsRaw(*smO)

	return e.EncodeElement(smR, start)
}
