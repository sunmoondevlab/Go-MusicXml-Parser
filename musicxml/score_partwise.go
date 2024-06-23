package musicxml

import (
	"encoding/xml"

	tIdentification "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification"
	tWork "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work"
)

type ScorePartwise struct {
	XMLName        xml.Name                       `xml:"score-partwise"`
	Work           tWork.Work                     `xml:"work"`
	Identification tIdentification.Identification `xml:"identification"`
}
