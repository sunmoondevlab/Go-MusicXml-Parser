package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
)

type LyricLanguageL []LyricLanguage
type LyricLanguage struct {
	XMLName xml.Name          `xml:"lyric-language"`
	XMLLang *datatype.XmlLang `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Name    *string           `xml:"name,attr,omitempty"`
	Number  *string           `xml:"number,attr,omitempty"`
}

type LyricLanguageRawL []LyricLanguageRaw
type LyricLanguageRaw struct {
	XMLName xml.Name           `xml:"lyric-language"`
	XMLLang *attribute.XmlLang `xml:"http://www.w3.org/XML/1998/namespace lang,attr,omitempty"`
	Name    *string            `xml:"name,attr,omitempty"`
	Number  *string            `xml:"number,attr,omitempty"`
}

func (lfL *LyricLanguageL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start

	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		lfR := LyricLanguageRaw{}
		err := d.DecodeElement(&lfR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, lfR.XMLName.Space, lfR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		if lfR.XMLLang == nil {
			return fmt.Errorf("attribute xml:lang is required in <lyric-language/>")
		}

		lfE := LyricLanguage{
			XMLName: lfR.XMLName,
			XMLLang: lfR.XMLLang.ToEntityData(),
			Name:    lfR.Name,
			Number:  lfR.Number,
		}
		*lfL = append(*lfL, lfE)
	}

	return nil
}

func (lfL *LyricLanguageL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	lfRL := LyricLanguageRawL{}
	for _, lfE := range *lfL {
		if lfE.XMLLang == nil {
			return fmt.Errorf("attribute xml:lang is required in <lyric-language/>")
		}
		lfR := LyricLanguageRaw{
			XMLName: lfE.XMLName,
			XMLLang: (*attribute.XmlLang)(lfE.XMLLang.StringPtr()),
			Name:    lfE.Name,
			Number:  lfE.Number,
		}
		lfRL = append(lfRL, lfR)
		start.Name = lfE.XMLName
	}
	return e.EncodeElement(lfRL, start)
}
