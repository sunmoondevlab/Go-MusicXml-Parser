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

type TopMargin struct {
	XMLName   xml.Name        `xml:"top-margin"`
	TopMargin datatype.Tenths `xml:",chardata"`
}

type TopMarginRaw struct {
	XMLName   xml.Name         `xml:"top-margin"`
	TopMargin valuetext.Tenths `xml:",chardata"`
}

func (tmO *TopMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	tmR := &TopMarginRaw{}
	err := d.DecodeElement(tmR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, tmR.XMLName.Space, tmR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}
	*tmO = TopMargin{
		XMLName:   tmR.XMLName,
		TopMargin: tmR.TopMargin.ToEntityData(),
	}

	return nil
}

func (tmO *TopMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tmR := TopMarginRaw{
		XMLName:   tmO.XMLName,
		TopMargin: valuetext.Tenths(tmO.TopMargin.String()),
	}

	start.Name = tmO.XMLName

	return e.EncodeElement(tmR, start)
}
