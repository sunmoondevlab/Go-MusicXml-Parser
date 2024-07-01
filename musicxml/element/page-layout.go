package element

import (
	"encoding/xml"
	"fmt"
	"io"
	"reflect"
)

type PageLayout struct {
	XMLName     xml.Name      `xml:"page-layout"`
	PageHeight  *PageHeight   `xml:"page-height,omitempty"`
	PageWidth   *PageWidth    `xml:"page-width,omitempty"`
	PageMargins *PageMarginsL `xml:"page-margins,omitempty"`
}

type PageLayoutRaw struct {
	XMLName     xml.Name      `xml:"page-layout"`
	PageHeight  *PageHeight   `xml:"page-height,omitempty"`
	PageWidth   *PageWidth    `xml:"page-width,omitempty"`
	PageMargins *PageMarginsL `xml:"page-margins,omitempty"`
}

func (plO *PageLayout) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	plR := &PageLayoutRaw{}

	err := d.DecodeElement(plR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}
	err = plR.ValidatePageHeightWidth()
	if err != nil {
		return err
	}
	err = plR.PageMargins.ValidatePageMargins()
	if err != nil {
		return err
	}
	*plO = PageLayout(*plR)

	return err
}

func (plO *PageLayout) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	start.Name = plO.XMLName

	plR := PageLayoutRaw(*plO)

	err := plR.ValidatePageHeightWidth()
	if err != nil {
		return err
	}
	err = plR.PageMargins.ValidatePageMargins()
	if err != nil {
		return err
	}

	return e.EncodeElement(plR, start)
}

func (plR *PageLayoutRaw) ValidatePageHeightWidth() error {
	if plR == nil {
		return nil
	}
	if plR.PageHeight == nil && plR.PageWidth == nil ||
		plR.PageHeight != nil && plR.PageWidth != nil {
		return nil
	}

	if plR.PageHeight != nil {
		return fmt.Errorf(`only <page-height/> cannot be set for <page-layout/>.
<page-layout/> must have both <page-height/> and <page-width/> set, or both must be unset`)
	}

	return fmt.Errorf(`only <page-width/> cannot be set for <page-layout/>.
	<page-layout/> must have both <page-height/> and <page-width/> set, or both must be unset`)
}
