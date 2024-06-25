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

type RightMargin struct {
	XMLName     xml.Name        `xml:"right-margin"`
	RightMargin datatype.Tenths `xml:",chardata"`
}

type RightMarginRaw struct {
	XMLName     xml.Name         `xml:"right-margin"`
	RightMargin valuetext.Tenths `xml:",chardata"`
}

func (rmO *RightMargin) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	rmR := &RightMarginRaw{}
	err := d.DecodeElement(rmR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, rmR.XMLName.Space, rmR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*rmO = RightMargin{
		XMLName:     rmR.XMLName,
		RightMargin: rmR.RightMargin.ToEntityData(),
	}

	return nil
}

func (rmO *RightMargin) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	rmR := RightMarginRaw{
		XMLName:     rmO.XMLName,
		RightMargin: valuetext.Tenths(rmO.RightMargin.String()),
	}

	start.Name = rmO.XMLName

	return e.EncodeElement(rmR, start)
}
