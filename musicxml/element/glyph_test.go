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

func TestGlyphL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		gL      *GlyphL
		args    args
		wantErr bool
		wantObj GlyphL
	}{
		{
			name: "glyph invalid decode",
			gL:   &GlyphL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<glyph type="f-clef">fClef19thCentury</glyph>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "glyp",
					},
				},
			},
			wantErr: true,
			wantObj: GlyphL{},
		},
		{
			name: "glyph omit type",
			gL:   &GlyphL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<glyph >fClef19thCentury</glyph>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: GlyphL{},
		},
		{
			name: "glyph invalid type",
			gL:   &GlyphL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<glyph type="f-cle">fClef19thCentury</glyph>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: GlyphL{},
		},
		{
			name: "glyph invalid Smufl",
			gL:   &GlyphL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<glyph type="f-clef">fClef19thCentur</glyph>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: GlyphL{},
		},
		{
			name: "glyph",
			gL:   &GlyphL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<glyph type="f-clef">fClef19thCentury</glyph>
<glyph type="quarter-rest">restQuarterOld</glyph>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: GlyphL{
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
		},
		{
			name: "glyph empty value",
			gL:   &GlyphL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<glyph type="f-clef"></glyph>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: GlyphL{
				{
					XMLName: xml.Name{
						Local: "glyph",
					},
					Type: &enum.GlyphType.FClef,
				},
			},
		},
		{
			name: "glyph empty",
			gL:   &GlyphL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: GlyphL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("GlyphL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.gL, tt.wantObj); diff != "" {
				t.Errorf("GlyphL.UnmarshalXML() value is mismatch (-gL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestGlyphL_MarshalXML(t *testing.T) {
	type args struct {
		gL *GlyphL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "glyph omit type",
			args: args{
				gL: &GlyphL{
					{
						XMLName: xml.Name{
							Local: "glyph",
						},
						Glyph: datatype.SmuflGlyphName{
							Val: testutil.ToStringPtr("fClef19thCentury"),
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "glyph",
			args: args{
				gL: &GlyphL{
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
			},
			wantErr: false,
			wantXml: `<glyph type="f-clef">fClef19thCentury</glyph>
<glyph type="quarter-rest">restQuarterOld</glyph>`,
		},
		{
			name: "glyph empty",
			args: args{
				gL: &GlyphL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.gL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Glyph.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Glyph.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
