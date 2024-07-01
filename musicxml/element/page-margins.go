package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
	"slices"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element/internal/rawtype/attribute"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type PageMarginsL []PageMargins
type PageMargins struct {
	XMLName      xml.Name             `xml:"page-margins"`
	Type         *enum.MarginTypeEnum `xml:"type,attr,omitempty"`
	LeftMargin   *LeftMargin          `xml:"left-margin,omitempty"`
	RightMargin  *RightMargin         `xml:"right-margin,omitempty"`
	TopMargin    *TopMargin           `xml:"top-margin,omitempty"`
	BottomMargin *BottomMargin        `xml:"bottom-margin,omitempty"`
}

type PageMarginsRawL []PageMarginsRaw
type PageMarginsRaw struct {
	XMLName      xml.Name              `xml:"page-margins"`
	Type         *attribute.MarginType `xml:"type,attr,omitempty"`
	LeftMargin   *LeftMargin           `xml:"left-margin,omitempty"`
	RightMargin  *RightMargin          `xml:"right-margin,omitempty"`
	TopMargin    *TopMargin            `xml:"top-margin,omitempty"`
	BottomMargin *BottomMargin         `xml:"bottom-margin,omitempty"`
}

func (pmL *PageMarginsL) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	for {
		pmR := &PageMarginsRaw{}
		err := d.DecodeElement(pmR, startP)
		if err == io.EOF {
			break
		}
		if err != nil {
			if strings.HasPrefix(err.Error(), "invalid attr:") {
				return fmt.Errorf(`invalid element attr. element:{space:"%s", <%s/>}. error => %s`, pmR.XMLName.Space, pmR.XMLName.Local, err.Error())
			} else {
				return err
			}
		}

		if pmR.LeftMargin == nil {
			return fmt.Errorf("element <left-margin/> is required in <page-margins/>")
		}
		if pmR.RightMargin == nil {
			return fmt.Errorf("element <right-margin/> is required in <page-margins/>")
		}
		if pmR.TopMargin == nil {
			return fmt.Errorf("element <top-margin/> is required in <page-margins/>")
		}
		if pmR.BottomMargin == nil {
			return fmt.Errorf("element <bottom-margin/> is required in <page-margins/>")
		}

		pmE := PageMargins{
			XMLName:      pmR.XMLName,
			Type:         (*enum.MarginTypeEnum)(pmR.Type),
			LeftMargin:   pmR.LeftMargin,
			RightMargin:  pmR.RightMargin,
			TopMargin:    pmR.TopMargin,
			BottomMargin: pmR.BottomMargin,
		}
		(*pmL) = append(*pmL, pmE)
	}

	return nil
}

func (pmL *PageMarginsL) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	pmRL := PageMarginsRawL{}

	for _, pmE := range *pmL {
		if pmE.LeftMargin == nil {
			return fmt.Errorf("element <left-margin/> is required in <page-margins/>")
		}
		if pmE.RightMargin == nil {
			return fmt.Errorf("element <right-margin/> is required in <page-margins/>")
		}
		if pmE.TopMargin == nil {
			return fmt.Errorf("element <top-margin/> is required in <page-margins/>")
		}
		if pmE.BottomMargin == nil {
			return fmt.Errorf("element <bottom-margin/> is required in <page-margins/>")
		}

		pmR := PageMarginsRaw{
			XMLName:      pmE.XMLName,
			Type:         (*attribute.MarginType)(pmE.Type),
			LeftMargin:   pmE.LeftMargin,
			RightMargin:  pmE.RightMargin,
			TopMargin:    pmE.TopMargin,
			BottomMargin: pmE.BottomMargin,
		}
		pmRL = append(pmRL, pmR)
		start.Name = pmE.XMLName
	}

	return e.EncodeElement(pmRL, start)
}

func (pmL *PageMarginsL) ValidatePageMargins() error {
	if pmL == nil || len(*pmL) < 2 {
		return nil
	}
	if len(*pmL) > 2 {
		return fmt.Errorf("<page-margins/> cannot be set to 3 or more")
	}
	mtL := make(map[string]struct{})
	for _, pmE := range *pmL {
		var mt string
		if pmE.Type != nil {
			mt = (*pmE.Type).String()
		}
		mtL[mt] = struct{}{}
	}
	mtKs := make([]string, 0, len(mtL))
	for k := range mtL {
		mtKs = append(mtKs, k)
	}

	if len(mtL) < len(*pmL) {
		return fmt.Errorf("duplicate values cannot be set for attribute:type in <page-margins/>, type=%s", mtKs[0])
	}
	if !(slices.Contains(mtKs, enum.MarginType.Even.String()) && slices.Contains(mtKs, enum.MarginType.Odd.String())) {
		return fmt.Errorf(`if you set multiple <page-margins/>, you must set both "even" and "odd" for attribute:type`)
	}

	return nil
}
