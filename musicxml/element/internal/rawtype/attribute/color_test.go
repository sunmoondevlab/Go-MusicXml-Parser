package attribute

import (
	"encoding/xml"
	"image/color"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestColor_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		c       *Color
		args    args
		wantErr bool
		want    *Color
	}{
		{
			name: "invalid argb",
			c:    (*Color)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "color",
					},
					Value: "#FFFFFFFFF",
				},
			},
			wantErr: true,
			want:    (*Color)(testutil.ToStringPtr("")),
		},
		{
			name: "invalid rgb",
			c:    (*Color)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "color",
					},
					Value: "#FFFFF",
				},
			},
			wantErr: true,
			want:    (*Color)(testutil.ToStringPtr("")),
		},
		{
			name: "valid rgb",
			c:    (*Color)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "color",
					},
					Value: "#FFFEFD",
				},
			},
			wantErr: false,
			want:    (*Color)(testutil.ToStringPtr("#FFFEFD")),
		},
		{
			name: "valid argb",
			c:    (*Color)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "color",
					},
					Value: "#FCFFFEFD",
				},
			},
			wantErr: false,
			want:    (*Color)(testutil.ToStringPtr("#FCFFFEFD")),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.c.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("Color.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.c, tt.want); diff != "" {
				t.Errorf("Color.UnmarshalXMLAttr() value is mismatch (-c, +want):%s\n", diff)
			}
		})
	}
}

func TestColor_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		c    *Color
		want *datatype.Color
	}{
		{
			name: "upper rgb invalid",
			c:    (*Color)(nil),
			want: nil,
		},
		{
			name: "upper rgb invalid",
			c:    (*Color)(testutil.ToStringPtr("#FFFFFFF")),
			want: nil,
		},
		{
			name: "upper rgb white",
			c:    (*Color)(testutil.ToStringPtr("#FFFFFF")),
			want: &datatype.Color{
				WebColorCode: "#FFFFFF",
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
		},
		{
			name: "lower rgb white",
			c:    (*Color)(testutil.ToStringPtr("#ffffff")),
			want: nil,
		},
		{
			name: "upper argb white",
			c:    (*Color)(testutil.ToStringPtr("#FFFFFFFF")),
			want: &datatype.Color{
				WebColorCode: "#FFFFFFFF",
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
		},
		{
			name: "lower argb white",
			c:    (*Color)(testutil.ToStringPtr("#ffffffff")),
			want: nil,
		},
		{
			name: "upper rgb",
			c:    (*Color)(testutil.ToStringPtr("#EEDDCC")),
			want: &datatype.Color{
				WebColorCode: "#EEDDCC",
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
		},
		{
			name: "upper argb",
			c:    (*Color)(testutil.ToStringPtr("#FFEEDDCC")),
			want: &datatype.Color{
				WebColorCode: "#FFEEDDCC",
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.c.ToEntityData()
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("Color.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
