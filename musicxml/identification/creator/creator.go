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
	XMLName    xml.Name `xml:"creator"`
	Composer   *CreatorElement
	Arranger   *CreatorElement
	Lyricist   *CreatorElement
	Poet       *CreatorElement
	Translator *CreatorElement
}

func (cO *Creator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		type cerT struct {
			XMLName xml.Name `xml:"creator"`
			Type    string   `xml:"type,attr"`
			Creator string   `xml:",chardata"`
		}
		cer := &cerT{}
		err := d.DecodeElement(cer, startP)
		if err == io.EOF {
			break
		}
		cO.XMLName = cer.XMLName
		if err != nil {
			return err
		}
		cet, err := enum.ToCreatorType(cer.Type)
		if err != nil {
			return err
		}
		ce := &CreatorElement{
			XMLName: cer.XMLName,
			Type:    *cet,
			Creator: cer.Creator,
		}
		switch *cet {
		case enum.CreatorType.Composer:
			cO.Composer = ce
		case enum.CreatorType.Arranger:
			cO.Arranger = ce
		case enum.CreatorType.Lyricist:
			cO.Lyricist = ce
		case enum.CreatorType.Poet:
			cO.Poet = ce
		case enum.CreatorType.Translator:
			cO.Translator = ce
		}
	}

	return nil
}

func (cO *Creator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	cl := []*CreatorElement{}

	if cO.Composer != nil {
		cl = append(cl, cO.Composer)
	}
	if cO.Arranger != nil {
		cl = append(cl, cO.Arranger)
	}
	if cO.Lyricist != nil {
		cl = append(cl, cO.Lyricist)
	}
	if cO.Poet != nil {
		cl = append(cl, cO.Poet)
	}
	if cO.Translator != nil {
		cl = append(cl, cO.Translator)
	}

	type cerT struct {
		XMLName xml.Name `xml:"creator"`
		Type    string   `xml:"type,attr"`
		Creator string   `xml:",chardata"`
	}
	type cerL []cerT
	cerl := cerL{}
	for _, ce := range cl {
		start.Name = ce.XMLName
		cer := cerT{
			XMLName: ce.XMLName,
			Type:    ce.Type.String(),
			Creator: ce.Creator,
		}
		cerl = append(cerl, cer)
	}

	return e.EncodeElement(cerl, start)
}
