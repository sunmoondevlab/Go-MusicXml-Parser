package languages

import (
	_ "embed"
	"encoding/json"
	"strings"
)

// The embedded file generated from below  https://github.com/sunmoondevlab/get-locales-json-java
// This JSON file contains locale settings supported by Java.
// Some of the Japanese locales have been added manually.

//go:embed support_languages.json
var SupportLanguagesJson []byte

type LanguageExtension struct {
	UExtension *string `json:"u,omitempty"`
	XExtension *string `json:"x,omitempty"`
}

type SupportLanguagesL []SupportLanguage
type SupportLanguage struct {
	Name         string             `json:"name"`
	Language     string             `json:"language"`
	Region       *string            `json:"region,omitempty"`
	Script       *string            `json:"script,omitempty"`
	Variant      *string            `json:"variant,omitempty"`
	Extension    *LanguageExtension `json:"extension,omitempty"`
	ISO3Language *string            `json:"iso3Language"`
	ISO3Region   *string            `json:"iso3region"`
}

type SupportLanguagesJsonData struct {
	Languages SupportLanguagesL `json:"languages"`
}

var SupportLanguages *SupportLanguagesL

func init() {
	SupportLanguages = getSupportLanguages()
}

func getSupportLanguages() *SupportLanguagesL {
	var jsonData SupportLanguagesJsonData
	json.Unmarshal(SupportLanguagesJson, &jsonData)

	return &jsonData.Languages
}

func (slL *SupportLanguagesL) BinarySearchByName(nm string) (int, bool) {
	if slL == nil {
		return 0, false
	}
	l := 0
	r := len(*slL)
	if r == 0 {
		return 0, false
	}
	m := (l + r) / 2
	for l < r {
		c := strings.Compare(nm, (*slL)[m].Name)
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
