package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestFontStyle_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		fs      *FontStyle
		args    args
		wantErr bool
		want    *enum.FontStyleEnum
	}{
		{
			name: "invalid",
			fs:   (*FontStyle)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-style",
					},
					Value: "itali",
				},
			},
			wantErr: true,
			want:    (*enum.FontStyleEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			fs:   (*FontStyle)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-style",
					},
					Value: "italic",
				},
			},
			wantErr: false,
			want:    &enum.FontStyle.Italic,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fs.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("FontStyle.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.FontStyleEnum)(tt.fs), tt.want); diff != "" {
				t.Errorf("FontStyle.UnmarshalXMLAttr() value is mismatch (-fs, +want):%s\n", diff)
			}
		})
	}
}
