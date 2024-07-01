package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestValign_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		v       *Valign
		args    args
		wantErr bool
		want    *enum.ValignEnum
	}{
		{
			name: "invalid",
			v:    (*Valign)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "baselin",
				},
			},
			wantErr: true,
			want:    (*enum.ValignEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			v:    (*Valign)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "baseline",
				},
			},
			wantErr: false,
			want:    &enum.Valign.Baseline,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.v.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("Valign.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.ValignEnum)(tt.v), tt.want); diff != "" {
				t.Errorf("Valign.UnmarshalXMLAttr() value is mismatch (-v, +want):%s\n", diff)
			}
		})
	}
}
