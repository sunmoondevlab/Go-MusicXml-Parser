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

type EncodingDate struct {
	XMLName      xml.Name          `xml:"encoding-date"`
	EncodingDate datatype.YyyyMmDd `xml:",chardata"`
}

type EncodingDateRaw struct {
	XMLName      xml.Name           `xml:"encoding-date"`
	EncodingDate valuetext.YyyyMmDd `xml:",chardata"`
}

func (edO *EncodingDate) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	edR := &EncodingDateRaw{}
	err := d.DecodeElement(edR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, edR.XMLName.Space, edR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*edO = EncodingDate{
		XMLName:      edR.XMLName,
		EncodingDate: edR.EncodingDate.ToEntityData(),
	}

	return nil
}

func (edO *EncodingDate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	edR := EncodingDateRaw{
		XMLName:      edO.XMLName,
		EncodingDate: valuetext.YyyyMmDd(edO.EncodingDate.String()),
	}

	start.Name = edO.XMLName

	return e.EncodeElement(edR, start)
}
