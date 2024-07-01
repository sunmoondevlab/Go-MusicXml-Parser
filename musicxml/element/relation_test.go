package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestRelationL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		rL      *RelationL
		args    args
		wantErr bool
		wantObj RelationL
	}{
		{
			name: "relation invalid decode",
			rL:   &RelationL{},
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
						Local: "relatio",
					},
				},
			},
			wantErr: true,
			wantObj: RelationL{},
		},
		{
			name: "relation",
			rL:   &RelationL{},
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
			wantObj: RelationL{
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Relation: "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("music"),
					CopyrightsType: &enum.CopyrightsType.Music,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("arrangement"),
					CopyrightsType: &enum.CopyrightsType.Arrangement,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("words"),
					CopyrightsType: &enum.CopyrightsType.Words,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("translation"),
					CopyrightsType: &enum.CopyrightsType.Translation,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("parody"),
					CopyrightsType: &enum.CopyrightsType.Parody,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
				{
					XMLName: xml.Name{
						Local: "relation",
					},
					Type:           testutil.ToStringPtr("???"),
					CopyrightsType: &enum.CopyrightsType.Other,
					Relation:       "urn:ISBN:0-000-00000-0",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RelationL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.rL, tt.wantObj); diff != "" {
			t.Errorf("RelationL.UnmarshalXML() value is mismatch (-rL, +wantObj):%s\n", diff)
		}
	}
}

func TestRelationL_MarshalXML(t *testing.T) {
	type args struct {
		rL *RelationL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "relation",
			args: args{
				rL: &RelationL{
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Relation: "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("music"),
						CopyrightsType: &enum.CopyrightsType.Music,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyrightsType: &enum.CopyrightsType.Arrangement,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyrightsType: &enum.CopyrightsType.Words,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyrightsType: &enum.CopyrightsType.Translation,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyrightsType: &enum.CopyrightsType.Parody,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyrightsType: &enum.CopyrightsType.Other,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
				},
			},
			wantErr: false,
			wantXml: `<relation>urn:ISBN:0-000-00000-0</relation>
<relation type="music">urn:ISBN:0-000-00000-0</relation>
<relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
<relation type="words">urn:ISBN:0-000-00000-0</relation>
<relation type="translation">urn:ISBN:0-000-00000-0</relation>
<relation type="parody">urn:ISBN:0-000-00000-0</relation>
<relation type="???">urn:ISBN:0-000-00000-0</relation>`,
		},
		{
			name: "relation empty",
			args: args{
				rL: &RelationL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.rL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("RelationL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("RelationL.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
