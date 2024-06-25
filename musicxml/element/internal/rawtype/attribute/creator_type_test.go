package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestCreatorType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		ct      *CreatorType
		args    args
		wantErr bool
		want    *enum.CreatorTypeEnum
	}{
		{
			name: "invalid",
			ct:   (*CreatorType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "arrange",
				},
			},
			wantErr: true,
			want:    (*enum.CreatorTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			ct:   (*CreatorType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "arranger",
				},
			},
			wantErr: false,
			want:    &enum.CreatorType.Arranger,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ct.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("CreatorType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.CreatorTypeEnum)(tt.ct), tt.want); diff != "" {
				t.Errorf("CreatorType.UnmarshalXMLAttr() value is mismatch (-ct, +want):%s\n", diff)
			}
		})
	}
}
