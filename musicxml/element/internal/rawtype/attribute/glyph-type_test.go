package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestGlyphType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		gt      *GlyphType
		args    args
		wantErr bool
		want    *enum.GlyphTypeEnum
	}{
		{
			name: "invalid",
			gt:   (*GlyphType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "g-clef-ottava-bass",
				},
			},
			wantErr: true,
			want:    (*enum.GlyphTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			gt:   (*GlyphType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "g-clef-ottava-bassa",
				},
			},
			wantErr: false,
			want:    &enum.GlyphType.GClefOttavaBassa,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gt.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("GlyphType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.GlyphTypeEnum)(tt.gt), tt.want); diff != "" {
				t.Errorf("GlyphType.UnmarshalXMLAttr() value is mismatch (-gt, +want):%s\n", diff)
			}
		})
	}
}
