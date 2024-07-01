package datatype

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNewFontFamilyL(t *testing.T) {
	type args struct {
		ffR string
	}
	tests := []struct {
		name string
		args args
		want *FontFamily
	}{
		{
			name: "empty",
			args: args{
				ffR: "",
			},
			want: nil,
		},
		{
			name: "One Font ",
			args: args{
				ffR: "A Font ",
			},
			want: &FontFamily{"A Font"},
		},
		{
			name: "Threee Font ",
			args: args{
				ffR: "A Font , B Font, C Font",
			},
			want: &FontFamily{"A Font", "B Font", "C Font"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewFontFamilyL(tt.args.ffR)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewFontFamilyL() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontFamily_String(t *testing.T) {
	tests := []struct {
		name string
		ff   *FontFamily
		want string
	}{
		{
			name: "nil",
			ff:   nil,
			want: "",
		},
		{
			name: "One Font ",
			ff:   &FontFamily{"A Font"},
			want: "A Font",
		},
		{
			name: "Threee Font ",
			ff:   &FontFamily{"A Font", "B Font", "C Font"},
			want: "A Font,B Font,C Font",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ff.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontFamily.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontFamily_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		ff   *FontFamily
		want *string
	}{
		{
			name: "nil",
			ff:   nil,
			want: nil,
		},
		{
			name: "One Font ",
			ff:   &FontFamily{"A Font"},
			want: testutil.ToStringPtr("A Font"),
		},
		{
			name: "Threee Font ",
			ff:   &FontFamily{"A Font", "B Font", "C Font"},
			want: testutil.ToStringPtr("A Font,B Font,C Font"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ff.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontFamily.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
