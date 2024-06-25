package attribute

import (
	"encoding/xml"
	"reflect"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestTenths_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		t       *Tenths
		args    args
		wantErr bool
		want    *Tenths
	}{
		{
			name: "invalid",
			t:    (*Tenths)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "default-x",
					},
					Value: "G0.0",
				},
			},
			wantErr: true,
			want:    (*Tenths)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			t:    (*Tenths)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "default-x",
					},
					Value: "105.01",
				},
			},
			wantErr: false,
			want:    (*Tenths)(testutil.ToStringPtr("105.01")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("Tenths.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.t, tt.want); diff != "" {
				t.Errorf("Tenths.UnmarshalXMLAttr() value is mismatch (-t, +want):%s\n", diff)
			}
		})
	}
}

func TestTenths_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		t    *Tenths
		want *datatype.Tenths
	}{
		{
			name: "nil",
			t:    nil,
			want: nil,
		},
		{
			name: "invalid",
			t:    (*Tenths)(testutil.ToStringPtr("G105.01")),
			want: nil,
		},
		{
			name: "valid",
			t:    (*Tenths)(testutil.ToStringPtr("105.01")),
			want: datatype.NewTenthsFixedPoint(10501, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.ToEntityData()
			if got := tt.t.ToEntityData(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Tenths.ToEntityData() = %v, want %v", got, tt.want)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tenths.ToEntityData() value is mismatch -got, want):%s\n", diff)
			}
		})
	}
}
