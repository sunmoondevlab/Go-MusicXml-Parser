package element

import (
	"encoding/xml"
)

type ScorePartwise struct {
	XMLName        xml.Name        `xml:"score-partwise"`
	Work           *Work           `xml:"work,omitempty"`
	MovementNumber *string         `xml:"movement-number,omitempty"`
	MovementTitle  *string         `xml:"movement-title,omitempty"`
	Identification *Identification `xml:"identification,omitempty"`
}
