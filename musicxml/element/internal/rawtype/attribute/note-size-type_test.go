package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNoteSizeType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		nst     *NoteSizeType
		args    args
		wantErr bool
		want    *enum.NoteSizeTypeEnum
	}{
		{
			name: "invalid",
			nst:  (*NoteSizeType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "cu",
				},
			},
			wantErr: true,
			want:    (*enum.NoteSizeTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			nst:  (*NoteSizeType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "cue",
				},
			},
			wantErr: false,
			want:    &enum.NoteSizeType.Cue,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.nst.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("NoteSizeType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.NoteSizeTypeEnum)(tt.nst), tt.want); diff != "" {
				t.Errorf("NoteSizeType.UnmarshalXMLAttr() value is mismatch (-nst, +want):%s\n", diff)
			}
		})
	}
}
