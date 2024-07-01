package attribute

import (
	"encoding/xml"
	"fmt"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

type XmlLang string

func (xl *XmlLang) UnmarshalXMLAttr(attr xml.Attr) error {
	_, err := datatype.NewXmlLang(attr.Value)
	if err != nil {
		return fmt.Errorf("invalid attr:{Space: %s, Local: %s}: invalid xml:lang, %s", attr.Name.Space, attr.Name.Local, err.Error())
	}
	*xl = (XmlLang)(attr.Value)

	return nil
}

func (xl *XmlLang) ToEntityData() *datatype.XmlLang {
	if xl == nil {
		return nil
	}

	xlV, err := datatype.NewXmlLang((string)(*xl))
	if err != nil {
		return nil
	}
	return xlV
}
