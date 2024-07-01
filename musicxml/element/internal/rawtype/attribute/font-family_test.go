package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestFontFamily_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		ff      *FontFamily
		args    args
		wantErr bool
		want    *FontFamily
	}{
		{
			name: "valid",
			ff:   (*FontFamily)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-family",
					},
					Value: "Edwin",
				},
			},
			wantErr: false,
			want:    (*FontFamily)(testutil.ToStringPtr("Edwin")),
		},
		{
			name: "valid",
			ff:   (*FontFamily)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-family",
					},
					Value: "Edwin,Leland",
				},
			},
			wantErr: false,
			want:    (*FontFamily)(testutil.ToStringPtr("Edwin,Leland")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ff.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("FontFamily.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.ff, tt.want); diff != "" {
				t.Errorf("FontFamily.UnmarshalXMLAttr() value is mismatch (-ff, +want):%s\n", diff)
			}
		})
	}
}

func TestFontFamily_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		ff   *FontFamily
		want *datatype.FontFamily
	}{
		{
			name: "nil",
			ff:   nil,
			want: nil,
		},
		{
			name: "valid",
			ff:   (*FontFamily)(testutil.ToStringPtr("Edwin")),
			want: &datatype.FontFamily{"Edwin"},
		},
		{
			name: "valid",
			ff:   (*FontFamily)(testutil.ToStringPtr("Edwin, Leland, MS P Mincho ")),
			want: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ff.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontFamily.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
