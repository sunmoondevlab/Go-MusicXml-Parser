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

type PageHeight struct {
	XMLName    xml.Name        `xml:"page-height"`
	PageHeight datatype.Tenths `xml:",chardata"`
}

type PageHeightRaw struct {
	XMLName    xml.Name         `xml:"page-height"`
	PageHeight valuetext.Tenths `xml:",chardata"`
}

func (phO *PageHeight) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	phR := &PageHeightRaw{}
	err := d.DecodeElement(phR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		if strings.HasPrefix(err.Error(), "invalid value:") {
			return fmt.Errorf(`invalid element value. element:{space:"%s", <%s/>}. error => %s`, phR.XMLName.Space, phR.XMLName.Local, err.Error())
		} else {
			return err
		}
	}

	*phO = PageHeight{
		XMLName:    phR.XMLName,
		PageHeight: phR.PageHeight.ToEntityData(),
	}

	return nil
}

func (phO *PageHeight) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	phR := PageHeightRaw{
		XMLName:    phO.XMLName,
		PageHeight: valuetext.Tenths(phO.PageHeight.String()),
	}

	start.Name = phO.XMLName

	return e.EncodeElement(phR, start)
}
