package valuetext

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestSmuflGlyphName_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		sgn     *SmuflGlyphName
		args    args
		wantErr bool
		want    SmuflGlyphName
	}{
		{
			name: "empty",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("")),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("aaaa")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			sgn:  (*SmuflGlyphName)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("fClefFrench")),
			},
			wantErr: false,
			want:    (SmuflGlyphName)("fClefFrench"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sgn.UnmarshalText(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("SmuflGlyphName.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.sgn, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.UnmarshalText() value is mismatch (-sgn, +want):%s\n", diff)
			}
		})
	}
}

func TestSmuflGlyphName_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		sgnv SmuflGlyphName
		want datatype.SmuflGlyphName
	}{
		{
			name: "empty",
			sgnv: "",
			want: datatype.SmuflGlyphName{},
		},
		{
			name: "invalid",
			sgnv: (SmuflGlyphName)("aaaaaa"),
			want: datatype.SmuflGlyphName{},
		},
		{
			name: "valid",
			sgnv: (SmuflGlyphName)("fClefFrench"),
			want: datatype.SmuflGlyphName{
				Val:      testutil.ToStringPtr("fClefFrench"),
				Stringer: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.sgnv.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
