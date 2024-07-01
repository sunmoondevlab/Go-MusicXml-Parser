package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestStaffDistance_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName       xml.Name
		StaffDistance datatype.Tenths
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
		wantObj StaffDistance
	}{
		{
			name:   "staff-distance invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<staff-distance>10</staff-distance>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "staff-distanc",
					},
				},
			},
			wantErr: true,
			wantObj: StaffDistance{},
		},
		{
			name:   "staff-distance invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<staff-distance>G10</staff-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: StaffDistance{},
		},
		{
			name:   "staff-distance",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<staff-distance>10</staff-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: StaffDistance{
				XMLName: xml.Name{
					Local: "staff-distance",
				},
				StaffDistance: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "staff-distance empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<staff-distance></staff-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: StaffDistance{
				XMLName: xml.Name{
					Local: "staff-distance",
				},
			},
		},
		{
			name:   "staff-distance empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: StaffDistance{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdO := &StaffDistance{
				XMLName:       tt.fields.XMLName,
				StaffDistance: tt.fields.StaffDistance,
			}
			if err := sdO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("StaffDistance.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*sdO, tt.wantObj); diff != "" {
				t.Errorf("StaffDistance.UnmarshalXML() value is mismatch (-sdO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestStaffDistance_MarshalXML(t *testing.T) {
	type args struct {
		sdO *StaffDistance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "staff-distance",
			args: args{
				sdO: &StaffDistance{
					XMLName: xml.Name{
						Local: "staff-distance",
					},
					StaffDistance: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<staff-distance>10</staff-distance>`,
		},
		{
			name:    "staff-distance empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.sdO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffDistance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("StaffDistance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
