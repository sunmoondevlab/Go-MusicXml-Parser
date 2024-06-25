package datatype

import (
	"fmt"
	"strings"

	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype/internal/languages"
)

type XmlLang struct {
	Val *string
	fmt.Stringer
}

func (xl *XmlLang) String() string {
	if xl == nil {
		return ""
	}
	if xl.Val == nil {
		return ""
	}
	return *xl.Val
}

func (xl *XmlLang) StringPtr() *string {
	if xl == nil {
		return nil
	}
	return xl.Val
}

func NewXmlLang(value string) (*XmlLang, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, nil
	}
	_, fnd := languages.SupportLanguages.BinarySearchByName(value)
	if !fnd {
		return nil, fmt.Errorf("can not convert to XmlLang. See the definition in the W3C Extensible Markup Language recommendation. Language names come from ISO 639, with optional country subcodes from ISO 3166. see berow: [https://www.w3.org/TR/xml/#sec-lang-tag] and [https://docs.oracle.com/en/java/javase/22/docs/api/java.base/java/util/Locale.html]. value=%s", value)
	}

	return &XmlLang{
		Val: &value,
	}, nil
}
