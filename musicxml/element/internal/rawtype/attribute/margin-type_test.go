package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestMarginType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		mt      *MarginType
		args    args
		wantErr bool
		want    *enum.MarginTypeEnum
	}{
		{
			name: "invalid",
			mt:   (*MarginType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "bot",
				},
			},
			wantErr: true,
			want:    (*enum.MarginTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			mt:   (*MarginType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "both",
				},
			},
			wantErr: false,
			want:    &enum.MarginType.Both,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mt.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("MarginType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.MarginTypeEnum)(tt.mt), tt.want); diff != "" {
				t.Errorf("MarginType.UnmarshalXMLAttr() value is mismatch (-mt, +want):%s\n", diff)
			}
		})
	}
}
