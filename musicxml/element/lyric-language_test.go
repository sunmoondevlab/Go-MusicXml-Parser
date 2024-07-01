package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestLyricLanguage_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		lfL     *LyricLanguageL
		args    args
		wantErr bool
		wantObj LyricLanguageL
	}{
		{
			name: "lyric-language invalid decode",
			lfL:  &LyricLanguageL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-language xml:lang="ja-JP"/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "lyric-languag",
					},
				},
			},
			wantErr: true,
			wantObj: LyricLanguageL{},
		},
		{
			name: "lyric-language omit xml:lang",
			lfL:  &LyricLanguageL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-language/>
		`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricLanguageL{},
		},
		{
			name: "lyric-language invalid xml:lang",
			lfL:  &LyricLanguageL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-language xml:lang="aa-AA"/>
		`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricLanguageL{},
		},
		{
			name: "lyric-language",
			lfL:  &LyricLanguageL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-language xml:lang="ja-JP" name="Tenor" number="Tenor 1"/>
<lyric-language xml:lang="ja-Hira-JP" name="Bass" number="Bass 1"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LyricLanguageL{
				{
					XMLName: xml.Name{Local: "lyric-language"},
					XMLLang: &datatype.XmlLang{
						Val: testutil.ToStringPtr("ja-JP"),
					},
					Name:   testutil.ToStringPtr("Tenor"),
					Number: testutil.ToStringPtr("Tenor 1"),
				},
				{
					XMLName: xml.Name{Local: "lyric-language"},
					XMLLang: &datatype.XmlLang{
						Val: testutil.ToStringPtr("ja-Hira-JP"),
					},
					Name:   testutil.ToStringPtr("Bass"),
					Number: testutil.ToStringPtr("Bass 1"),
				},
			},
		},
		{
			name: "lyric-language omit optional",
			lfL:  &LyricLanguageL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-language xml:lang="ja-JP" />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: LyricLanguageL{
				{
					XMLName: xml.Name{Local: "lyric-language"},
					XMLLang: &datatype.XmlLang{
						Val: testutil.ToStringPtr("ja-JP"),
					},
				},
			},
		},
		{
			name: "lyric-language empty",
			lfL:  &LyricLanguageL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LyricLanguageL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lfL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("LyricLanguage.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*tt.lfL, tt.wantObj, opt); diff != "" {
				t.Errorf("LyricLanguage.UnmarshalXML() value is mismatch (-lfL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestLyricLanguage_MarshalXML(t *testing.T) {
	type args struct {
		lfL *LyricLanguageL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "lyric-language omit xml:lang",
			args: args{
				lfL: &LyricLanguageL{
					{
						XMLName: xml.Name{Local: "lyric-language"},
						Name:    testutil.ToStringPtr("Tenor"),
						Number:  testutil.ToStringPtr("Tenor 1"),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "lyric-language",
			args: args{
				lfL: &LyricLanguageL{
					{
						XMLName: xml.Name{Local: "lyric-language"},
						XMLLang: &datatype.XmlLang{
							Val: testutil.ToStringPtr("ja-JP"),
						},
						Name:   testutil.ToStringPtr("Tenor"),
						Number: testutil.ToStringPtr("Tenor 1"),
					},
					{
						XMLName: xml.Name{Local: "lyric-language"},
						XMLLang: &datatype.XmlLang{
							Val: testutil.ToStringPtr("ja-Hira-JP"),
						},
						Name:   testutil.ToStringPtr("Bass"),
						Number: testutil.ToStringPtr("Bass 1"),
					},
				},
			},
			wantErr: false,
			wantXml: `<lyric-language xml:lang="ja-JP" name="Tenor" number="Tenor 1"></lyric-language>
<lyric-language xml:lang="ja-Hira-JP" name="Bass" number="Bass 1"></lyric-language>`,
		},
		{
			name: "lyric-language omit optional",
			args: args{
				lfL: &LyricLanguageL{
					{
						XMLName: xml.Name{
							Local: "lyric-language",
						},
						XMLLang: &datatype.XmlLang{
							Val: testutil.ToStringPtr("ja-Hira-JP"),
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<lyric-language xml:lang="ja-Hira-JP"></lyric-language>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.lfL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("LyricLanguage.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("LyricLanguage.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
