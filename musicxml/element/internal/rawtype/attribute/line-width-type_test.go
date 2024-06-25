package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestLineWidthType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		lwt     *LineWidthType
		args    args
		wantErr bool
		want    *enum.LineWidthTypeEnum
	}{
		{
			name: "invalid",
			lwt:  (*LineWidthType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "bea",
				},
			},
			wantErr: true,
			want:    (*enum.LineWidthTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			lwt:  (*LineWidthType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "heavy barline",
					},
					Value: "heavy barline",
				},
			},
			wantErr: false,
			want:    &enum.LineWidthType.HeavyBarline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lwt.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("LineWidthType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.LineWidthTypeEnum)(tt.lwt), tt.want); diff != "" {
				t.Errorf("LineWidthType.UnmarshalXMLAttr() value is mismatch (-lwt, +want):%s\n", diff)
			}
		})
	}
}
