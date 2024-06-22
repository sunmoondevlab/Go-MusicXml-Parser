package musicxml

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml"
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
		want musicxml.ScorePartwise
	}{
		{
			name: "test",
			args: args{
				mx: sample1,
			},
			want: musicxml.ScorePartwise{
				XMLName: xml.Name{Local: "score-partwise"},
				Work: musicxml.Work{
					XMLName: xml.Name{Local: "work"}, WorkTitle: "無題のスコア", WorkNumber: "Work Number",
					WorkOpus: musicxml.WorkOpus{
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
