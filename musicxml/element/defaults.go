package element

import (
	"encoding/xml"
	"io"
	"reflect"
)

type Defaults struct {
	XMLName       xml.Name        `xml:"defaults"`
	Scaling       *Scaling        `xml:"scaling,omitempty"`
	ConcertScore  *ConcertScore   `xml:"concert-score,omitempty"`
	PageLayout    *PageLayout     `xml:"page-layout,omitempty"`
	SystemLayout  *SystemLayout   `xml:"system-layout,omitempty"`
	StaffLayout   *StaffLayoutL   `xml:"staff-layout,omitempty"`
	Appearance    *Appearance     `xml:"appearance,omitempty"`
	MusicFont     *MusicFont      `xml:"music-font,omitempty"`
	WordFont      *WordFont       `xml:"word-font,omitempty"`
	LyricFont     *LyricFontL     `xml:"lyric-font,omitempty"`
	LyricLanguage *LyricLanguageL `xml:"lyric-language,omitempty"`
}

type DefaultsRaw struct {
	XMLName       xml.Name        `xml:"defaults"`
	Scaling       *Scaling        `xml:"scaling,omitempty"`
	ConcertScore  *ConcertScore   `xml:"concert-score,omitempty"`
	PageLayout    *PageLayout     `xml:"page-layout,omitempty"`
	SystemLayout  *SystemLayout   `xml:"system-layout,omitempty"`
	StaffLayout   *StaffLayoutL   `xml:"staff-layout,omitempty"`
	Appearance    *Appearance     `xml:"appearance,omitempty"`
	MusicFont     *MusicFont      `xml:"music-font,omitempty"`
	WordFont      *WordFont       `xml:"word-font,omitempty"`
	LyricFont     *LyricFontL     `xml:"lyric-font,omitempty"`
	LyricLanguage *LyricLanguageL `xml:"lyric-language,omitempty"`
}

func (dO *Defaults) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	startP := &start
	if reflect.DeepEqual(start, xml.StartElement{}) {
		startP = nil
	}
	dR := &DefaultsRaw{}
	err := d.DecodeElement(dR, startP)
	if err == io.EOF {
		return nil
	}
	if err != nil {
		return err
	}

	*dO = Defaults(*dR)

	return nil
}

func (dO *Defaults) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	dR := DefaultsRaw(*dO)

	start.Name = dO.XMLName

	return e.EncodeElement(dR, start)
}
