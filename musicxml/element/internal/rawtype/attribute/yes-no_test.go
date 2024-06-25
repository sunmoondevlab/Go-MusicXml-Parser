package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestYesNo_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		yn      *YesNo
		args    args
		wantErr bool
		want    *enum.YesNoEnum
	}{
		{
			name: "invalid",
			yn:   (*YesNo)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "ye",
				},
			},
			wantErr: true,
			want:    (*enum.YesNoEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			yn:   (*YesNo)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "yes",
				},
			},
			wantErr: false,
			want:    &enum.YesNo.Yes,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.yn.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("YesNo.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.YesNoEnum)(tt.yn), tt.want); diff != "" {
				t.Errorf("YesNo.UnmarshalXMLAttr() value is mismatch (-yn, +want):%s\n", diff)
			}
		})
	}
}
