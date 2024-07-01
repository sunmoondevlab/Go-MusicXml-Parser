package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestXlinkActuate_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		xa      *XlinkActuate
		args    args
		wantErr bool
		want    *enum.XlinkActuateEnum
	}{
		{
			name: "invalid",
			xa:   (*XlinkActuate)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "actuate",
					},
					Value: "onLoa",
				},
			},
			wantErr: true,
			want:    (*enum.XlinkActuateEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			xa:   (*XlinkActuate)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "http://www.w3.org/1999/xlink",
						Local: "actuate",
					},
					Value: "onLoad",
				},
			},
			wantErr: false,
			want:    &enum.XlinkActuate.OnLoad,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.xa.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("XlinkActuate.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.XlinkActuateEnum)(tt.xa), tt.want); diff != "" {
				t.Errorf("XlinkActuate.UnmarshalXMLAttr() value is mismatch (-xa, +want):%s\n", diff)
			}
		})
	}
}
