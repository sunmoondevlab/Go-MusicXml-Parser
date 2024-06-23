package creator

import (
	"encoding/xml"
	"io"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type CreatorElement struct {
	XMLName xml.Name             `xml:"creator"`
	Type    enum.CreatorTypeEnum `xml:"type,attr"`
	Creator string               `xml:",chardata"`
}

type Creator struct {
	Composer   CreatorElement
	Arranger   CreatorElement
	Lyricist   CreatorElement
	Poet       CreatorElement
	Translator CreatorElement
}

func (c *Creator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		cer := &struct {
			XMLName xml.Name `xml:"creator"`
			Type    string   `xml:"type,attr"`
			Creator string   `xml:",chardata"`
		}{}
		err := d.DecodeElement(cer, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		cet, err := enum.ToCreatorType(cer.Type)
		if err != nil {
			return err
		}
		ce := CreatorElement{
			XMLName: cer.XMLName,
			Type:    *cet,
			Creator: cer.Creator,
		}
		switch *cet {
		case enum.CreatorType.Composer:
			c.Composer = ce
		case enum.CreatorType.Arranger:
			c.Arranger = ce
		case enum.CreatorType.Lyricist:
			c.Lyricist = ce
		case enum.CreatorType.Poet:
			c.Poet = ce
		case enum.CreatorType.Translator:
			c.Translator = ce
		}
	}

	return nil
}
