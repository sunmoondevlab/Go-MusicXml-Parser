package encoder

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestEncoderL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		eL      *EncoderL
		args    args
		wantErr bool
		wantObj EncoderL
	}{
		{
			name: "encoder invalid decode",
			eL:   &EncoderL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoder>Hogeran A.Fugas</encoder>
      <encoder type="music">Hogeran A.Fugas</encoder>
      <encoder type="arrangement">Hogeran A.Fugas</encoder>
      <encoder type="words">Hogeran A.Fugas</encoder>
      <encoder type="translation">Hogeran A.Fugas</encoder>
      <encoder type="parody">Hogeran A.Fugas</encoder>
      <encoder type="???">Hogeran A.Fugas</encoder>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "right",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: []Encoder{},
		},
		{
			name: "encoder",
			eL:   &EncoderL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoder>Hogeran A.Fugas</encoder>
      <encoder type="music">Hogeran A.Fugas</encoder>
      <encoder type="arrangement">Hogeran A.Fugas</encoder>
      <encoder type="words">Hogeran A.Fugas</encoder>
      <encoder type="translation">Hogeran A.Fugas</encoder>
      <encoder type="parody">Hogeran A.Fugas</encoder>
      <encoder type="???">Hogeran A.Fugas</encoder>`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []Encoder{
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Encoder: "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("music"),
					CopyRightsType: &enum.CopyRightsType.Music,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("arrangement"),
					CopyRightsType: &enum.CopyRightsType.Arrangement,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("words"),
					CopyRightsType: &enum.CopyRightsType.Words,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("translation"),
					CopyRightsType: &enum.CopyRightsType.Translation,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("parody"),
					CopyRightsType: &enum.CopyRightsType.Parody,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "encoder",
					},
					Type:           testutil.ToStringPtr("???"),
					CopyRightsType: &enum.CopyRightsType.Other,
					Encoder:        "Hogeran A.Fugas",
				},
			},
		},
		{
			name: "encoder empty",
			eL:   &EncoderL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []Encoder{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.eL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("EncoderL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.eL, tt.wantObj); diff != "" {
			t.Errorf("EncoderL.UnmarshalXML() value is mismatch (-*tt.el +tt.wantObj):%s\n", diff)
		}
	}
}

func TestEncoderL_MarshalXML(t *testing.T) {
	type args struct {
		eL *EncoderL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "encoder",
			args: args{
				eL: &EncoderL{
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Encoder: "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("music"),
						CopyRightsType: &enum.CopyRightsType.Music,
						Encoder:        "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyRightsType: &enum.CopyRightsType.Arrangement,
						Encoder:        "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyRightsType: &enum.CopyRightsType.Words,
						Encoder:        "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyRightsType: &enum.CopyRightsType.Translation,
						Encoder:        "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyRightsType: &enum.CopyRightsType.Parody,
						Encoder:        "Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyRightsType: &enum.CopyRightsType.Other,
						Encoder:        "Hogeran A.Fugas",
					},
				},
			},
			wantErr: false,
			wantXml: `<encoder>Hogeran A.Fugas</encoder>
<encoder type="music">Hogeran A.Fugas</encoder>
<encoder type="arrangement">Hogeran A.Fugas</encoder>
<encoder type="words">Hogeran A.Fugas</encoder>
<encoder type="translation">Hogeran A.Fugas</encoder>
<encoder type="parody">Hogeran A.Fugas</encoder>
<encoder type="???">Hogeran A.Fugas</encoder>`,
		},
		{
			name: "encoder empty",
			args: args{
				eL: &EncoderL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.eL, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("EncoderL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("EncoderL.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
