package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type Creator struct {
	XMLName    xml.Name `xml:"creator"`
	Composer   *CreatorInfo
	Arranger   *CreatorInfo
	Lyricist   *CreatorInfo
	Poet       *CreatorInfo
	Translator *CreatorInfo
}

type CreatorInfoL []CreatorInfo
type CreatorInfo struct {
	XMLName xml.Name             `xml:"creator"`
	Type    enum.CreatorTypeEnum `xml:"type,attr"`
	Creator string               `xml:",chardata"`
}

type CreatorRawL []CreatorRaw
type CreatorRaw struct {
	XMLName xml.Name              `xml:"creator"`
	Type    attribute.CreatorType `xml:"type,attr"`
	Creator string                `xml:",chardata"`
}

func (cO *Creator) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		cR := &CreatorRaw{}
		err := d.DecodeElement(cR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, cR.XMLName.Space, cR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}
		cO.XMLName = cR.XMLName
		ciE := &CreatorInfo{
			XMLName: cR.XMLName,
			Type:    enum.CreatorTypeEnum(cR.Type),
			Creator: cR.Creator,
		}
		switch ciE.Type {
		case enum.CreatorType.Composer:
			cO.Composer = ciE
		case enum.CreatorType.Arranger:
			cO.Arranger = ciE
		case enum.CreatorType.Lyricist:
			cO.Lyricist = ciE
		case enum.CreatorType.Poet:
			cO.Poet = ciE
		case enum.CreatorType.Translator:
			cO.Translator = ciE
		}
	}

	return nil
}

func (cO *Creator) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	ciL := CreatorInfoL{}

	if cO.Composer != nil {
		ciL = append(ciL, *cO.Composer)
	}
	if cO.Arranger != nil {
		ciL = append(ciL, *cO.Arranger)
	}
	if cO.Lyricist != nil {
		ciL = append(ciL, *cO.Lyricist)
	}
	if cO.Poet != nil {
		ciL = append(ciL, *cO.Poet)
	}
	if cO.Translator != nil {
		ciL = append(ciL, *cO.Translator)
	}

	cRL := CreatorRawL{}
	for _, ciE := range ciL {
		start.Name = ciE.XMLName
		cR := CreatorRaw{
			XMLName: ciE.XMLName,
			Type:    attribute.CreatorType(ciE.Type),
			Creator: ciE.Creator,
		}
		cRL = append(cRL, cR)
	}

	return e.EncodeElement(cRL, start)
}
