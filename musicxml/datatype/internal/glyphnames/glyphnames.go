package glyphnames

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// The following embed file is the file in which glyphname and codepoint are extracted from this file "latest/searchindex.json".
// https://github.com/w3c/smufl/blob/gh-pages/latest/searchindex.js

//go:embed support_glyph_names.json
var supportGlyphNamesJson []byte

type SupportSmuflGlyphNamesL []SupportSmuflGlyphName
type SupportSmuflGlyphName struct {
	Name      string `json:"name"`
	CodePoint string `json:"codepoint"`
}

type SupportGlyphNamesJsonData struct {
	Glyphnames SupportSmuflGlyphNamesL `json:"glyph-names"`
}

var SpecGlyphNames *SupportSmuflGlyphNamesL

func init() {
	SpecGlyphNames = getSupportGlyphNames()
}

func getSupportGlyphNames() *SupportSmuflGlyphNamesL {
	var jsonData SupportGlyphNamesJsonData
	json.Unmarshal(supportGlyphNamesJson, &jsonData)

	return &jsonData.Glyphnames
}

func (sgnL *SupportSmuflGlyphNamesL) BinarySearchByName(nm string) (int, bool) {
	if sgnL == nil {
		return 0, false
	}
	l := 0
	r := len(*sgnL)
	if r == 0 {
		return 0, false
	}
	m := (l + r) / 2
	for l < r {
		c := strings.Compare(nm, (*sgnL)[m].Name)
		if c == 0 {
			return m, true
		}

		if c < 0 {
			r = m
		} else {
			l = m + 1
		}
		m = (l + r) / 2
	}

	return m, false
}
