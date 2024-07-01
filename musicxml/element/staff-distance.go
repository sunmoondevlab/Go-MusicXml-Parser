package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/valuetext"
)

type StaffDistance struct {
	XMLName       xml.Name        `xml:"staff-distance"`
	StaffDistance datatype.Tenths `xml:",chardata"`
}

type StaffDistanceRaw struct {
	XMLName       xml.Name         `xml:"staff-distance"`
	StaffDistance valuetext.Tenths `xml:",chardata"`
}

func (sdO *StaffDistance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	sdR := &StaffDistanceRaw{}
	err := d.DecodeElement(sdR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, sdR.XMLName.Space, sdR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}
	*sdO = StaffDistance{
		XMLName:       sdR.XMLName,
		StaffDistance: sdR.StaffDistance.ToEntityData(),
	}

	return nil
}

func (sdO *StaffDistance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	sdR := StaffDistanceRaw{
		XMLName:       sdO.XMLName,
		StaffDistance: valuetext.Tenths(sdO.StaffDistance.String()),
	}

	start.Name = sdO.XMLName

	return e.EncodeElement(sdR, start)
}
