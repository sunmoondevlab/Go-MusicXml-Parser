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
	SetType(enum.XlinkTypeEnum)
	SetRole(string)
	SetTitle(string)
	SetShow(enum.XlinkShowEnum)
	SetActuate(enum.XlinkActuateEnum)
}

type LinkAttributesRaw interface {
	GetXMLName() xml.Name
	GetXLink() string
	GetHref() string
	GetType() string
	GetRole() string
	GetTitle() string
	GetShow() string
	GetActuate() string
}

func UnmarshalLinkAttributes(d *xml.Decoder, start xml.StartElement, l LinkAttributes, lr LinkAttributesRaw) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	if err := d.DecodeElement(lr, startP); err != nil {
		return err
	}
	xt, e := enum.ToXlinkType(lr.GetType(), lr.GetXMLName().Local)
	if e != nil {
		return e
	}
	xs, e := enum.ToXlinkShow(lr.GetShow(), lr.GetXMLName().Local)
	if e != nil {
		return e
	}
	xa, e := enum.ToXlinkActuate(lr.GetActuate(), lr.GetXMLName().Local)
	if e != nil {
		return e
	}
	l.SetXMLName(lr.GetXMLName())
	l.SetXLink(lr.GetXLink())
	l.SetHref(lr.GetHref())
	l.SetType(*xt)
	l.SetRole(lr.GetRole())
	l.SetTitle(lr.GetTitle())
	l.SetShow(*xs)
	l.SetActuate(*xa)

	return nil
}
