package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestLeftCenterRight_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		lcr     *LeftCenterRight
		args    args
		wantErr bool
		want    *enum.LeftCenterRightEnum
	}{
		{
			name: "invalid",
			lcr:  (*LeftCenterRight)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "halign",
					},
					Value: "lef",
				},
			},
			wantErr: true,
			want:    (*enum.LeftCenterRightEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			lcr:  (*LeftCenterRight)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "halign",
					},
					Value: "left",
				},
			},
			wantErr: false,
			want:    &enum.LeftCenterRight.Left,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lcr.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("LeftCenterRight.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.LeftCenterRightEnum)(tt.lcr), tt.want); diff != "" {
				t.Errorf("LeftCenterRight.UnmarshalXMLAttr() value is mismatch (-lcr, +want):%s\n", diff)
			}
		})
	}
}
