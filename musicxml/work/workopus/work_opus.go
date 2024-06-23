package workopus

import (
	"encoding/xml"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/attrgrp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type WorkOpus struct {
	XMLName xml.Name              `xml:"opus"`
	XLink   string                `xml:"xlink,attr"`
	Href    string                `xml:"href,attr"`
	Type    enum.XlinkTypeEnum    `xml:"type,attr"`
	Role    string                `xml:"role,attr"`
	Title   string                `xml:"title,attr"`
	Show    enum.XlinkShowEnum    `xml:"show,attr"`
	Actuate enum.XlinkActuateEnum `xml:"actuate,attr"`
	attrgrp.LinkAttributes
}

type WorkOpusRaw struct {
	XMLName xml.Name `xml:"opus"`
	XLink   string   `xml:"xlink,attr"`
	Href    string   `xml:"href,attr"`
	Type    string   `xml:"type,attr"`
	Role    string   `xml:"role,attr"`
	Title   string   `xml:"title,attr"`
	Show    string   `xml:"show,attr"`
	Actuate string   `xml:"actuate,attr"`
	attrgrp.LinkAttributesRaw
}

func (wo *WorkOpus) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	t := WorkOpusRaw{}
	*wo = WorkOpus{}

	err := attrgrp.UnmarshalLinkAttributes(d, start, wo, &t)
	return err
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

func (l *WorkOpus) SetType(xt enum.XlinkTypeEnum) {
	l.Type = xt
}

func (l *WorkOpus) SetRole(r string) {
	l.Role = r
}

func (l *WorkOpus) SetTitle(t string) {
	l.Title = t
}

func (l *WorkOpus) SetShow(xs enum.XlinkShowEnum) {
	l.Show = xs
}

func (l *WorkOpus) SetActuate(xa enum.XlinkActuateEnum) {
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

func (l *WorkOpusRaw) GetType() string {
	return l.Type
}

func (l *WorkOpusRaw) GetRole() string {
	return l.Role
}

func (l *WorkOpusRaw) GetTitle() string {
	return l.Title
}

func (l *WorkOpusRaw) GetShow() string {
	return l.Show
}

func (l *WorkOpusRaw) GetActuate() string {
	return l.Actuate
}
