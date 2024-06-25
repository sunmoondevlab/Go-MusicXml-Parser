package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestTenths_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		Tenths  datatype.Tenths
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
		wantObj Tenths
	}{
		{
			name:   "tenths invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<tenths>10</tenths>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "tenth",
					},
				},
			},
			wantErr: true,
			wantObj: Tenths{},
		},
		{
			name:   "tenths invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<tenths>G10</tenths>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Tenths{},
		},
		{
			name:   "tenths",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<tenths>10</tenths>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Tenths{
				XMLName: xml.Name{
					Local: "tenths",
				},
				Tenths: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "tenths empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<tenths></tenths>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Tenths{
				XMLName: xml.Name{
					Local: "tenths",
				},
			},
		},
		{
			name:   "tenths empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Tenths{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tO := &Tenths{
				XMLName: tt.fields.XMLName,
				Tenths:  tt.fields.Tenths,
			}
			if err := tO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Tenths.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tO, tt.wantObj); diff != "" {
				t.Errorf("Tenths.UnmarshalXML() value is mismatch (-tO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestTenths_MarshalXML(t *testing.T) {
	type args struct {
		tO *Tenths
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "tenths",
			args: args{
				tO: &Tenths{
					XMLName: xml.Name{
						Local: "tenths",
					},
					Tenths: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<tenths>10</tenths>`,
		},
		{
			name:    "tenths empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.tO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Tenths.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Tenths.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
