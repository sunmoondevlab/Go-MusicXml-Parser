package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestLineWidthL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		lwL     *LineWidthL
		args    args
		wantErr bool
		wantObj LineWidthL
	}{
		{
			name: "line-width invalid decode",
			lwL:  &LineWidthL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<line-width type="beam">10</line-width>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "line-widt",
					},
				},
			},
			wantErr: true,
			wantObj: LineWidthL{},
		},
		{
			name: "line-width omit type",
			lwL:  &LineWidthL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<line-width >10</line-width>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LineWidthL{},
		},
		{
			name: "line-width invalid type",
			lwL:  &LineWidthL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<line-width type="bean">10</line-width>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LineWidthL{},
		},
		{
			name: "line-width invalid decimal",
			lwL:  &LineWidthL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<line-width type="beam">G10</line-width>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LineWidthL{},
		},
		{
			name: "line-width",
			lwL:  &LineWidthL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<line-width type="slur middle">10</line-width>
<line-width type="ending">10</line-width>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LineWidthL{
				{
					XMLName: xml.Name{
						Local: "line-width",
					},
					Type:      &enum.LineWidthType.SlurMiddle,
					LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
				},
				{
					XMLName: xml.Name{
						Local: "line-width",
					},
					Type:      &enum.LineWidthType.Ending,
					LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
		},
		{
			name: "line-width empty value",
			lwL:  &LineWidthL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<line-width type="beam"></line-width>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LineWidthL{
				{
					XMLName: xml.Name{
						Local: "line-width",
					},
					Type: &enum.LineWidthType.Beam,
				},
			},
		},
		{
			name: "line-width empty",
			lwL:  &LineWidthL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LineWidthL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lwL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("LineWidthL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.lwL, tt.wantObj); diff != "" {
				t.Errorf("LineWidthL.UnmarshalXML() value is mismatch (-lwL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestLineWidthL_MarshalXML(t *testing.T) {
	type args struct {
		lwL *LineWidthL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "line-width omit type",
			args: args{
				lwL: &LineWidthL{
					{
						XMLName: xml.Name{
							Local: "line-width",
						},
						LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "line-width",
			args: args{
				lwL: &LineWidthL{
					{
						XMLName: xml.Name{
							Local: "line-width",
						},
						Type:      &enum.LineWidthType.LightBarline,
						LineWidth: *datatype.NewTenthsFixedPoint(11, 0),
					},
					{
						XMLName: xml.Name{
							Local: "line-width",
						},
						Type:      &enum.LineWidthType.Ending,
						LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
					},
				},
			},
			wantErr: false,
			wantXml: `<line-width type="light barline">11</line-width>
<line-width type="ending">10</line-width>`,
		},
		{
			name: "line-width empty",
			args: args{
				lwL: &LineWidthL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.lwL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("LineWidth.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("LineWidth.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
