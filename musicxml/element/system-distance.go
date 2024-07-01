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

type SystemDistance struct {
	XMLName        xml.Name        `xml:"system-distance"`
	SystemDistance datatype.Tenths `xml:",chardata"`
}

type SystemDistanceRaw struct {
	XMLName        xml.Name         `xml:"system-distance"`
	SystemDistance valuetext.Tenths `xml:",chardata"`
}

func (sdO *SystemDistance) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	sdR := &SystemDistanceRaw{}
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
	*sdO = SystemDistance{
		XMLName:        sdR.XMLName,
		SystemDistance: sdR.SystemDistance.ToEntityData(),
	}

	return nil
}

func (sdO *SystemDistance) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	sdR := SystemDistanceRaw{
		XMLName:        sdO.XMLName,
		SystemDistance: valuetext.Tenths(sdO.SystemDistance.String()),
	}

	start.Name = sdO.XMLName

	return e.EncodeElement(sdR, start)
}
