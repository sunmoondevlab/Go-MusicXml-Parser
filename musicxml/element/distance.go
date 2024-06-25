package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/valuetext"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type DistanceL []Distance
type Distance struct {
	XMLName  xml.Name               `xml:"distance"`
	Type     *enum.DistanceTypeEnum `xml:"type,attr,omitempty"`
	Distance datatype.Tenths        `xml:",chardata"`
}

type DistanceRawL []DistanceRaw
type DistanceRaw struct {
	XMLName  xml.Name                `xml:"distance"`
	Type     *attribute.DistanceType `xml:"type,attr,omitempty"`
	Distance valuetext.Tenths        `xml:",chardata"`
}

func (dL *DistanceL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		dR := &DistanceRaw{}
		err := d.DecodeElement(dR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, dR.XMLName.Space, dR.XMLName.Local, err.Error())
			} else if strings.HasPrefix(err.Error(), "invalid value:") {
				return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, dR.XMLName.Space, dR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		if dR.Type == nil {
			return fmt.Errorf("attribute type is required in <distance/>")
		}

		dE := Distance{
			XMLName:  dR.XMLName,
			Type:     (*enum.DistanceTypeEnum)(dR.Type),
			Distance: dR.Distance.ToEntityData(),
		}
		*dL = append(*dL, dE)
	}

	return nil
}

func (dL *DistanceL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	dRL := DistanceRawL{}
	for _, dE := range *dL {
		if dE.Type == nil {
			return fmt.Errorf("attribute type is required in <distance/>")
		}

		dR := DistanceRaw{
			XMLName:  dE.XMLName,
			Type:     (*attribute.DistanceType)(dE.Type),
			Distance: valuetext.Tenths(dE.Distance.String()),
		}
		dRL = append(dRL, dR)
		start.Name = dE.XMLName
	}

	return e.EncodeElement(dRL, start)
}
