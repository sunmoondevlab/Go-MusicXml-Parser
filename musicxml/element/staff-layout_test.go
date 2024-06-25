package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestStaffLayoutL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		slL     *StaffLayoutL
		args    args
		wantErr bool
		wantObj StaffLayoutL
	}{
		{
			name: "staff-layout invalid decode",
			slL:  &StaffLayoutL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<staff-layout number="1">
	<staff-distance>60</staff-distance>
</staff-layout>
<staff-layout number="2">
	<staff-distance>60</staff-distance>
</staff-layout>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "staff-layou",
					},
				},
			},
			wantErr: true,
			wantObj: StaffLayoutL{},
		},
		{
			name: "staff-layout",
			slL:  &StaffLayoutL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<staff-layout>
</staff-layout>
<staff-layout number="1">
	<staff-distance>60</staff-distance>
</staff-layout>
<staff-layout number="2">
	<staff-distance>40</staff-distance>
</staff-layout>`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: StaffLayoutL{
				{
					XMLName: xml.Name{
						Local: "staff-layout",
					},
					Number:        nil,
					StaffDistance: nil,
				},
				{
					XMLName: xml.Name{
						Local: "staff-layout",
					},
					Number: &datatype.StaffNumber{
						Val:      testutil.ToUint64Ptr(1),
						Stringer: nil,
					},
					StaffDistance: &StaffDistance{
						XMLName: xml.Name{
							Local: "staff-distance",
						},
						StaffDistance: *datatype.NewTenthsFixedPoint(6, 1),
					},
				},
				{
					XMLName: xml.Name{
						Local: "staff-layout",
					},
					Number: &datatype.StaffNumber{
						Val:      testutil.ToUint64Ptr(2),
						Stringer: nil,
					},
					StaffDistance: &StaffDistance{
						XMLName: xml.Name{
							Local: "staff-distance",
						},
						StaffDistance: *datatype.NewTenthsFixedPoint(4, 1),
					},
				},
			},
		},
		{
			name: "StaffLayout empty",
			slL:  &StaffLayoutL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: StaffLayoutL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.slL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("StaffLayoutL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.slL, tt.wantObj); diff != "" {
			t.Errorf("StaffLayoutL.UnmarshalXML() value is mismatch (-slL, +wantObj):%s\n", diff)
		}
	}
}

func TestStaffLayoutL_MarshalXML(t *testing.T) {
	type args struct {
		slL *StaffLayoutL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "staff-layout",
			args: args{
				slL: &StaffLayoutL{
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number:        nil,
						StaffDistance: nil,
					},
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number: &datatype.StaffNumber{
							Val:      testutil.ToUint64Ptr(1),
							Stringer: nil,
						},
						StaffDistance: &StaffDistance{
							XMLName: xml.Name{
								Local: "staff-distance",
							},
							StaffDistance: *datatype.NewTenthsFixedPoint(6, 1),
						},
					},
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number: &datatype.StaffNumber{
							Val:      testutil.ToUint64Ptr(2),
							Stringer: nil,
						},
						StaffDistance: &StaffDistance{
							XMLName: xml.Name{
								Local: "staff-distance",
							},
							StaffDistance: *datatype.NewTenthsFixedPoint(4, 1),
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<staff-layout></staff-layout>
<staff-layout number="1">
	<staff-distance>60</staff-distance>
</staff-layout>
<staff-layout number="2">
	<staff-distance>40</staff-distance>
</staff-layout>`,
		},
		{
			name: "staff-layout empty",
			args: args{
				slL: &StaffLayoutL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.slL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("StaffLayoutL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("StaffLayoutL.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
