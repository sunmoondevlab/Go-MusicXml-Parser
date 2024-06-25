package element

import (
	"bytes"
	"encoding/xml"
	"image/color"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/google/go-cmp/cmp/cmpopts"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestDefaults_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName       xml.Name
		Creator       *Creator
		Rights        *RightsL
		Encoding      *Encoding
		Source        *string
		Relation      *RelationL
		Miscellaneous *Miscellaneous
	}
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
		wantObj Defaults
	}{
		{
			name:   "defaults",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<defaults>
	<scaling>
		<millimeters>10.3921</millimeters>
		<tenths>10</tenths>
	</scaling>
	<concert-score/>
	<page-layout>
		<page-height>1696.94</page-height>
		<page-width>1200.48</page-width>
		<page-margins type="even">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
		<page-margins type="odd">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
	</page-layout>
	<system-layout>
		<system-margins>
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
		</system-margins>
		<system-distance>10</system-distance>
		<top-system-distance>12</top-system-distance>
		<system-dividers>
			<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
			<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
		</system-dividers>
	</system-layout>
	<staff-layout/>
	<staff-layout number="1">
		<staff-distance>60</staff-distance>
	</staff-layout>
	<staff-layout number="2">
		<staff-distance>40</staff-distance>
	</staff-layout>
	<appearance>
		<line-width type="slur middle">10</line-width>
		<line-width type="ending">10</line-width>
		<note-size type="large">10</note-size>
		<note-size type="grace-cue">10</note-size>
		<distance type="hyphen">10</distance>
		<distance type="beam">1</distance>
		<glyph type="f-clef">fClef19thCentury</glyph>
		<glyph type="quarter-rest">restQuarterOld</glyph>
		<other-appearance type="misc">other</other-appearance>
		<other-appearance type="other">misc</other-appearance>
	</appearance>
	<music-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
	<word-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
	<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="large" font-style="normal" font-weight="normal" name="Tenor" number="Tenor 1"/>
	<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold" name="Bass" number="Bass 1"/>
	<lyric-language xml:lang="ja-JP" name="Tenor" number="Tenor 1"/>
	<lyric-language xml:lang="ja-Hira-JP" name="Bass" number="Bass 1"/>
</defaults>`))),
				start: xml.StartElement{Name: xml.Name{Local: "default"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Defaults{},
		},
		{
			name:   "defaults",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<defaults>
	<scaling>
		<millimeters>10.3921</millimeters>
		<tenths>10</tenths>
	</scaling>
	<concert-score/>
	<page-layout>
		<page-height>1696.94</page-height>
		<page-width>1200.48</page-width>
		<page-margins type="even">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
		<page-margins type="odd">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
	</page-layout>
	<system-layout>
		<system-margins>
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
		</system-margins>
		<system-distance>10</system-distance>
		<top-system-distance>12</top-system-distance>
		<system-dividers>
			<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
			<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
		</system-dividers>
	</system-layout>
	<staff-layout/>
	<staff-layout number="1">
		<staff-distance>60</staff-distance>
	</staff-layout>
	<staff-layout number="2">
		<staff-distance>40</staff-distance>
	</staff-layout>
	<appearance>
		<line-width type="slur middle">10</line-width>
		<line-width type="ending">10</line-width>
		<note-size type="large">10</note-size>
		<note-size type="grace-cue">10</note-size>
		<distance type="hyphen">10</distance>
		<distance type="beam">1</distance>
		<glyph type="f-clef">fClef19thCentury</glyph>
		<glyph type="quarter-rest">restQuarterOld</glyph>
		<other-appearance type="misc">other</other-appearance>
		<other-appearance type="other">misc</other-appearance>
	</appearance>
	<music-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
	<word-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
	<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="large" font-style="normal" font-weight="normal" name="Tenor" number="Tenor 1"/>
	<lyric-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold" name="Bass" number="Bass 1"/>
	<lyric-language xml:lang="ja-JP" name="Tenor" number="Tenor 1"/>
	<lyric-language xml:lang="ja-Hira-JP" name="Bass" number="Bass 1"/>
</defaults>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Defaults{
				XMLName: xml.Name{
					Local: "defaults",
				},
				Scaling: &Scaling{
					XMLName: xml.Name{
						Local: "scaling",
					},
					Millimeters: &Millimeters{
						XMLName: xml.Name{
							Local: "millimeters",
						},
						Millimeters: *datatype.NewMillimetersFixedPoint(103921, -4),
					},
					Tenths: &Tenths{
						XMLName: xml.Name{
							Local: "tenths",
						},
						Tenths: *datatype.NewTenthsFixedPoint(1, 1),
					},
				},
				ConcertScore: &ConcertScore{
					XMLName: xml.Name{
						Local: "concert-score",
					},
				},
				PageLayout: &PageLayout{
					XMLName: xml.Name{Local: "page-layout"},
					PageHeight: &PageHeight{
						XMLName: xml.Name{
							Local: "page-height",
						},
						PageHeight: *datatype.NewTenthsFromFloat(1696.94),
					},
					PageWidth: &PageWidth{
						XMLName: xml.Name{
							Local: "page-width",
						},
						PageWidth: *datatype.NewTenthsFromFloat(1200.48),
					},
					PageMargins: &PageMarginsL{
						{
							XMLName: xml.Name{
								Local: "page-margins",
							},
							Type: &enum.MarginType.Even,
							LeftMargin: &LeftMargin{
								XMLName: xml.Name{
									Local: "left-margin",
								},
								LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
							},
							RightMargin: &RightMargin{
								XMLName: xml.Name{
									Local: "right-margin",
								},
								RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
							},
							TopMargin: &TopMargin{
								XMLName: xml.Name{
									Local: "top-margin",
								},
								TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
							},
							BottomMargin: &BottomMargin{
								XMLName: xml.Name{
									Local: "bottom-margin",
								},
								BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
							},
						},
						{
							XMLName: xml.Name{
								Local: "page-margins",
							},
							Type: &enum.MarginType.Odd,
							LeftMargin: &LeftMargin{
								XMLName: xml.Name{
									Local: "left-margin",
								},
								LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
							},
							RightMargin: &RightMargin{
								XMLName: xml.Name{
									Local: "right-margin",
								},
								RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
							},
							TopMargin: &TopMargin{
								XMLName: xml.Name{
									Local: "top-margin",
								},
								TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
							},
							BottomMargin: &BottomMargin{
								XMLName: xml.Name{
									Local: "bottom-margin",
								},
								BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
							},
						},
					},
				},
				SystemLayout: &SystemLayout{
					XMLName: xml.Name{Local: "system-layout"},
					SystemMargins: &SystemMargins{
						XMLName: xml.Name{
							Local: "system-margins",
						},
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
					},
					SystemDistance: &SystemDistance{
						XMLName: xml.Name{
							Local: "system-distance",
						},
						SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
					},
					TopSystemDistance: &TopSystemDistance{
						XMLName: xml.Name{
							Local: "top-system-distance",
						},
						TopSystemDistance: *datatype.NewTenthsFixedPoint(12, 0),
					},
					SystemDividers: &SystemDividers{
						XMLName: xml.Name{
							Local: "system-dividers",
						},
						LeftDivider: &LeftDivider{
							XMLName: xml.Name{Local: "left-divider"},
							Color: &datatype.Color{
								WebColorCode: "#FFFFEF",
								Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
							},
							DefaultX: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("11.01"),
							},
							DefaultY: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("12.01"),
							},
							FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
							FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
							FontStyle:   &enum.FontStyle.Italic,
							FontWeight:  &enum.FontWeight.Bold,
							Halign:      &enum.LeftCenterRight.Left,
							PrintObject: &enum.YesNo.Yes,
							RelativeX: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("0.01"),
							},
							RelativeY: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("1.01"),
							},
							Valign: &enum.Valign.Top,
						},
						RightDivider: &RightDivider{
							XMLName: xml.Name{Local: "right-divider"},
							Color: &datatype.Color{
								WebColorCode: "#FFFFEF",
								Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
							},
							DefaultX: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("701.01"),
							},
							DefaultY: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("13.01"),
							},
							FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
							FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
							FontStyle:   &enum.FontStyle.Normal,
							FontWeight:  &enum.FontWeight.Normal,
							Halign:      &enum.LeftCenterRight.Right,
							PrintObject: &enum.YesNo.Yes,
							RelativeX: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("-0.01"),
							},
							RelativeY: &datatype.Tenths{
								Val: testutil.ToDecimalPtr("-1.01"),
							},
							Valign: &enum.Valign.Top,
						},
					},
				},
				StaffLayout: &StaffLayoutL{
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number:        nil,
						StaffDistance: nil,
					},
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number: &datatype.StaffNumber{
							Val:      testutil.ToUint64Ptr(1),
							Stringer: nil,
						},
						StaffDistance: &StaffDistance{
							XMLName: xml.Name{
								Local: "staff-distance",
							},
							StaffDistance: *datatype.NewTenthsFixedPoint(6, 1),
						},
					},
					{
						XMLName: xml.Name{
							Local: "staff-layout",
						},
						Number: &datatype.StaffNumber{
							Val:      testutil.ToUint64Ptr(2),
							Stringer: nil,
						},
						StaffDistance: &StaffDistance{
							XMLName: xml.Name{
								Local: "staff-distance",
							},
							StaffDistance: *datatype.NewTenthsFixedPoint(4, 1),
						},
					},
				},
				Appearance: &Appearance{
					XMLName: xml.Name{
						Local: "appearance",
					},
					LineWidth: &LineWidthL{
						{
							XMLName: xml.Name{
								Local: "line-width",
							},
							Type:      &enum.LineWidthType.SlurMiddle,
							LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
						},
						{
							XMLName: xml.Name{
								Local: "line-width",
							},
							Type:      &enum.LineWidthType.Ending,
							LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
						},
					},
					NoteSize: &NoteSizeL{
						{
							XMLName: xml.Name{
								Local: "note-size",
							},
							Type:     &enum.NoteSizeType.Large,
							NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
						},
						{
							XMLName: xml.Name{
								Local: "note-size",
							},
							Type:     &enum.NoteSizeType.GraceCue,
							NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
						},
					},
					Distance: &DistanceL{
						{
							XMLName: xml.Name{
								Local: "distance",
							},
							Type:     &enum.DistanceType.Hyphen,
							Distance: *datatype.NewTenthsFixedPoint(1, 1),
						},
						{
							XMLName: xml.Name{
								Local: "distance",
							},
							Type:     &enum.DistanceType.Beam,
							Distance: *datatype.NewTenthsFixedPoint(1, 0),
						},
					},
					Glyph: &GlyphL{
						{
							XMLName: xml.Name{
								Local: "glyph",
							},
							Type: &enum.GlyphType.FClef,
							Glyph: datatype.SmuflGlyphName{
								Val: testutil.ToStringPtr("fClef19thCentury"),
							},
						},
						{
							XMLName: xml.Name{
								Local: "glyph",
							},
							Type: &enum.GlyphType.QuarterRest,
							Glyph: datatype.SmuflGlyphName{
								Val: testutil.ToStringPtr("restQuarterOld"),
							},
						},
					},
					OtherAppearance: &OtherAppearanceL{
						{
							XMLName: xml.Name{
								Local: "other-appearance",
							},
							Type:            testutil.ToStringPtr("misc"),
							OtherAppearance: "other",
						},
						{
							XMLName: xml.Name{
								Local: "other-appearance",
							},
							Type:            testutil.ToStringPtr("other"),
							OtherAppearance: "misc",
						},
					},
				},
				MusicFont: &MusicFont{
					XMLName:    xml.Name{Local: "music-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
					FontStyle:  &enum.FontStyle.Italic,
					FontWeight: &enum.FontWeight.Bold,
				},
				WordFont: &WordFont{
					XMLName:    xml.Name{Local: "word-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
					FontStyle:  &enum.FontStyle.Italic,
					FontWeight: &enum.FontWeight.Bold,
				},
				LyricFont: &LyricFontL{
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
				LyricLanguage: &LyricLanguageL{
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
		},
		{
			name:   "defaults empty children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<defaults>
	</defaults>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Defaults{
				XMLName: xml.Name{
					Local: "defaults",
				},
			},
		},
		{
			name:   "defaults empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Defaults{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			dO := &Defaults{
				XMLName:       tt.fields.XMLName,
				Scaling:       tt.wantObj.Scaling,
				ConcertScore:  tt.wantObj.ConcertScore,
				PageLayout:    tt.wantObj.PageLayout,
				SystemLayout:  tt.wantObj.SystemLayout,
				StaffLayout:   tt.wantObj.StaffLayout,
				Appearance:    tt.wantObj.Appearance,
				MusicFont:     tt.wantObj.MusicFont,
				WordFont:      tt.wantObj.WordFont,
				LyricFont:     tt.wantObj.LyricFont,
				LyricLanguage: tt.wantObj.LyricLanguage,
			}
			if err := dO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Defaults.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*dO, tt.wantObj, opt); diff != "" {
				t.Errorf("Defaults.UnmarshalXML() value is mismatch (-dO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestDefaults_MarshalXML(t *testing.T) {
	type args struct {
		dO *Defaults
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "defaults",
			args: args{
				dO: &Defaults{
					XMLName: xml.Name{
						Local: "defaults",
					},
					Scaling: &Scaling{
						XMLName: xml.Name{
							Local: "scaling",
						},
						Millimeters: &Millimeters{
							XMLName: xml.Name{
								Local: "millimeters",
							},
							Millimeters: *datatype.NewMillimetersFixedPoint(103921, -4),
						},
						Tenths: &Tenths{
							XMLName: xml.Name{
								Local: "tenths",
							},
							Tenths: *datatype.NewTenthsFixedPoint(1, 1),
						},
					},
					ConcertScore: &ConcertScore{
						XMLName: xml.Name{
							Local: "concert-score",
						},
					},
					PageLayout: &PageLayout{
						XMLName: xml.Name{Local: "page-layout"},
						PageHeight: &PageHeight{
							XMLName: xml.Name{
								Local: "page-height",
							},
							PageHeight: *datatype.NewTenthsFromFloat(1696.94),
						},
						PageWidth: &PageWidth{
							XMLName: xml.Name{
								Local: "page-width",
							},
							PageWidth: *datatype.NewTenthsFromFloat(1200.48),
						},
						PageMargins: &PageMarginsL{
							{
								XMLName: xml.Name{
									Local: "page-margins",
								},
								Type: &enum.MarginType.Even,
								LeftMargin: &LeftMargin{
									XMLName: xml.Name{
										Local: "left-margin",
									},
									LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
								},
								RightMargin: &RightMargin{
									XMLName: xml.Name{
										Local: "right-margin",
									},
									RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
								},
								TopMargin: &TopMargin{
									XMLName: xml.Name{
										Local: "top-margin",
									},
									TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
								},
								BottomMargin: &BottomMargin{
									XMLName: xml.Name{
										Local: "bottom-margin",
									},
									BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
								},
							},
							{
								XMLName: xml.Name{
									Local: "page-margins",
								},
								Type: &enum.MarginType.Odd,
								LeftMargin: &LeftMargin{
									XMLName: xml.Name{
										Local: "left-margin",
									},
									LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
								},
								RightMargin: &RightMargin{
									XMLName: xml.Name{
										Local: "right-margin",
									},
									RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
								},
								TopMargin: &TopMargin{
									XMLName: xml.Name{
										Local: "top-margin",
									},
									TopMargin: *datatype.NewTenthsFixedPoint(857254, -4),
								},
								BottomMargin: &BottomMargin{
									XMLName: xml.Name{
										Local: "bottom-margin",
									},
									BottomMargin: *datatype.NewTenthsFixedPoint(857255, -4),
								},
							},
						},
					},
					SystemLayout: &SystemLayout{
						XMLName: xml.Name{Local: "system-layout"},
						SystemMargins: &SystemMargins{
							XMLName: xml.Name{
								Local: "system-margins",
							},
							LeftMargin: &LeftMargin{
								XMLName: xml.Name{
									Local: "left-margin",
								},
								LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
							},
							RightMargin: &RightMargin{
								XMLName: xml.Name{
									Local: "right-margin",
								},
								RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
							},
						},
						SystemDistance: &SystemDistance{
							XMLName: xml.Name{
								Local: "system-distance",
							},
							SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
						},
						TopSystemDistance: &TopSystemDistance{
							XMLName: xml.Name{
								Local: "top-system-distance",
							},
							TopSystemDistance: *datatype.NewTenthsFixedPoint(12, 0),
						},
						SystemDividers: &SystemDividers{
							XMLName: xml.Name{
								Local: "system-dividers",
							},
							LeftDivider: &LeftDivider{
								XMLName: xml.Name{Local: "left-divider"},
								Color: &datatype.Color{
									WebColorCode: "#FFFFEF",
									Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
								},
								DefaultX: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("11.01"),
								},
								DefaultY: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("12.01"),
								},
								FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
								FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
								FontStyle:   &enum.FontStyle.Italic,
								FontWeight:  &enum.FontWeight.Bold,
								Halign:      &enum.LeftCenterRight.Left,
								PrintObject: &enum.YesNo.Yes,
								RelativeX: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("0.01"),
								},
								RelativeY: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("1.01"),
								},
								Valign: &enum.Valign.Top,
							},
							RightDivider: &RightDivider{
								XMLName: xml.Name{Local: "right-divider"},
								Color: &datatype.Color{
									WebColorCode: "#FFFFEF",
									Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
								},
								DefaultX: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("701.01"),
								},
								DefaultY: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("13.01"),
								},
								FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
								FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
								FontStyle:   &enum.FontStyle.Normal,
								FontWeight:  &enum.FontWeight.Normal,
								Halign:      &enum.LeftCenterRight.Right,
								PrintObject: &enum.YesNo.Yes,
								RelativeX: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("-0.01"),
								},
								RelativeY: &datatype.Tenths{
									Val: testutil.ToDecimalPtr("-1.01"),
								},
								Valign: &enum.Valign.Top,
							},
						},
					},
					StaffLayout: &StaffLayoutL{
						{
							XMLName: xml.Name{
								Local: "staff-layout",
							},
							Number:        nil,
							StaffDistance: nil,
						},
						{
							XMLName: xml.Name{
								Local: "staff-layout",
							},
							Number: &datatype.StaffNumber{
								Val:      testutil.ToUint64Ptr(1),
								Stringer: nil,
							},
							StaffDistance: &StaffDistance{
								XMLName: xml.Name{
									Local: "staff-distance",
								},
								StaffDistance: *datatype.NewTenthsFixedPoint(6, 1),
							},
						},
						{
							XMLName: xml.Name{
								Local: "staff-layout",
							},
							Number: &datatype.StaffNumber{
								Val:      testutil.ToUint64Ptr(2),
								Stringer: nil,
							},
							StaffDistance: &StaffDistance{
								XMLName: xml.Name{
									Local: "staff-distance",
								},
								StaffDistance: *datatype.NewTenthsFixedPoint(4, 1),
							},
						},
					},
					Appearance: &Appearance{
						XMLName: xml.Name{
							Local: "appearance",
						},
						LineWidth: &LineWidthL{
							{
								XMLName: xml.Name{
									Local: "line-width",
								},
								Type:      &enum.LineWidthType.SlurMiddle,
								LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
							},
							{
								XMLName: xml.Name{
									Local: "line-width",
								},
								Type:      &enum.LineWidthType.Ending,
								LineWidth: *datatype.NewTenthsFixedPoint(1, 1),
							},
						},
						NoteSize: &NoteSizeL{
							{
								XMLName: xml.Name{
									Local: "note-size",
								},
								Type:     &enum.NoteSizeType.Large,
								NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
							},
							{
								XMLName: xml.Name{
									Local: "note-size",
								},
								Type:     &enum.NoteSizeType.GraceCue,
								NoteSize: *datatype.NewNonNegativeDecimalFixedPoint(1, 1),
							},
						},
						Distance: &DistanceL{
							{
								XMLName: xml.Name{
									Local: "distance",
								},
								Type:     &enum.DistanceType.Hyphen,
								Distance: *datatype.NewTenthsFixedPoint(1, 1),
							},
							{
								XMLName: xml.Name{
									Local: "distance",
								},
								Type:     &enum.DistanceType.Beam,
								Distance: *datatype.NewTenthsFixedPoint(1, 0),
							},
						},
						Glyph: &GlyphL{
							{
								XMLName: xml.Name{
									Local: "glyph",
								},
								Type: &enum.GlyphType.FClef,
								Glyph: datatype.SmuflGlyphName{
									Val: testutil.ToStringPtr("fClef19thCentury"),
								},
							},
							{
								XMLName: xml.Name{
									Local: "glyph",
								},
								Type: &enum.GlyphType.QuarterRest,
								Glyph: datatype.SmuflGlyphName{
									Val: testutil.ToStringPtr("restQuarterOld"),
								},
							},
						},
						OtherAppearance: &OtherAppearanceL{
							{
								XMLName: xml.Name{
									Local: "other-appearance",
								},
								Type:            testutil.ToStringPtr("misc"),
								OtherAppearance: "other",
							},
							{
								XMLName: xml.Name{
									Local: "other-appearance",
								},
								Type:            testutil.ToStringPtr("other"),
								OtherAppearance: "misc",
							},
						},
					},
					MusicFont: &MusicFont{
						XMLName:    xml.Name{Local: "music-font"},
						FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
						FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
						FontStyle:  &enum.FontStyle.Italic,
						FontWeight: &enum.FontWeight.Bold,
					},
					WordFont: &WordFont{
						XMLName:    xml.Name{Local: "word-font"},
						FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
						FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
						FontStyle:  &enum.FontStyle.Italic,
						FontWeight: &enum.FontWeight.Bold,
					},
					LyricFont: &LyricFontL{
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
					LyricLanguage: &LyricLanguageL{
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
			},
			wantErr: false,
			wantXml: `<defaults>
	<scaling>
		<millimeters>10.3921</millimeters>
		<tenths>10</tenths>
	</scaling>
	<concert-score></concert-score>
	<page-layout>
		<page-height>1696.94</page-height>
		<page-width>1200.48</page-width>
		<page-margins type="even">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
		<page-margins type="odd">
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
			<top-margin>85.7254</top-margin>
			<bottom-margin>85.7255</bottom-margin>
		</page-margins>
	</page-layout>
	<system-layout>
		<system-margins>
			<left-margin>85.7252</left-margin>
			<right-margin>85.7253</right-margin>
		</system-margins>
		<system-distance>10</system-distance>
		<top-system-distance>12</top-system-distance>
		<system-dividers>
			<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
			<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
		</system-dividers>
	</system-layout>
	<staff-layout></staff-layout>
	<staff-layout number="1">
		<staff-distance>60</staff-distance>
	</staff-layout>
	<staff-layout number="2">
		<staff-distance>40</staff-distance>
	</staff-layout>
	<appearance>
		<line-width type="slur middle">10</line-width>
		<line-width type="ending">10</line-width>
		<note-size type="large">10</note-size>
		<note-size type="grace-cue">10</note-size>
		<distance type="hyphen">10</distance>
		<distance type="beam">1</distance>
		<glyph type="f-clef">fClef19thCentury</glyph>
		<glyph type="quarter-rest">restQuarterOld</glyph>
		<other-appearance type="misc">other</other-appearance>
		<other-appearance type="other">misc</other-appearance>
	</appearance>
	<music-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold"></music-font>
	<word-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold"></word-font>
	<lyric-font font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" name="Tenor" number="Tenor 1"></lyric-font>
	<lyric-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" name="Bass" number="Bass 1"></lyric-font>
	<lyric-language xml:lang="ja-JP" name="Tenor" number="Tenor 1"></lyric-language>
	<lyric-language xml:lang="ja-Hira-JP" name="Bass" number="Bass 1"></lyric-language>
</defaults>`,
		},
		{
			name: "defaults omit children",
			args: args{
				dO: &Defaults{
					XMLName: xml.Name{
						Local: "defaults",
					},
				},
			},
			wantErr: false,
			wantXml: `<defaults></defaults>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.dO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Defaults.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Defaults.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
