package identification

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	tCreator "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/creator"
	tEncoding "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
	tMiscellaneous "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous"
	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"

	tRelation "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/relation"
	tRights "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/rights"
)

func TestIdentification_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName       xml.Name
		Creator       *tCreator.Creator
		Rights        *tRights.RightsL
		Encoding      *tEncoding.Encoding
		Source        *string
		Relation      *tRelation.RelationL
		Miscellaneous *tMiscellaneous.Miscellaneous
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
		wantObj Identification
	}{
		{
			name:   "identification",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<identification>
	<creator type="composer">作曲者:composs</creator>
	<creator type="arranger">編曲者:arran</creator>
	<creator type="lyricist">作詞者:lilly</creator>
	<creator type="poet">作詩者:poet</creator>
	<creator type="translator">訳詞者:trann</creator>
	<rights>Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>
	<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      <supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
	</encoding>
	<source>Based on A. Abbb Abbb of 1000, republished by Dover in 1000.</source>
    <relation>urn:ISBN:0-000-00000-0</relation>
    <relation type="music">urn:ISBN:0-000-00000-0</relation>
    <relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
    <relation type="words">urn:ISBN:0-000-00000-0</relation>
    <relation type="translation">urn:ISBN:0-000-00000-0</relation>
    <relation type="parody">urn:ISBN:0-000-00000-0</relation>
    <relation type="???">urn:ISBN:0-000-00000-0</relation>
	<miscellaneous><miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field></miscellaneous>
</identification>`))),
				start: xml.StartElement{Name: xml.Name{Local: "identificatio"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Identification{},
		},
		{
			name:   "identification",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<identification>
	<creator type="composer">作曲者:composs</creator>
	<creator type="arranger">編曲者:arran</creator>
	<creator type="lyricist">作詞者:lilly</creator>
	<creator type="poet">作詩者:poet</creator>
	<creator type="translator">訳詞者:trann</creator>
	<rights>Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>
	<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      <supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
	</encoding>
	<source>Based on A. Abbb Abbb of 1000, republished by Dover in 1000.</source>
    <relation>urn:ISBN:0-000-00000-0</relation>
    <relation type="music">urn:ISBN:0-000-00000-0</relation>
    <relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
    <relation type="words">urn:ISBN:0-000-00000-0</relation>
    <relation type="translation">urn:ISBN:0-000-00000-0</relation>
    <relation type="parody">urn:ISBN:0-000-00000-0</relation>
    <relation type="???">urn:ISBN:0-000-00000-0</relation>
	<miscellaneous><miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field></miscellaneous>
</identification>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Identification{
				XMLName: xml.Name{
					Local: "identification",
				},
				Creator: &tCreator.Creator{
					XMLName: xml.Name{
						Local: "creator",
					},
					Composer: &tCreator.CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Composer,
						Creator: "作曲者:composs",
					},
					Arranger: &tCreator.CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Arranger,
						Creator: "編曲者:arran",
					},
					Lyricist: &tCreator.CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Lyricist,
						Creator: "作詞者:lilly",
					},
					Poet: &tCreator.CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Poet,
						Creator: "作詩者:poet",
					},
					Translator: &tCreator.CreatorElement{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Translator,
						Creator: "訳詞者:trann",
					},
				},
				Rights: &tRights.RightsL{
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						CopyRight: "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("music"),
						CopyRightsType: &enum.CopyRightsType.Music,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyRightsType: &enum.CopyRightsType.Arrangement,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyRightsType: &enum.CopyRightsType.Words,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyRightsType: &enum.CopyRightsType.Translation,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyRightsType: &enum.CopyRightsType.Parody,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyRightsType: &enum.CopyRightsType.Other,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
				},
				Encoding: &tEncoding.Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &encoder.EncoderL{
						{
							XMLName: xml.Name{
								Local: "encoder",
							},
							Encoder: "Hogeran A.Fugas",
						},
					},
					EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
					Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
					EncodingDescription: testutil.ToStringPtr("MusicXML example"),
					Supports: &supports.SupportsL{
						{
							XMLName: xml.Name{
								Local: "supports",
							},
							Element: "accidental",
							Type:    enum.YesNo.Yes,
						},
						{
							XMLName: xml.Name{
								Local: "supports",
							},
							Element: "beam",
							Type:    enum.YesNo.Yes,
						},
						{
							XMLName: xml.Name{
								Local: "supports",
							},
							Element:   "print",
							Type:      enum.YesNo.No,
							Attribute: testutil.ToStringPtr("new-page"),
						},
						{
							XMLName: xml.Name{
								Local: "supports",
							},
							Element:   "print",
							Type:      enum.YesNo.No,
							Attribute: testutil.ToStringPtr("new-system"),
						},
						{
							XMLName: xml.Name{
								Local: "supports",
							},
							Element: "stem",
							Type:    enum.YesNo.Yes,
						},
					},
				},
				Source: testutil.ToStringPtr("Based on A. Abbb Abbb of 1000, republished by Dover in 1000."),
				Relation: &tRelation.RelationL{
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Relation: "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("music"),
						CopyRightsType: &enum.CopyRightsType.Music,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyRightsType: &enum.CopyRightsType.Arrangement,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyRightsType: &enum.CopyRightsType.Words,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyRightsType: &enum.CopyRightsType.Translation,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyRightsType: &enum.CopyRightsType.Parody,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyRightsType: &enum.CopyRightsType.Other,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
				},
				Miscellaneous: &tMiscellaneous.Miscellaneous{
					XMLName: xml.Name{
						Local: "miscellaneous",
					},
					MiscellaneousField: &tMiscellaneousfield.MiscellaneousFieldL{
						{
							XMLName: xml.Name{
								Local: "miscellaneous-field",
							},
							Name:               "difficulty-level",
							MiscellaneousField: "3",
						},
						{
							XMLName: xml.Name{
								Local: "miscellaneous-field",
							},
							Name:               "target-ages",
							MiscellaneousField: "older than 12",
						},
						{
							XMLName: xml.Name{
								Local: "miscellaneous-field",
							},
							Name:               "parts-num",
							MiscellaneousField: "4",
						},
					},
				},
			},
		},
		{
			name:   "identification empty children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<identification>
	</identification>`))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Identification{
				XMLName: xml.Name{
					Local: "identification",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			iO := &Identification{
				XMLName:       tt.fields.XMLName,
				Creator:       tt.fields.Creator,
				Rights:        tt.fields.Rights,
				Encoding:      tt.fields.Encoding,
				Source:        tt.fields.Source,
				Relation:      tt.fields.Relation,
				Miscellaneous: tt.fields.Miscellaneous,
			}
			if err := iO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Identification.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*iO, tt.wantObj); diff != "" {
				t.Errorf("Identification.UnmarshalXML() value is mismatch (-*iO +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestIdentification_MarshalXML(t *testing.T) {
	type args struct {
		iO *Identification
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "identification",
			args: args{
				iO: &Identification{
					XMLName: xml.Name{
						Local: "identification",
					},
					Creator: &tCreator.Creator{
						XMLName: xml.Name{
							Local: "creator",
						},
						Composer: &tCreator.CreatorElement{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Composer,
							Creator: "作曲者:composs",
						},
						Arranger: &tCreator.CreatorElement{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Arranger,
							Creator: "編曲者:arran",
						},
						Lyricist: &tCreator.CreatorElement{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Lyricist,
							Creator: "作詞者:lilly",
						},
						Poet: &tCreator.CreatorElement{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Poet,
							Creator: "作詩者:poet",
						},
						Translator: &tCreator.CreatorElement{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Translator,
							Creator: "訳詞者:trann",
						},
					},
					Rights: &tRights.RightsL{
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							CopyRight: "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("music"),
							CopyRightsType: &enum.CopyRightsType.Music,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("arrangement"),
							CopyRightsType: &enum.CopyRightsType.Arrangement,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("words"),
							CopyRightsType: &enum.CopyRightsType.Words,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("translation"),
							CopyRightsType: &enum.CopyRightsType.Translation,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("parody"),
							CopyRightsType: &enum.CopyRightsType.Parody,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("???"),
							CopyRightsType: &enum.CopyRightsType.Other,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
					},
					Encoding: &tEncoding.Encoding{
						XMLName: xml.Name{
							Local: "encoding",
						},
						Encoder: &encoder.EncoderL{
							{
								XMLName: xml.Name{
									Local: "encoder",
								},
								Encoder: "Hogeran A.Fugas",
							},
						},
						EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
						Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
						EncodingDescription: testutil.ToStringPtr("MusicXML example"),
						Supports: &supports.SupportsL{
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "accidental",
								Type:    enum.YesNo.Yes,
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "beam",
								Type:    enum.YesNo.Yes,
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: testutil.ToStringPtr("new-page"),
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element:   "print",
								Type:      enum.YesNo.No,
								Attribute: testutil.ToStringPtr("new-system"),
							},
							{
								XMLName: xml.Name{
									Local: "supports",
								},
								Element: "stem",
								Type:    enum.YesNo.Yes,
							},
						},
					},
					Source: testutil.ToStringPtr("Based on A. Abbb Abbb of 1000, republished by Dover in 1000."),
					Relation: &tRelation.RelationL{
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Relation: "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("music"),
							CopyRightsType: &enum.CopyRightsType.Music,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("arrangement"),
							CopyRightsType: &enum.CopyRightsType.Arrangement,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("words"),
							CopyRightsType: &enum.CopyRightsType.Words,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("translation"),
							CopyRightsType: &enum.CopyRightsType.Translation,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("parody"),
							CopyRightsType: &enum.CopyRightsType.Parody,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("???"),
							CopyRightsType: &enum.CopyRightsType.Other,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
					},
					Miscellaneous: &tMiscellaneous.Miscellaneous{
						XMLName: xml.Name{
							Local: "miscellaneous",
						},
						MiscellaneousField: &tMiscellaneousfield.MiscellaneousFieldL{
							{
								XMLName: xml.Name{
									Local: "miscellaneous-field",
								},
								Name:               "difficulty-level",
								MiscellaneousField: "3",
							},
							{
								XMLName: xml.Name{
									Local: "miscellaneous-field",
								},
								Name:               "target-ages",
								MiscellaneousField: "older than 12",
							},
							{
								XMLName: xml.Name{
									Local: "miscellaneous-field",
								},
								Name:               "parts-num",
								MiscellaneousField: "4",
							},
						},
					},
				},
			},
			wantErr: false,
			wantXml: `<identification>
  <creator type="composer">作曲者:composs</creator>
  <creator type="arranger">編曲者:arran</creator>
  <creator type="lyricist">作詞者:lilly</creator>
  <creator type="poet">作詩者:poet</creator>
  <creator type="translator">訳詞者:trann</creator>
  <rights>Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
  <rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>
  <encoding>
    <encoder>Hogeran A.Fugas</encoder>
    <encoding-date>2024-06-21</encoding-date>
    <software>MuseScore 4.3.2</software>
    <encoding-description>MusicXML example</encoding-description>
    <supports element="accidental" type="yes"></supports>
    <supports element="beam" type="yes"></supports>
    <supports element="print" type="no" attribute="new-page"></supports>
    <supports element="print" type="no" attribute="new-system"></supports>
    <supports element="stem" type="yes"></supports>
  </encoding>
  <source>Based on A. Abbb Abbb of 1000, republished by Dover in 1000.</source>
  <relation>urn:ISBN:0-000-00000-0</relation>
  <relation type="music">urn:ISBN:0-000-00000-0</relation>
  <relation type="arrangement">urn:ISBN:0-000-00000-0</relation>
  <relation type="words">urn:ISBN:0-000-00000-0</relation>
  <relation type="translation">urn:ISBN:0-000-00000-0</relation>
  <relation type="parody">urn:ISBN:0-000-00000-0</relation>
  <relation type="???">urn:ISBN:0-000-00000-0</relation>
  <miscellaneous>
    <miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field>
  </miscellaneous>
</identification>`,
		},
		{
			name: "identification omit miscelaneous-field",
			args: args{
				iO: &Identification{
					XMLName: xml.Name{
						Local: "identification",
					},
				},
			},
			wantErr: false,
			wantXml: `<identification></identification>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.iO, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("Identification.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Identification.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
