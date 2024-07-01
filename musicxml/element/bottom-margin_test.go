package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestBottomMargin_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName      xml.Name
		BottomMargin datatype.Tenths
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
		wantObj BottomMargin
	}{
		{
			name:   "bottom-margin invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<bottom-margin>10</bottom-margin>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "bottom-margi",
					},
				},
			},
			wantErr: true,
			wantObj: BottomMargin{},
		},
		{
			name:   "bottom-margin invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<bottom-margin>G10</bottom-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: BottomMargin{},
		},
		{
			name:   "bottom-margin",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<bottom-margin>10</bottom-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: BottomMargin{
				XMLName: xml.Name{
					Local: "bottom-margin",
				},
				BottomMargin: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "bottom-margin empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<bottom-margin></bottom-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: BottomMargin{
				XMLName: xml.Name{
					Local: "bottom-margin",
				},
			},
		},
		{
			name:   "bottom-margin empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: BottomMargin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			bmO := &BottomMargin{
				XMLName:      tt.fields.XMLName,
				BottomMargin: tt.fields.BottomMargin,
			}
			if err := bmO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("BottomMargin.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*bmO, tt.wantObj); diff != "" {
				t.Errorf("BottomMargin.UnmarshalXML() value is mismatch (-bmO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestBottomMargin_MarshalXML(t *testing.T) {
	type args struct {
		bmO *BottomMargin
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "bottom-margin",
			args: args{
				bmO: &BottomMargin{
					XMLName: xml.Name{
						Local: "bottom-margin",
					},
					BottomMargin: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<bottom-margin>10</bottom-margin>`,
		},
		{
			name:    "bottom-margin empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.bmO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("BottomMargin.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("BottomMargin.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
