package attribute

import (
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestSmuflGlyphName_UnmarshalText(t *testing.T) {
	type args struct {
		attr xml.Attr
	}
	tests := []struct {
		name    string
		sgnv    *SmuflGlyphName
		args    args
		wantErr bool
		want    SmuflGlyphName
	}{
		{
			name: "empty",
			sgnv: (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "smufl",
					},
					Value: "",
				},
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			sgnv: (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "smufl",
					},
					Value: "aaaa",
				},
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			sgnv: (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				attr: xml.Attr{
					Name: xml.Name{
						Space: "",
						Local: "smufl",
					},
					Value: "fClefFrench",
				},
			},
			wantErr: false,
			want:    (SmuflGlyphName)("fClefFrench"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sgnv.UnmarshalXMLAttr(tt.args.attr); (err != nil) != tt.wantErr {
				t.Errorf("SmuflGlyphName.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.sgnv, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.UnmarshalText() value is mismatch (-sgn, +want):%s\n", diff)
			}
		})
	}
}

func TestSmuflGlyphName_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		sgn  *SmuflGlyphName
		want *datatype.SmuflGlyphName
	}{
		{
			name: "nil",
			sgn:  nil,
			want: nil,
		},
		{
			name: "empty",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("")),
			want: nil,
		},
		{
			name: "invalid",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("aaaaa")),
			want: nil,
		},
		{
			name: "valid",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("fClefFrench")),
			want: &datatype.SmuflGlyphName{
				Val:      testutil.ToStringPtr("fClefFrench"),
				Stringer: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sgn.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
