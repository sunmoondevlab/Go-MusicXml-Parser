package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestLeftMargin_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName    xml.Name
		LeftMargin datatype.Tenths
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
		wantObj LeftMargin
	}{
		{
			name:   "left-margin invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-margin>10</left-margin>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "left-margi",
					},
				},
			},
			wantErr: true,
			wantObj: LeftMargin{},
		},
		{
			name:   "left-margin invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<left-margin>G10</left-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftMargin{},
		},
		{
			name:   "left-margin",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<left-margin>10</left-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LeftMargin{
				XMLName: xml.Name{
					Local: "left-margin",
				},
				LeftMargin: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "left-margin empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<left-margin></left-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LeftMargin{
				XMLName: xml.Name{
					Local: "left-margin",
				},
			},
		},
		{
			name:   "left-margin empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LeftMargin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lmO := &LeftMargin{
				XMLName:    tt.fields.XMLName,
				LeftMargin: tt.fields.LeftMargin,
			}
			if err := lmO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("LeftMargin.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*lmO, tt.wantObj); diff != "" {
				t.Errorf("LeftMargin.UnmarshalXML() value is mismatch (-lmO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestLeftMargin_MarshalXML(t *testing.T) {
	type args struct {
		lmO *LeftMargin
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "left-margin",
			args: args{
				lmO: &LeftMargin{
					XMLName: xml.Name{
						Local: "left-margin",
					},
					LeftMargin: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<left-margin>10</left-margin>`,
		},
		{
			name:    "left-margin empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.lmO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("LeftMargin.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("LeftMargin.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
