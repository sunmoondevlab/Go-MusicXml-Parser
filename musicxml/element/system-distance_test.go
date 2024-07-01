package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestSystemDistance_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		SystemDistance datatype.Tenths
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
		wantObj SystemDistance
	}{
		{
			name:   "system-distance invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-distance>10</system-distance>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "system-distanc",
					},
				},
			},
			wantErr: true,
			wantObj: SystemDistance{},
		},
		{
			name:   "system-distance invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<system-distance>G10</system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SystemDistance{},
		},
		{
			name:   "system-distance",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<system-distance>10</system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDistance{
				XMLName: xml.Name{
					Local: "system-distance",
				},
				SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "system-distance empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<system-distance></system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDistance{
				XMLName: xml.Name{
					Local: "system-distance",
				},
			},
		},
		{
			name:   "system-distance empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDistance{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdO := &SystemDistance{
				XMLName:        tt.fields.XMLName,
				SystemDistance: tt.fields.SystemDistance,
			}
			if err := sdO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SystemDistance.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*sdO, tt.wantObj); diff != "" {
				t.Errorf("SystemDistance.UnmarshalXML() value is mismatch (-sdO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestSystemDistance_MarshalXML(t *testing.T) {
	type args struct {
		sdO *SystemDistance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "system-distance",
			args: args{
				sdO: &SystemDistance{
					XMLName: xml.Name{
						Local: "system-distance",
					},
					SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<system-distance>10</system-distance>`,
		},
		{
			name:    "system-distance empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.sdO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemDistance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("SystemDistance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
