package languages

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func Test_getSupportLanguages(t *testing.T) {
	tests := []struct {
		name     string
		wantLen  int
		wantHead SupportLanguagesL
		wantTail SupportLanguagesL
	}{
		{
			name:    "get support languages",
			wantLen: 1128,
			wantHead: SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			wantTail: SupportLanguagesL{
				{
					Name:         "zh-SG",
					Language:     "zh",
					Region:       testutil.ToStringPtr("SG"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("zho"),
					ISO3Region:   testutil.ToStringPtr("SGP"),
				},
				{
					Name:         "zh-TW",
					Language:     "zh",
					Region:       testutil.ToStringPtr("TW"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("zho"),
					ISO3Region:   testutil.ToStringPtr("TWN"),
				},
				{
					Name:         "zu",
					Language:     "zu",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("zul"),
					ISO3Region:   new(string),
				},
				{
					Name:         "zu-Latn-ZA",
					Language:     "zu",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("zul"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "zu-ZA",
					Language:     "zu",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("zul"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := getSupportLanguages()
			gotHead := (*got)[0:5]
			gotTail := (*got)[len(*got)-5:]
			if diff := cmp.Diff(len(*got), tt.wantLen); diff != "" {
				t.Errorf("getSupportLanguages() value is mismatch (-len(got) +wantLen):%s\n", diff)
			}
			if diff := cmp.Diff(gotHead, tt.wantHead); diff != "" {
				t.Errorf("getSupportLanguages() value is mismatch (-gotHead, +wantHead):%s\n", diff)
			}
			if diff := cmp.Diff(gotTail, tt.wantTail); diff != "" {
				t.Errorf("getSupportLanguages() value is mismatch (-gotTail, +wantTail):%s\n", diff)
			}
		})
	}
}

func TestSupportLanguagesL_BinarySearch(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name  string
		slL   *SupportLanguagesL
		args  args
		wantI int
		want  bool
	}{
		{
			name: "nil",
			slL:  nil,
			args: args{
				value: "af-NA",
			},
			wantI: 0,
			want:  false,
		},
		{
			name: "empty",
			slL:  &SupportLanguagesL{},
			args: args{
				value: "af-NA",
			},
			wantI: 0,
			want:  false,
		},
		{
			name: "found mid",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "af-NA",
			},
			wantI: 2,
			want:  true,
		},
		{
			name: "found left",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "af-Latn-ZA",
			},
			wantI: 1,
			want:  true,
		},
		{
			name: "found first",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "af",
			},
			wantI: 0,
			want:  true,
		},
		{
			name: "found right",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "af-ZA",
			},
			wantI: 3,
			want:  true,
		},
		{
			name: "found last",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "agq",
			},
			wantI: 4,
			want:  true,
		},
		{
			name: "not found left",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "aa",
			},
			wantI: 0,
			want:  false,
		},
		{
			name: "not found last",
			slL: &SupportLanguagesL{
				{
					Name:         "af",
					Language:     "af",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   new(string),
				},
				{
					Name:         "af-Latn-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       testutil.ToStringPtr("Latn"),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "af-NA",
					Language:     "af",
					Region:       testutil.ToStringPtr("NA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("NAM"),
				},
				{
					Name:         "af-ZA",
					Language:     "af",
					Region:       testutil.ToStringPtr("ZA"),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("afr"),
					ISO3Region:   testutil.ToStringPtr("ZAF"),
				},
				{
					Name:         "agq",
					Language:     "agq",
					Region:       new(string),
					Script:       new(string),
					Variant:      new(string),
					Extension:    &LanguageExtension{},
					ISO3Language: testutil.ToStringPtr("agq"),
					ISO3Region:   new(string),
				},
			},
			args: args{
				value: "agr",
			},
			wantI: 5,
			want:  false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gotI, got := tt.slL.BinarySearchByName(tt.args.value)
			if diff := cmp.Diff(gotI, tt.wantI); diff != "" {
				t.Errorf("SupportLanguagesL.BinarySearchByName() value is mismatch (-gotI, +wantI):%s\n", diff)
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SupportLanguagesL.BinarySearchByName() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
