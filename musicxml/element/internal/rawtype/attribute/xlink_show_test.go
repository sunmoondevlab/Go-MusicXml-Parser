package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestXlinkShow_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		xs      *XlinkShow
		args    args
		wantErr bool
		want    *enum.XlinkShowEnum
	}{
		{
			name: "invalid",
			xs:   (*XlinkShow)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "show",
					},
					Value: "non",
				},
			},
			wantErr: true,
			want:    (*enum.XlinkShowEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			xs:   (*XlinkShow)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "show",
					},
					Value: "none",
				},
			},
			wantErr: false,
			want:    &enum.XlinkShow.None,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.xs.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("XlinkShow.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.XlinkShowEnum)(tt.xs), tt.want); diff != "" {
				t.Errorf("XlinkShow.UnmarshalXMLAttr() value is mismatch (-xs, +want):%s\n", diff)
			}
		})
	}
}
