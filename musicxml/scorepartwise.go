package musicxml

import (
	"encoding/xml"
)

type ScorePartwise struct {
	XMLName xml.Name `xml:"score-partwise"`
	Work    Work     `xml:"work"`
}
