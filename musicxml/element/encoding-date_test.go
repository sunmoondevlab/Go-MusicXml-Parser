package element

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestEncodingDate_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName      xml.Name
		EncodingDate datatype.YyyyMmDd
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
		wantObj EncodingDate
	}{
		{
			name:   "encoding-date invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding-date>2024-06-21</encoding-date>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "encoding-dat",
					},
				},
			},
			wantErr: true,
			wantObj: EncodingDate{},
		},
		{
			name:   "encoding-date invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<encoding-date>2024-06-21G</encoding-date>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: EncodingDate{},
		},
		{
			name:   "encoding-date",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<encoding-date>2024-06-21</encoding-date>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: EncodingDate{
				XMLName: xml.Name{
					Local: "encoding-date",
				},
				EncodingDate: datatype.YyyyMmDd{
					Val: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
				},
			},
		},
		{
			name:   "encoding-date empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<encoding-date></encoding-date>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: EncodingDate{
				XMLName: xml.Name{
					Local: "encoding-date",
				},
			},
		},
		{
			name:   "encoding-date empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: EncodingDate{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			edO := &EncodingDate{
				XMLName:      tt.fields.XMLName,
				EncodingDate: tt.fields.EncodingDate,
			}
			if err := edO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("EncodingDate.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*edO, tt.wantObj); diff != "" {
				t.Errorf("EncodingDate.UnmarshalXML() value is mismatch (-edO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestEncodingDate_MarshalXML(t *testing.T) {
	type args struct {
		edO *EncodingDate
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "encoding-date",
			args: args{
				edO: &EncodingDate{
					XMLName: xml.Name{
						Local: "encoding-date",
					},
					EncodingDate: datatype.YyyyMmDd{
						Val: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
					},
				},
			},
			wantErr: false,
			wantXml: `<encoding-date>2024-06-21</encoding-date>`,
		},
		{
			name:    "encoding-date empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.edO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("EncodingDate.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("EncodingDate.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
