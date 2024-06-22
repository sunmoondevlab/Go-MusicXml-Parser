package musicxml

import (
	"encoding/xml"
)

type Work struct {
	XMLName    xml.Name `xml:"work"`
	WorkTitle  string   `xml:"work-title"`
	WorkNumber string   `xml:"work-number"`
	WorkOpus   WorkOpus `xml:"opus"`
}
