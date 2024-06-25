package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestLyricFont_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		lfL     *LyricFontL
		args    args
		wantErr bool
		wantObj LyricFontL
	}{
		{
			name: "lyric-font invalid decode",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font />`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "lyric-fon",
					},
				},
			},
			wantErr: true,
			wantObj: LyricFontL{},
		},
		{
			name: "lyric-font invalid font-size css-font-size",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-size="xxx-large"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricFontL{},
		},
		{
			name: "lyric-font invalid font-size font-point-size",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-size="14.0G"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricFontL{},
		},
		{
			name: "lyric-font valid font-size font-point-size",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-size="14.0"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LyricFontL{
				{
					XMLName: xml.Name{
						Local: "lyric-font",
					},
					FontSize: &datatype.FontSize{
						FontPointSize: testutil.ToDecimalPtr("14.0"),
					},
				},
			},
		},
		{
			name: "lyric-font invalid font-style",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-style="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricFontL{},
		},
		{
			name: "lyric-font invalid font-weight",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-weight="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LyricFontL{},
		},
		{
			name: "lyric-font",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="large" font-style="normal" font-weight="normal" name="Tenor" number="Tenor 1"/>
<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold" name="Bass" number="Bass 1"/>
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: LyricFontL{
				{
					XMLName:    xml.Name{Local: "lyric-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
					FontStyle:  &enum.FontStyle.Normal,
					FontWeight: &enum.FontWeight.Normal,
					Name:       testutil.ToStringPtr("Tenor"),
					Number:     testutil.ToStringPtr("Tenor 1"),
				},
				{
					XMLName:    xml.Name{Local: "lyric-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
					FontStyle:  &enum.FontStyle.Italic,
					FontWeight: &enum.FontWeight.Bold,
					Name:       testutil.ToStringPtr("Bass"),
					Number:     testutil.ToStringPtr("Bass 1"),
				},
			},
		},
		{
			name: "lyric-font omit optional",
			lfL:  &LyricFontL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<lyric-font />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: LyricFontL{
				{
					XMLName: xml.Name{Local: "lyric-font"},
				},
			},
		},
		{
			name: "lyric-font empty",
			lfL:  &LyricFontL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LyricFontL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.lfL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("LyricFont.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*tt.lfL, tt.wantObj, opt); diff != "" {
				t.Errorf("LyricFont.UnmarshalXML() value is mismatch (-lfL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestLyricFont_MarshalXML(t *testing.T) {
	type args struct {
		lfL *LyricFontL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "lyric-font",
			args: args{
				lfL: &LyricFontL{
					{
						XMLName:    xml.Name{Local: "lyric-font"},
						FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
						FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
						FontStyle:  &enum.FontStyle.Normal,
						FontWeight: &enum.FontWeight.Normal,
						Name:       testutil.ToStringPtr("Tenor"),
						Number:     testutil.ToStringPtr("Tenor 1"),
					},
					{
						XMLName:    xml.Name{Local: "lyric-font"},
						FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
						FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
						FontStyle:  &enum.FontStyle.Italic,
						FontWeight: &enum.FontWeight.Bold,
						Name:       testutil.ToStringPtr("Bass"),
						Number:     testutil.ToStringPtr("Bass 1"),
					},
				},
			},
			wantErr: false,
			wantXml: `<lyric-font font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" name="Tenor" number="Tenor 1"></lyric-font>
<lyric-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" name="Bass" number="Bass 1"></lyric-font>`,
		},
		{
			name: "lyric-font omit optional",
			args: args{
				lfL: &LyricFontL{
					{
						XMLName: xml.Name{
							Local: "lyric-font",
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<lyric-font></lyric-font>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.lfL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("LyricFont.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("LyricFont.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
