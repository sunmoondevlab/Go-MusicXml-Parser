package encoder

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestEncoderL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		el      *EncoderL
		args    args
		wantErr bool
		wantObj EncoderL
	}{
		{
			name: "encoder invalid decode",
			el:   &EncoderL{},
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
						Space: "",
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
			el:   &EncoderL{},
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
						Space: "",
						Local: "encoder",
					},
					Type:           "",
					CopyRightsType: enum.CopyRightsType.All,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "music",
					CopyRightsType: enum.CopyRightsType.Music,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "arrangement",
					CopyRightsType: enum.CopyRightsType.Arrangement,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "words",
					CopyRightsType: enum.CopyRightsType.Words,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "translation",
					CopyRightsType: enum.CopyRightsType.Translation,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "parody",
					CopyRightsType: enum.CopyRightsType.Parody,
					Encoder:        "Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "encoder",
					},
					Type:           "???",
					CopyRightsType: enum.CopyRightsType.Other,
					Encoder:        "Hogeran A.Fugas",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.el.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("EncoderL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.el, tt.wantObj); diff != "" {
			t.Errorf("EncoderL.UnmarshalXML() value is mismatch (-*tt.el +tt.wantObj):%s\n", diff)
		}
	}
}
