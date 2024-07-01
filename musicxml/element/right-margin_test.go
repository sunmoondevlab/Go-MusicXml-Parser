package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestRightMargin_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		RightMargin datatype.Tenths
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
		wantObj RightMargin
	}{
		{
			name:   "right-margin invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-margin>10</right-margin>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "right-margi",
					},
				},
			},
			wantErr: true,
			wantObj: RightMargin{},
		},
		{
			name:   "right-margin invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<right-margin>G10</right-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightMargin{},
		},
		{
			name:   "right-margin",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<right-margin>10</right-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: RightMargin{
				XMLName: xml.Name{
					Local: "right-margin",
				},
				RightMargin: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "right-margin empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<right-margin></right-margin>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: RightMargin{
				XMLName: xml.Name{
					Local: "right-margin",
				},
			},
		},
		{
			name:   "right-margin empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: RightMargin{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rmO := &RightMargin{
				XMLName:     tt.fields.XMLName,
				RightMargin: tt.fields.RightMargin,
			}
			if err := rmO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RightMargin.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*rmO, tt.wantObj); diff != "" {
				t.Errorf("RightMargin.UnmarshalXML() value is mismatch (-rmO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestRightMargin_MarshalXML(t *testing.T) {
	type args struct {
		rmO *RightMargin
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "right-margin",
			args: args{
				rmO: &RightMargin{
					XMLName: xml.Name{
						Local: "right-margin",
					},
					RightMargin: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<right-margin>10</right-margin>`,
		},
		{
			name:    "right-margin empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.rmO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("RightMargin.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("RightMargin.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
