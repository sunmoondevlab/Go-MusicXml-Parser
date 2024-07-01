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

type BottomMargin struct {
	XMLName      xml.Name        `xml:"bottom-margin"`
	BottomMargin datatype.Tenths `xml:",chardata"`
}

type BottomMarginRaw struct {
	XMLName      xml.Name         `xml:"bottom-margin"`
	BottomMargin valuetext.Tenths `xml:",chardata"`
}

func (bmO *BottomMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	bmR := &BottomMarginRaw{}
	err := d.DecodeElement(bmR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, bmR.XMLName.Space, bmR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*bmO = BottomMargin{
		XMLName:      bmR.XMLName,
		BottomMargin: bmR.BottomMargin.ToEntityData(),
	}

	return nil
}

func (bmO *BottomMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	bmR := BottomMarginRaw{
		XMLName:      bmO.XMLName,
		BottomMargin: valuetext.Tenths(bmO.BottomMargin.String()),
	}

	start.Name = bmO.XMLName

	return e.EncodeElement(bmR, start)
}
