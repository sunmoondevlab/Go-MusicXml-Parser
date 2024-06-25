package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestXlinkType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		xt      *XlinkType
		args    args
		wantErr bool
		want    *enum.XlinkTypeEnum
	}{
		{
			name: "invalid",
			xt:   (*XlinkType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "type",
					},
					Value: "simpl",
				},
			},
			wantErr: true,
			want:    (*enum.XlinkTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			xt:   (*XlinkType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "type",
					},
					Value: "simple",
				},
			},
			wantErr: false,
			want:    &enum.XlinkType.Simple,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.xt.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("XlinkType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.XlinkTypeEnum)(tt.xt), tt.want); diff != "" {
				t.Errorf("XlinkType.UnmarshalXMLAttr() value is mismatch (-xt, +want):%s\n", diff)
			}
		})
	}
}
