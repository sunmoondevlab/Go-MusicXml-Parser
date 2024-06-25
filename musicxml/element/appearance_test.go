package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestAppearance_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName         xml.Name
		LineWidth       *LineWidthL
		NoteSize        *NoteSizeL
		Distance        *DistanceL
		Glyph           *GlyphL
		OtherAppearance *OtherAppearanceL
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
		wantObj Appearance
	}{
		{
			name:   "appearance invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<appearance>
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
</appearance>`))),
				start: xml.StartElement{Name: xml.Name{Local: "appearanc"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Appearance{},
		},
		{
			name:   "appearance",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<appearance>
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
</appearance>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Appearance{
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
		},
		{
			name:   "appearance empty children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<appearance>
	</appearance>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Appearance{
				XMLName: xml.Name{
					Local: "appearance",
				},
			},
		},
		{
			name:   "appearance empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Appearance{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			aO := &Appearance{
				XMLName:         tt.fields.XMLName,
				LineWidth:       tt.fields.LineWidth,
				NoteSize:        tt.fields.NoteSize,
				Distance:        tt.fields.Distance,
				Glyph:           tt.fields.Glyph,
				OtherAppearance: tt.fields.OtherAppearance,
			}
			if err := aO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Appearance.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*aO, tt.wantObj); diff != "" {
				t.Errorf("Appearance.UnmarshalXML() value is mismatch (-aO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestAppearance_MarshalXML(t *testing.T) {
	type args struct {
		aO *Appearance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "appearance",
			args: args{
				aO: &Appearance{
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
			},
			wantErr: false,
			wantXml: `<appearance>
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
</appearance>`,
		},
		{
			name: "appearance omit children",
			args: args{
				aO: &Appearance{
					XMLName: xml.Name{
						Local: "appearance",
					},
				},
			},
			wantErr: false,
			wantXml: `<appearance></appearance>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.aO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Appearance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Appearance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
