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

type Tenths struct {
	XMLName xml.Name        `xml:"tenths"`
	Tenths  datatype.Tenths `xml:",chardata"`
}

type TenthsRaw struct {
	XMLName xml.Name         `xml:"tenths"`
	Tenths  valuetext.Tenths `xml:",chardata"`
}

func (tO *Tenths) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	tR := &TenthsRaw{}
	err := d.DecodeElement(tR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, tR.XMLName.Space, tR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*tO = Tenths{
		XMLName: tR.XMLName,
		Tenths:  tR.Tenths.ToEntityData(),
	}

	return nil
}

func (tO *Tenths) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	tR := TenthsRaw{
		XMLName: tO.XMLName,
		Tenths:  valuetext.Tenths(tO.Tenths.String()),
	}

	start.Name = tO.XMLName

	return e.EncodeElement(tR, start)
}
