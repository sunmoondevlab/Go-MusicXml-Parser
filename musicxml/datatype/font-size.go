package datatype

import (
	"fmt"
	"strings"

	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

type FontSize struct {
	FontPointSize *decimal.Decimal
	CssFontSize   *enum.CssFontSizeEnum
	fmt.Stringer
}

func (fs *FontSize) String() string {
	if fs == nil {
		return ""
	}
	if (FontSize{}) == (*fs) {
		return ""
	}
	if fs.CssFontSize != nil {
		return fs.CssFontSize.String()
	} else {
		return fs.FontPointSize.String()
	}
}

func (fs *FontSize) StringPtr() *string {
	if fs == nil {
		return nil
	}
	if (FontSize{}) == (*fs) {
		return nil
	}
	if fs.CssFontSize != nil {
		s := fs.CssFontSize.String()
		return &s
	} else {
		s := fs.FontPointSize.String()
		return &s
	}
}

func NewFontSize(value string) (*FontSize, error) {
	if len(strings.TrimSpace(value)) == 0 {
		return nil, nil
	}
	csft, err := enum.ToCssFontSize(value)
	if err == nil {
		return &FontSize{
			CssFontSize: csft,
		}, nil
	}

	d, err := decimal.NewFromString(value)
	if err != nil {
		return nil, fmt.Errorf("invalid font-size. font-size=%s", value)
	}
	return &FontSize{
		FontPointSize: &d,
	}, nil
}
