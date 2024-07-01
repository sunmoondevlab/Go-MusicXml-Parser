package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestMillimeters_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		Millimeters datatype.Millimeters
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
		wantObj Millimeters
	}{
		{
			name:   "millimeters invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<millimeters>100.2319</millimeters>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "millimeter",
					},
				},
			},
			wantErr: true,
			wantObj: Millimeters{},
		},
		{
			name:   "millimeters invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<millimeters>G100.2319</millimeters>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Millimeters{},
		},
		{
			name:   "millimeters",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<millimeters>100.2319</millimeters>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Millimeters{
				XMLName: xml.Name{
					Local: "millimeters",
				},
				Millimeters: *datatype.NewMillimetersFixedPoint(1002319, -4),
			},
		},
		{
			name:   "millimeters empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<millimeters></millimeters>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Millimeters{
				XMLName: xml.Name{
					Local: "millimeters",
				},
			},
		},
		{
			name:   "millimeters empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Millimeters{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mO := &Millimeters{
				XMLName:     tt.fields.XMLName,
				Millimeters: tt.fields.Millimeters,
			}
			if err := mO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Millimeters.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*mO, tt.wantObj); diff != "" {
				t.Errorf("Millimeters.UnmarshalXML() value is mismatch (-mO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestMillimeters_MarshalXML(t *testing.T) {
	type args struct {
		mO *Millimeters
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "millimeters",
			args: args{
				mO: &Millimeters{
					XMLName: xml.Name{
						Local: "millimeters",
					},
					Millimeters: *datatype.NewMillimetersFixedPoint(1002319, -4),
				},
			},
			wantErr: false,
			wantXml: `<millimeters>100.2319</millimeters>`,
		},
		{
			name:    "millimeters empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.mO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Millimeters.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Millimeters.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
