package glyphnames

import (
	_ "embed"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func Test_getSupportGlyphNames(t *testing.T) {
	tests := []struct {
		name     string
		wantLen  int
		wantHead SupportSmuflGlyphNamesL
		wantTail SupportSmuflGlyphNamesL
	}{
		{
			name:    "get spec glyph names",
			wantLen: 3450,
			wantHead: SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			wantTail: SupportSmuflGlyphNamesL{
				{
					Name:      "windThreeQuartersClosedHole",
					CodePoint: "U+E5F5",
				},
				{
					Name:      "windTightEmbouchure",
					CodePoint: "U+E5FF",
				},
				{
					Name:      "windTrillKey",
					CodePoint: "U+E5FA",
				},
				{
					Name:      "windVeryTightEmbouchure",
					CodePoint: "U+E601",
				},
				{
					Name:      "windWeakAirPressure",
					CodePoint: "U+E602",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSupportGlyphNames()
			gotHead := (*got)[0:5]
			gotTail := (*got)[len(*got)-5:]
			if diff := cmp.Diff(len(*got), tt.wantLen); diff != "" {
				t.Errorf("GetSupportGlyphNames() value is mismatch (-len(got) +wantLen):%s\n", diff)
			}
			if diff := cmp.Diff(gotHead, tt.wantHead); diff != "" {
				t.Errorf("GetSupportGlyphNames() value is mismatch (-gotHead, +wantHead):%s\n", diff)
			}
			if diff := cmp.Diff(gotTail, tt.wantTail); diff != "" {
				t.Errorf("GetSupportGlyphNames() value is mismatch (-gotTail, +wantTail):%s\n", diff)
			}
		})
	}
}

func TestSupportGlyphNamesL_BinarySearch(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		sgnL  *SupportSmuflGlyphNamesL
		args  args
		wantI int
		want  bool
	}{
		{
			name: "nil",
			sgnL: nil,
			args: args{
				value: "4stringTabClefTall",
			},
			wantI: 0,
			want:  false,
		},
		{
			name: "empty",
			sgnL: &SupportSmuflGlyphNamesL{},
			args: args{
				value: "4stringTabClefTall",
			},
			wantI: 0,
			want:  false,
		},
		{
			name: "found mid",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "4stringTabClefTall",
			},
			wantI: 2,
			want:  true,
		},
		{
			name: "found left",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "4stringTabClefSerif",
			},
			wantI: 1,
			want:  true,
		},
		{
			name: "found first",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "4stringTabClef",
			},
			wantI: 0,
			want:  true,
		},
		{
			name: "found right",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "6stringTabClef",
			},
			wantI: 3,
			want:  true,
		},
		{
			name: "found last",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "6stringTabClefSerif",
			},
			wantI: 4,
			want:  true,
		},
		{
			name: "not found left",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "4stringTabCled",
			},
			wantI: -0,
			want:  false,
		},
		{
			name: "not found last",
			sgnL: &SupportSmuflGlyphNamesL{
				{
					Name:      "4stringTabClef",
					CodePoint: "U+E06E",
				},
				{
					Name:      "4stringTabClefSerif",
					CodePoint: "uniE06E.salt02",
				},
				{
					Name:      "4stringTabClefTall",
					CodePoint: "uniE06E.salt01",
				},
				{
					Name:      "6stringTabClef",
					CodePoint: "U+E06D",
				},
				{
					Name:      "6stringTabClefSerif",
					CodePoint: "uniE06D.salt02",
				},
			},
			args: args{
				value: "6stringTabClefSerifd",
			},
			wantI: 5,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, got := tt.sgnL.BinarySearchByName(tt.args.value)
			if diff := cmp.Diff(gotI, tt.wantI); diff != "" {
				t.Errorf("SupportGlyphNamesL.BinarySearchByName() value is mismatch (-gotI, +wantI):%s\n", diff)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SupportGlyphNamesL.BinarySearchByName() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
