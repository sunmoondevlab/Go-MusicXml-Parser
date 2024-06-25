package workopus

import (
	"encoding/xml"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/attrgrp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type WorkOpus struct {
	XMLName xml.Name               `xml:"opus"`
	XLink   string                 `xml:"xlink,attr,omitempty"`
	Href    string                 `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Type    *enum.XlinkTypeEnum    `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role    *string                `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title   *string                `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show    *enum.XlinkShowEnum    `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate *enum.XlinkActuateEnum `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	attrgrp.LinkAttributes
}

type WorkOpusRaw struct {
	XMLName xml.Name `xml:"opus"`
	XLink   string   `xml:"xlink,attr,omitempty"`
	Href    string   `xml:"http://www.w3.org/1999/xlink href,attr,omitempty"`
	Type    *string  `xml:"http://www.w3.org/1999/xlink type,attr,omitempty"`
	Role    *string  `xml:"http://www.w3.org/1999/xlink role,attr,omitempty"`
	Title   *string  `xml:"http://www.w3.org/1999/xlink title,attr,omitempty"`
	Show    *string  `xml:"http://www.w3.org/1999/xlink show,attr,omitempty"`
	Actuate *string  `xml:"http://www.w3.org/1999/xlink actuate,attr,omitempty"`
	attrgrp.LinkAttributesRaw
}

func (woo *WorkOpus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	t := WorkOpusRaw{}
	*woo = WorkOpus{}

	err := attrgrp.UnmarshalLinkAttributes(d, start, woo, &t)
	return err
}

func (woo *WorkOpus) MarshalXML(e *xml.Encoder, start xml.StartElement) error {

	wor := &WorkOpusRaw{}

	start.Name = woo.XMLName

	wor.SetXMLName(woo.GetXMLName())
	wor.SetHref(woo.GetHref())
	wor.SetType(woo.GetType())
	wor.SetRole(woo.GetRole())
	wor.SetTitle(woo.GetTitle())
	wor.SetShow(woo.GetShow())
	wor.SetActuate(woo.GetActuate())

	return e.EncodeElement(wor, start)
}

func (l *WorkOpus) GetXMLName() xml.Name {
	return l.XMLName
}

func (l *WorkOpus) GetXLink() string {
	return l.XLink
}

func (l *WorkOpus) GetHref() string {
	return l.Href
}

func (l *WorkOpus) GetType() *enum.XlinkTypeEnum {
	return l.Type
}

func (l *WorkOpus) GetRole() *string {
	return l.Role
}

func (l *WorkOpus) GetTitle() *string {
	return l.Title
}

func (l *WorkOpus) GetShow() *enum.XlinkShowEnum {
	return l.Show
}

func (l *WorkOpus) GetActuate() *enum.XlinkActuateEnum {
	return l.Actuate
}

func (l *WorkOpus) SetXMLName(xn xml.Name) {
	l.XMLName = xn
}

func (l *WorkOpus) SetXLink(xl string) {
	l.XLink = xl
}

func (l *WorkOpus) SetHref(h string) {
	l.Href = h
}

func (l *WorkOpus) SetType(xt *enum.XlinkTypeEnum) {
	if xt == nil {
		l.Type = nil
		return
	}
	l.Type = xt
}

func (l *WorkOpus) SetRole(r *string) {
	if r == nil {
		l.Role = nil
		return
	}
	l.Role = r
}

func (l *WorkOpus) SetTitle(t *string) {
	if t == nil {
		l.Title = nil
		return
	}
	l.Title = t
}

func (l *WorkOpus) SetShow(xs *enum.XlinkShowEnum) {
	if xs == nil {
		l.Show = nil
		return
	}
	l.Show = xs
}

func (l *WorkOpus) SetActuate(xa *enum.XlinkActuateEnum) {
	if xa == nil {
		l.Actuate = nil
		return
	}
	l.Actuate = xa
}

func (l *WorkOpusRaw) GetXMLName() xml.Name {
	return l.XMLName
}

func (l *WorkOpusRaw) GetXLink() string {
	return l.XLink
}

func (l *WorkOpusRaw) GetHref() string {
	return l.Href
}

func (l *WorkOpusRaw) GetType() *string {
	return l.Type
}

func (l *WorkOpusRaw) GetRole() *string {
	return l.Role
}

func (l *WorkOpusRaw) GetTitle() *string {
	return l.Title
}

func (l *WorkOpusRaw) GetShow() *string {
	return l.Show
}

func (l *WorkOpusRaw) GetActuate() *string {
	return l.Actuate
}

func (lr *WorkOpusRaw) SetXMLName(xn xml.Name) {
	lr.XMLName = xn
}

func (lr *WorkOpusRaw) SetXLink(xl string) {
	lr.XLink = xl
}

func (lr *WorkOpusRaw) SetHref(h string) {
	lr.Href = h
}

func (lr *WorkOpusRaw) SetType(xt *enum.XlinkTypeEnum) {
	if xt == nil {
		lr.Type = nil
		return
	}
	lr.Type = (*xt).StringPtr()
}

func (lr *WorkOpusRaw) SetRole(r *string) {
	if r == nil {
		lr.Role = nil
		return
	}
	lr.Role = r
}

func (lr *WorkOpusRaw) SetTitle(t *string) {
	if t == nil {
		lr.Title = nil
		return
	}
	lr.Title = t
}

func (lr *WorkOpusRaw) SetShow(xs *enum.XlinkShowEnum) {
	if xs == nil {
		lr.Show = nil
		return
	}
	lr.Show = (*xs).StringPtr()
}

func (lr *WorkOpusRaw) SetActuate(xa *enum.XlinkActuateEnum) {
	if xa == nil {
		lr.Actuate = nil
		return
	}
	lr.Actuate = (*xa).StringPtr()
}
