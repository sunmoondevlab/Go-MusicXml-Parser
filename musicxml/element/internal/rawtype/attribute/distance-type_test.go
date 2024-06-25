package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestDistanceType_UnmarshalXMLAttr(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		dt      *DistanceType
		args    args
		wantErr bool
		want    *enum.DistanceTypeEnum
	}{
		{
			name: "invalid",
			dt:   (*DistanceType)(testutil.ToStringPtr("")),
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
			want:    (*enum.DistanceTypeEnum)(testutil.ToStringPtr("")),
		},
		{
			name: "valid",
			dt:   (*DistanceType)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "type",
					},
					Value: "beam",
				},
			},
			wantErr: false,
			want:    &enum.DistanceType.Beam,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dt.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("DistanceType.UnmarshalXMLAttr() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff((*enum.DistanceTypeEnum)(tt.dt), tt.want); diff != "" {
				t.Errorf("DistanceType.UnmarshalXMLAttr() value is mismatch (-dt, +want):%s\n", diff)
			}
		})
	}
}
