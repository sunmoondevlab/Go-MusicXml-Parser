package musicxml

import (
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/element"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
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
		want element.ScorePartwise
	}{
		{
			name: "test",
			args: args{
				mx: sample1,
			},
			want: element.ScorePartwise{
				XMLName: xml.Name{Local: "score-partwise"},
				Work: &element.Work{
					XMLName: xml.Name{Local: "work"}, WorkTitle: testutil.ToStringPtr("無題のスコア"), WorkNumber: testutil.ToStringPtr("Work Number"),
					Opus: &element.Opus{
						XMLName: xml.Name{
							Local: "opus",
						},
						Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
						Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
						Type:    &enum.XlinkType.Simple,
						Role:    testutil.ToStringPtr("role"),
						Title:   testutil.ToStringPtr("winterreise"),
						Show:    &enum.XlinkShow.New,
						Actuate: &enum.XlinkActuate.None,
					},
				},
				MovementNumber: testutil.ToStringPtr("Move 1"),
				MovementTitle:  testutil.ToStringPtr("Move Move"),
				Identification: &element.Identification{
					XMLName: xml.Name{
						Local: "identification",
					},
					Creator: &element.Creator{
						XMLName: xml.Name{
							Local: "creator",
						},
						Composer: &element.CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Composer,
							Creator: "作曲者:composs",
						},
						Lyricist: &element.CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Lyricist,
							Creator: "作詞者:lilly",
						},
					},
					Rights: &element.RightsL{
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							CopyRight: "Copyright © 2024 Hogeran A.Fugas",
						},
					},
					Encoding: &element.Encoding{
						XMLName: xml.Name{
							Local: "encoding",
						},
						Encoder: &element.EncoderL{
							{
								XMLName: xml.Name{
									Local: "encoder",
								},
								Encoder: "Hogeran A.Fugas",
							},
						},
						EncodingDate: &element.EncodingDate{
							XMLName: xml.Name{
								Local: "encoding-date",
							},
							EncodingDate: datatype.YyyyMmDd{
								Val: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
							},
						},
						Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
						EncodingDescription: testutil.ToStringPtr("MusicXML example"),
						Supports: &element.SupportsL{
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "accidental",
								Type:    enum.YesNo.Yes,
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "beam",
								Type:    enum.YesNo.Yes,
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: testutil.ToStringPtr("new-page"),
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: testutil.ToStringPtr("new-system"),
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "stem",
								Type:    enum.YesNo.Yes,
							},
						},
					},
					Source: testutil.ToStringPtr("Based on A. Abbb Abbb of 1000, republished by Dover in 1000."),
					Relation: &element.RelationL{
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Relation: "urn:ISBN:0-000-00000-0",
						},
					},
					Miscellaneous: &element.Miscellaneous{
						XMLName: xml.Name{
							Local: "miscellaneous",
						},
						MiscellaneousField: &element.MiscellaneousFieldL{
							{
								XMLName: xml.Name{
									Local: "miscellaneous-field",
								},
								Name:               "difficulty-level",
								MiscellaneousField: "3",
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
				t.Errorf("ParseMusicXml() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
