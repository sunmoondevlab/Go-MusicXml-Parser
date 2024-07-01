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

type PageWidth struct {
	XMLName   xml.Name        `xml:"page-width"`
	PageWidth datatype.Tenths `xml:",chardata"`
}

type PageWidthRaw struct {
	XMLName   xml.Name         `xml:"page-width"`
	PageWidth valuetext.Tenths `xml:",chardata"`
}

func (pwO *PageWidth) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	pwR := &PageWidthRaw{}
	err := d.DecodeElement(pwR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, pwR.XMLName.Space, pwR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*pwO = PageWidth{
		XMLName:   pwR.XMLName,
		PageWidth: pwR.PageWidth.ToEntityData(),
	}

	return nil
}

func (pwO *PageWidth) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	pwR := PageWidthRaw{
		XMLName:   pwO.XMLName,
		PageWidth: valuetext.Tenths(pwO.PageWidth.String()),
	}

	start.Name = pwO.XMLName

	return e.EncodeElement(pwR, start)
}
