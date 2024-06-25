package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestFontWeight_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		fw      *FontWeight
		args    args
		wantErr bool
		want    *enum.FontWeightEnum
	}{
		{
			name: "invalid",
			fw:   (*FontWeight)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-weight",
					},
					Value: "bol",
				},
			},
			wantErr: true,
			want:    (*enum.FontWeightEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			fw:   (*FontWeight)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "font-weight",
					},
					Value: "bold",
				},
			},
			wantErr: false,
			want:    &enum.FontWeight.Bold,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.fw.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("FontWeight.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.FontWeightEnum)(tt.fw), tt.want); diff != "" {
				t.Errorf("FontWeight.UnmarshalXMLAttr() value is mismatch (-fw, +want):%s\n", diff)
			}
		})
	}
}
