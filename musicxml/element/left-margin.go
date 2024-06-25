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

type LeftMargin struct {
	XMLName    xml.Name        `xml:"left-margin"`
	LeftMargin datatype.Tenths `xml:",chardata"`
}

type LeftMarginRaw struct {
	XMLName    xml.Name         `xml:"left-margin"`
	LeftMargin valuetext.Tenths `xml:",chardata"`
}

func (lmO *LeftMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	lmR := &LeftMarginRaw{}
	err := d.DecodeElement(lmR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, lmR.XMLName.Space, lmR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*lmO = LeftMargin{
		XMLName:    lmR.XMLName,
		LeftMargin: lmR.LeftMargin.ToEntityData(),
	}

	return nil
}

func (lmO *LeftMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	lmR := LeftMarginRaw{
		XMLName:    lmO.XMLName,
		LeftMargin: valuetext.Tenths(lmO.LeftMargin.String()),
	}

	start.Name = lmO.XMLName

	return e.EncodeElement(lmR, start)
}
