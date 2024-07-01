package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type SystemDividers struct {
	XMLName      xml.Name      `xml:"system-dividers"`
	LeftDivider  *LeftDivider  `xml:"left-divider,omitempty"`
	RightDivider *RightDivider `xml:"right-divider,omitempty"`
}

type SystemDividersRaw struct {
	XMLName      xml.Name      `xml:"system-dividers"`
	LeftDivider  *LeftDivider  `xml:"left-divider,omitempty"`
	RightDivider *RightDivider `xml:"right-divider,omitempty"`
}

func (sdO *SystemDividers) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	sdR := &SystemDividersRaw{}

	err := d.DecodeElement(sdR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	if sdR.LeftDivider == nil {
		return fmt.Errorf("element <left-divider/> is required in <system-dividers/>")
	}
	if sdR.RightDivider == nil {
		return fmt.Errorf("element <right-divider/> is required in <system-dividers/>")
	}

	*sdO = SystemDividers(*sdR)

	return err
}

func (sdO *SystemDividers) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	if sdO.LeftDivider == nil {
		return fmt.Errorf("element <left-divider/> is required in <system-dividers/>")
	}
	if sdO.RightDivider == nil {
		return fmt.Errorf("element <right-divider/> is required in <system-dividers/>")
	}

	start.Name = sdO.XMLName

	sdR := SystemDividersRaw(*sdO)

	return e.EncodeElement(sdR, start)
}
