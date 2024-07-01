package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestXmlLang_UnmarshalText(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		xl      *XmlLang
		args    args
		wantErr bool
		want    XmlLang
	}{
		{
			name: "empty",
			xl:   (*XmlLang)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/XML/1998/namespace",
						Local: "lang",
					},
					Value: "",
				},
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			xl:   (*XmlLang)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/XML/1998/namespace",
						Local: "lang",
					},
					Value: "aa-AA",
				},
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			xl:   (*XmlLang)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/XML/1998/namespace",
						Local: "lang",
					},
					Value: "ja-JP",
				},
			},
			wantErr: false,
			want:    (XmlLang)("ja-JP"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.xl.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("XmlLang.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.xl, tt.want); diff != "" {
				t.Errorf("XmlLang.UnmarshalText() value is mismatch (-xl, +want):%s\n", diff)
			}
		})
	}
}

func TestXmlLang_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		xlv  *XmlLang
		want *datatype.XmlLang
	}{
		{
			name: "nil",
			xlv:  nil,
			want: nil,
		},
		{
			name: "empty",
			xlv:  (*XmlLang)(testutil.ToStringPtr("")),
			want: nil,
		},
		{
			name: "invalid",
			xlv:  (*XmlLang)(testutil.ToStringPtr("aa-AA")),
			want: nil,
		},
		{
			name: "valid",
			xlv:  (*XmlLang)(testutil.ToStringPtr("ja-JP")),
			want: &datatype.XmlLang{
				Val:      testutil.ToStringPtr("ja-JP"),
				Stringer: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.xlv.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XmlLang.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
