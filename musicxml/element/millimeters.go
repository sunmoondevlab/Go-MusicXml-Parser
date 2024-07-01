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

type Millimeters struct {
	XMLName     xml.Name             `xml:"millimeters"`
	Millimeters datatype.Millimeters `xml:",chardata"`
}

type MillimetersRaw struct {
	XMLName     xml.Name              `xml:"millimeters"`
	Millimeters valuetext.Millimeters `xml:",chardata"`
}

func (mO *Millimeters) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	mR := &MillimetersRaw{}
	err := d.DecodeElement(mR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, mR.XMLName.Space, mR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*mO = Millimeters{
		XMLName:     mR.XMLName,
		Millimeters: mR.Millimeters.ToEntityData(),
	}

	return nil
}

func (mO *Millimeters) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	mR := MillimetersRaw{
		XMLName:     mO.XMLName,
		Millimeters: valuetext.Millimeters(mO.Millimeters.String()),
	}

	start.Name = mO.XMLName

	return e.EncodeElement(mR, start)
}
