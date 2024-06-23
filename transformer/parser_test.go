package musicxml

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	tIdentification "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification"
	tCreator "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/creator"
	tEncoding "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
	tMiscellaneous "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous"
	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
	tRelation "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/relation"
	tRights "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/rights"
	tWork "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work"
	tWorkopus "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work/workopus"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestParseMusicXml_work(t *testing.T) {
	sample1 := testutil.ReadAll("testdata/sample1.musicxml")
	type args struct {
		mx string
	}
	tests := []struct {
		name string
		args args
		want musicxml.ScorePartwise
	}{
		{
			name: "test",
			args: args{
				mx: sample1,
			},
			want: musicxml.ScorePartwise{
				XMLName: xml.Name{Local: "score-partwise"},
				Work: tWork.Work{
					XMLName: xml.Name{Local: "work"}, WorkTitle: "無題のスコア", WorkNumber: "Work Number",
					WorkOpus: tWorkopus.WorkOpus{
						XMLName: xml.Name{
							Space: "",
							Local: "opus",
						},
						XLink:   "http://www.w3.org/1999/xlink",
						Href:    "opus/winterreise.musicxml",
						Type:    enum.XlinkType.Simple,
						Role:    "role",
						Title:   "winterreise",
						Show:    enum.XlinkShow.New,
						Actuate: enum.XlinkActuate.None,
					},
				},
				Identification: tIdentification.Identification{
					XMLName: xml.Name{
						Space: "",
						Local: "identification",
					},
					Creator: tCreator.Creator{
						Composer: tCreator.CreatorElement{
							XMLName: xml.Name{
								Space: "",
								Local: "creator",
							},
							Type:    enum.CreatorType.Composer,
							Creator: "作曲者:composs",
						},
						Arranger: tCreator.CreatorElement{},
						Lyricist: tCreator.CreatorElement{
							XMLName: xml.Name{
								Space: "",
								Local: "creator",
							},
							Type:    enum.CreatorType.Lyricist,
							Creator: "作詞者:lilly",
						},
						Poet:       tCreator.CreatorElement{},
						Translator: tCreator.CreatorElement{},
					},
					Rights: []tRights.Rights{
						{
							XMLName: xml.Name{
								Space: "",
								Local: "rights",
							},
							Type:           "",
							CopyRightsType: enum.CopyRightsType.All,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
					},
					Encoding: tEncoding.Encoding{
						XMLName: xml.Name{
							Space: "",
							Local: "encoding",
						},
						Encoder: encoder.EncoderL{
							{
								XMLName: xml.Name{
									Space: "",
									Local: "encoder",
								},
								Type:           "",
								CopyRightsType: enum.CopyRightsType.All,
								Encoder:        "Hogeran A.Fugas",
							},
						},
						EncodingDate:        time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
						Software:            "MuseScore 4.3.2",
						EncodingDescription: "MusicXML example",
						Supports: []supports.Supports{
							{
								XMLName: xml.Name{
									Space: "",
									Local: "supports",
								},
								Element:   "accidental",
								Type:      enum.YesNo.Yes,
								Attribute: "",
								Value:     "",
							},
							{
								XMLName: xml.Name{
									Space: "",
									Local: "supports",
								},
								Element:   "beam",
								Type:      enum.YesNo.Yes,
								Attribute: "",
								Value:     "",
							},
							{
								XMLName: xml.Name{
									Space: "",
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: "new-page",
								Value:     "",
							},
							{
								XMLName: xml.Name{
									Space: "",
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: "new-system",
								Value:     "",
							},
							{
								XMLName: xml.Name{
									Space: "",
									Local: "supports",
								},
								Element:   "stem",
								Type:      enum.YesNo.Yes,
								Attribute: "",
								Value:     "",
							},
						},
					},
					Source: "Based on A. Abbb Abbb of 1000, republished by Dover in 1000.",
					Relation: []tRelation.Relation{
						{
							XMLName: xml.Name{
								Space: "",
								Local: "relation",
							},
							Type:           "",
							CopyRightsType: enum.CopyRightsType.All,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
					},
					Miscellaneous: tMiscellaneous.Miscellaneous{
						XMLName: xml.Name{
							Space: "",
							Local: "miscellaneous",
						},
						MiscellaneousField: []tMiscellaneousfield.MiscellaneousField{
							{
								XMLName: xml.Name{
									Space: "",
									Local: "miscellaneous-field",
								},
								Name:  "difficulty-level",
								Value: "3",
							},
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ParseMusicXml(tt.args.mx)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ParseMusicXml() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}
