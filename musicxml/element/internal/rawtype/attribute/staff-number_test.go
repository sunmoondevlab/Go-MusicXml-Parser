package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestStaffNumber_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		sn      *StaffNumber
		args    args
		wantErr bool
		want    *StaffNumber
	}{
		{
			name: "invalid number",
			sn:   (*StaffNumber)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "number",
					},
					Value: "G000",
				},
			},
			wantErr: true,
			want:    (*StaffNumber)(testutil.ToStringPtr("")),
		},
		{
			name: "invalid number",
			sn:   (*StaffNumber)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "number",
					},
					Value: "-1",
				},
			},
			wantErr: true,
			want:    (*StaffNumber)(testutil.ToStringPtr("")),
		},
		{
			name: "valid number",
			sn:   (*StaffNumber)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "number",
					},
					Value: "1111",
				},
			},
			wantErr: false,
			want:    (*StaffNumber)(testutil.ToStringPtr("1111")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sn.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("StaffNumber.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.sn, tt.want); diff != "" {
				t.Errorf("StaffNumber.UnmarshalXMLAttr() value is mismatch (-sn, +want):%s\n", diff)
			}
		})
	}
}

func TestStaffNumber_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		sn   *StaffNumber
		want *datatype.StaffNumber
	}{
		{
			name: "number invalid",
			sn:   (*StaffNumber)(nil),
			want: nil,
		},
		{
			name: "number invalid",
			sn:   (*StaffNumber)(testutil.ToStringPtr("-1")),
			want: nil,
		},
		{
			name: "number",
			sn:   (*StaffNumber)(testutil.ToStringPtr("11")),
			want: &datatype.StaffNumber{
				Val: testutil.ToUint64Ptr(11),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sn.ToEntityData()
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("StaffNumber.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
