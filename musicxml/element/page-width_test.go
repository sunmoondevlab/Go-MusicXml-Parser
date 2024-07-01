package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestPageWidth_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName   xml.Name
		PageWidth datatype.Tenths
	}
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantObj PageWidth
	}{
		{
			name:   "page-width invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-width>10</page-width>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "page-widt",
					},
				},
			},
			wantErr: true,
			wantObj: PageWidth{},
		},
		{
			name:   "page-width invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-width>G10</page-width>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageWidth{},
		},
		{
			name:   "page-width",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-width>10</page-width>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageWidth{
				XMLName: xml.Name{
					Local: "page-width",
				},
				PageWidth: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "page-width empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-width></page-width>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageWidth{
				XMLName: xml.Name{
					Local: "page-width",
				},
			},
		},
		{
			name:   "page-width empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageWidth{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			pwO := &PageWidth{
				XMLName:   tt.fields.XMLName,
				PageWidth: tt.fields.PageWidth,
			}
			if err := pwO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("PageWidth.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*pwO, tt.wantObj); diff != "" {
				t.Errorf("PageWidth.UnmarshalXML() value is mismatch (-pwO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestPageWidth_MarshalXML(t *testing.T) {
	type args struct {
		pwO *PageWidth
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "page-width",
			args: args{
				pwO: &PageWidth{
					XMLName: xml.Name{
						Local: "page-width",
					},
					PageWidth: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<page-width>10</page-width>`,
		},
		{
			name:    "page-width empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.pwO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("PageWidth.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("PageWidth.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
