package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestFontSize_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		fs      *FontSize
		args    args
		wantErr bool
		want    *FontSize
	}{
		{
			name: "invalid",
			fs:   (*FontSize)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-size",
					},
					Value: "40Pt",
				},
			},
			wantErr: true,
			want:    (*FontSize)(testutil.ToStringPtr("")),
		},
		{
			name: "valid css-font-size xx-small",
			fs:   (*FontSize)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-size",
					},
					Value: "xx-small",
				},
			},
			wantErr: false,
			want:    (*FontSize)(testutil.ToStringPtr("xx-small")),
		},
		{
			name: "valid font-point-size",
			fs:   (*FontSize)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-size",
					},
					Value: "14",
				},
			},
			wantErr: false,
			want:    (*FontSize)(testutil.ToStringPtr("14")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("FontSize.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.fs, tt.want); diff != "" {
				t.Errorf("FontSize.UnmarshalXMLAttr() value is mismatch (-fs, +want):%s\n", diff)
			}
		})
	}
}

func TestFontSize_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		fs   *FontSize
		want *datatype.FontSize
	}{
		{
			name: "nil",
			fs:   nil,
			want: nil,
		},
		{
			name: "invalid",
			fs:   (*FontSize)(testutil.ToStringPtr("40Pt")),
			want: nil,
		},
		{
			name: "valid css-font-size xx-small",
			fs:   (*FontSize)(testutil.ToStringPtr("xx-small")),
			want: &datatype.FontSize{
				CssFontSize: &enum.CssFontSize.XXSmall,
			},
		},
		{
			name: "valid font-point-size",
			fs:   (*FontSize)(testutil.ToStringPtr("xx-small")),
			want: &datatype.FontSize{
				CssFontSize: &enum.CssFontSize.XXSmall,
			},
		},
		{
			name: "valid font-point-size",
			fs:   (*FontSize)(testutil.ToStringPtr("14")),
			want: &datatype.FontSize{
				FontPointSize: testutil.DecimalPtr(decimal.NewFromInt(14)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.fs.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontSize.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
