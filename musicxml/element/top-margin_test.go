package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestTopMargin_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName   xml.Name
		TopMargin datatype.Tenths
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
		wantObj TopMargin
	}{
		{
			name:   "top-margin invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<top-margin>10</top-margin>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "top-margi",
					},
				},
			},
			wantErr: true,
			wantObj: TopMargin{},
		},
		{
			name:   "top-margin invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-margin>G10</top-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: TopMargin{},
		},
		{
			name:   "top-margin",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-margin>10</top-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopMargin{
				XMLName: xml.Name{
					Local: "top-margin",
				},
				TopMargin: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "top-margin empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-margin></top-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopMargin{
				XMLName: xml.Name{
					Local: "top-margin",
				},
			},
		},
		{
			name:   "top-margin empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopMargin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmO := &TopMargin{
				XMLName:   tt.fields.XMLName,
				TopMargin: tt.fields.TopMargin,
			}
			if err := tmO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("TopMargin.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tmO, tt.wantObj); diff != "" {
				t.Errorf("TopMargin.UnmarshalXML() value is mismatch (-tmO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestTopMargin_MarshalXML(t *testing.T) {
	type args struct {
		tmO *TopMargin
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "top-margin",
			args: args{
				tmO: &TopMargin{
					XMLName: xml.Name{
						Local: "top-margin",
					},
					TopMargin: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<top-margin>10</top-margin>`,
		},
		{
			name:    "top-margin empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.tmO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("TopMargin.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("TopMargin.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
