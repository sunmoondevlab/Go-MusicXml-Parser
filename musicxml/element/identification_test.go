package element

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestIdentification_UnmarshalXML(t *testing.T) {
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
      <supports "accidental" type="yes"/>
      <supports "beam" type="yes"/>
      <supports "print" attribute="new-page" type="no"/>
      <supports "print" attribute="new-system" type="no"/>
      <supports "stem" type="yes"/>
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
      <supports element= "print" attribute="new-page" type="no"/>
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
				Creator: &Creator{
					XMLName: xml.Name{
						Local: "creator",
					},
					Composer: &CreatorInfo{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Composer,
						Creator: "作曲者:composs",
					},
					Arranger: &CreatorInfo{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Arranger,
						Creator: "編曲者:arran",
					},
					Lyricist: &CreatorInfo{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Lyricist,
						Creator: "作詞者:lilly",
					},
					Poet: &CreatorInfo{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Poet,
						Creator: "作詩者:poet",
					},
					Translator: &CreatorInfo{
						XMLName: xml.Name{
							Local: "creator",
						},
						Type:    enum.CreatorType.Translator,
						Creator: "訳詞者:trann",
					},
				},
				Rights: &RightsL{
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
						CopyrightsType: &enum.CopyrightsType.Music,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyrightsType: &enum.CopyrightsType.Arrangement,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyrightsType: &enum.CopyrightsType.Words,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyrightsType: &enum.CopyrightsType.Translation,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyrightsType: &enum.CopyrightsType.Parody,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyrightsType: &enum.CopyrightsType.Other,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
				},
				Encoding: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &EncoderL{
						{
							XMLName: xml.Name{
								Local: "encoder",
							},
							Encoder: "Hogeran A.Fugas",
						},
					},
					EncodingDate: &EncodingDate{
						XMLName: xml.Name{
							Local: "encoding-date",
						},
						EncodingDate: datatype.YyyyMmDd{
							Val: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
						},
					},
					Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
					EncodingDescription: testutil.ToStringPtr("MusicXML example"),
					Supports: &SupportsL{
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
				Relation: &RelationL{
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
						CopyrightsType: &enum.CopyrightsType.Music,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyrightsType: &enum.CopyrightsType.Arrangement,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyrightsType: &enum.CopyrightsType.Words,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyrightsType: &enum.CopyrightsType.Translation,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyrightsType: &enum.CopyrightsType.Parody,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Local: "relation",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyrightsType: &enum.CopyrightsType.Other,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
				},
				Miscellaneous: &Miscellaneous{
					XMLName: xml.Name{
						Local: "miscellaneous",
					},
					MiscellaneousField: &MiscellaneousFieldL{
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
		{
			name:   "identification empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{}},
			wantErr: false,
			wantObj: Identification{},
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
				t.Errorf("Identification.UnmarshalXML() value is mismatch (-iO, +wantObj):%s\n", diff)
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
					Creator: &Creator{
						XMLName: xml.Name{
							Local: "creator",
						},
						Composer: &CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Composer,
							Creator: "作曲者:composs",
						},
						Arranger: &CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Arranger,
							Creator: "編曲者:arran",
						},
						Lyricist: &CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Lyricist,
							Creator: "作詞者:lilly",
						},
						Poet: &CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Poet,
							Creator: "作詩者:poet",
						},
						Translator: &CreatorInfo{
							XMLName: xml.Name{
								Local: "creator",
							},
							Type:    enum.CreatorType.Translator,
							Creator: "訳詞者:trann",
						},
					},
					Rights: &RightsL{
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
							CopyrightsType: &enum.CopyrightsType.Music,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("arrangement"),
							CopyrightsType: &enum.CopyrightsType.Arrangement,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("words"),
							CopyrightsType: &enum.CopyrightsType.Words,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("translation"),
							CopyrightsType: &enum.CopyrightsType.Translation,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("parody"),
							CopyrightsType: &enum.CopyrightsType.Parody,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
						{
							XMLName: xml.Name{
								Local: "rights",
							},
							Type:           testutil.ToStringPtr("???"),
							CopyrightsType: &enum.CopyrightsType.Other,
							CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
						},
					},
					Encoding: &Encoding{
						XMLName: xml.Name{
							Local: "encoding",
						},
						Encoder: &EncoderL{
							{
								XMLName: xml.Name{
									Local: "encoder",
								},
								Encoder: "Hogeran A.Fugas",
							},
						},
						EncodingDate: &EncodingDate{
							XMLName: xml.Name{
								Local: "encoding-date",
							},
							EncodingDate: datatype.YyyyMmDd{
								Val: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
							},
						},
						Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
						EncodingDescription: testutil.ToStringPtr("MusicXML example"),
						Supports: &SupportsL{
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
					Relation: &RelationL{
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
							CopyrightsType: &enum.CopyrightsType.Music,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("arrangement"),
							CopyrightsType: &enum.CopyrightsType.Arrangement,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("words"),
							CopyrightsType: &enum.CopyrightsType.Words,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("translation"),
							CopyrightsType: &enum.CopyrightsType.Translation,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("parody"),
							CopyrightsType: &enum.CopyrightsType.Parody,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
						{
							XMLName: xml.Name{
								Local: "relation",
							},
							Type:           testutil.ToStringPtr("???"),
							CopyrightsType: &enum.CopyrightsType.Other,
							Relation:       "urn:ISBN:0-000-00000-0",
						},
					},
					Miscellaneous: &Miscellaneous{
						XMLName: xml.Name{
							Local: "miscellaneous",
						},
						MiscellaneousField: &MiscellaneousFieldL{
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
			name: "identification omit children",
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
			b, err := xml.MarshalIndent(tt.args.iO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Identification.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Identification.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
