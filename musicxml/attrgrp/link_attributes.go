package attrgrp

import (
	"encoding/xml"
	"reflect"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type LinkAttributes interface {
	SetXMLName(xml.Name)
	SetXLink(string)
	SetHref(string)
	SetType(*enum.XlinkTypeEnum)
	SetRole(*string)
	SetTitle(*string)
	SetShow(*enum.XlinkShowEnum)
	SetActuate(*enum.XlinkActuateEnum)

	GetXMLName() xml.Name
	GetXLink() string
	GetHref() string
	GetType() *enum.XlinkTypeEnum
	GetRole() *string
	GetTitle() *string
	GetShow() *enum.XlinkShowEnum
	GetActuate() *enum.XlinkActuateEnum
}

type LinkAttributesRaw interface {
	GetXMLName() xml.Name
	GetXLink() string
	GetHref() string
	GetType() *string
	GetRole() *string
	GetTitle() *string
	GetShow() *string
	GetActuate() *string

	SetXMLName(xml.Name)
	SetXLink(string)
	SetHref(string)
	SetType(*enum.XlinkTypeEnum)
	SetRole(*string)
	SetTitle(*string)
	SetShow(*enum.XlinkShowEnum)
	SetActuate(*enum.XlinkActuateEnum)
}

func UnmarshalLinkAttributes(d *xml.Decoder, start xml.StartElement, l LinkAttributes, lr LinkAttributesRaw) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	if err := d.DecodeElement(lr, startP); err != nil {
		return err
	}
	var xt *enum.XlinkTypeEnum
	var err error
	if lr.GetType() != nil {
		xt, err = enum.ToXlinkType(*lr.GetType(), lr.GetXMLName().Local)
		if err != nil {
			return err
		}
	}

	var xs *enum.XlinkShowEnum
	if lr.GetShow() != nil {
		xs, err = enum.ToXlinkShow(*lr.GetShow(), lr.GetXMLName().Local)
		if err != nil {
			return err
		}
	}
	var xa *enum.XlinkActuateEnum
	if lr.GetActuate() != nil {
		xa, err = enum.ToXlinkActuate(*lr.GetActuate(), lr.GetXMLName().Local)
		if err != nil {
			return err
		}
	}
	l.SetXMLName(lr.GetXMLName())
	l.SetXLink(lr.GetXLink())
	l.SetHref(lr.GetHref())
	l.SetType(xt)
	l.SetRole(lr.GetRole())
	l.SetTitle(lr.GetTitle())
	l.SetShow(xs)
	l.SetActuate(xa)

	return nil
}
