package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestPageHeight_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName    xml.Name
		PageHeight datatype.Tenths
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
		wantObj PageHeight
	}{
		{
			name:   "page-height invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<page-height>10</page-height>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "page-heigh",
					},
				},
			},
			wantErr: true,
			wantObj: PageHeight{},
		},
		{
			name:   "page-height invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-height>G10</page-height>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: PageHeight{},
		},
		{
			name:   "page-height",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-height>10</page-height>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageHeight{
				XMLName: xml.Name{
					Local: "page-height",
				},
				PageHeight: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "page-height empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<page-height></page-height>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageHeight{
				XMLName: xml.Name{
					Local: "page-height",
				},
			},
		},
		{
			name:   "page-height empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: PageHeight{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			phO := &PageHeight{
				XMLName:    tt.fields.XMLName,
				PageHeight: tt.fields.PageHeight,
			}
			if err := phO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("PageHeight.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*phO, tt.wantObj); diff != "" {
				t.Errorf("PageHeight.UnmarshalXML() value is mismatch (-phO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestPageHeight_MarshalXML(t *testing.T) {
	type args struct {
		phO *PageHeight
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "page-height",
			args: args{
				phO: &PageHeight{
					XMLName: xml.Name{
						Local: "page-height",
					},
					PageHeight: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<page-height>10</page-height>`,
		},
		{
			name:    "page-height empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.phO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("PageHeight.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("PageHeight.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
