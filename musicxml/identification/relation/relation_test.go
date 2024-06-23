package relation

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestRelationL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		rl      *RelationL
		args    args
		wantErr bool
		wantObj RelationL
	}{
		{
			name: "relation invalid decode",
			rl:   &RelationL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<relation>urn:ISBN:0-000-00000-0</relation>
    <relation type="music">urn:ISBN:0-000-00000-0</relation>
    <relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
    <relation type="words">urn:ISBN:0-000-00000-0</relation>
    <relation type="translation">urn:ISBN:0-000-00000-0</relation>
    <relation type="parody">urn:ISBN:0-000-00000-0</relation>
    <relation type="???">urn:ISBN:0-000-00000-0</relation>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "right",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: []Relation{},
		},
		{
			name: "relation",
			rl:   &RelationL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<relation>urn:ISBN:0-000-00000-0</relation>
    <relation type="music">urn:ISBN:0-000-00000-0</relation>
    <relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
    <relation type="words">urn:ISBN:0-000-00000-0</relation>
    <relation type="translation">urn:ISBN:0-000-00000-0</relation>
    <relation type="parody">urn:ISBN:0-000-00000-0</relation>
    <relation type="???">urn:ISBN:0-000-00000-0</relation>`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []Relation{
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "",
					CopyRightsType: enum.CopyRightsType.All,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "music",
					CopyRightsType: enum.CopyRightsType.Music,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "arrangement",
					CopyRightsType: enum.CopyRightsType.Arrangement,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "words",
					CopyRightsType: enum.CopyRightsType.Words,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "translation",
					CopyRightsType: enum.CopyRightsType.Translation,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "parody",
					CopyRightsType: enum.CopyRightsType.Parody,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "relation",
					},
					Type:           "???",
					CopyRightsType: enum.CopyRightsType.Other,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rl.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RelationL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.rl, tt.wantObj); diff != "" {
			t.Errorf("RelationL.UnmarshalXML() value is mismatch (-*tt.rl +tt.wantObj):%s\n", diff)
		}
	}
}
