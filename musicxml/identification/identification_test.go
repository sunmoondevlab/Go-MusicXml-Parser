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

	tRelation "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/relation"
	tRights "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/rights"
)

func TestIdentification_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName       xml.Name
		Creator       tCreator.Creator
		Rights        tRights.RightsL
		Encoding      tEncoding.Encoding
		Source        string
		Relation      tRelation.RelationL
		Miscellaneous tMiscellaneous.Miscellaneous
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
				start: xml.StartElement{Name: xml.Name{Space: "", Local: "identificatio"}, Attr: []xml.Attr{}}},
			wantErr: true,
			wantObj: Identification{},
		},
		{
			name:   "identification",
			fields: fields{XMLName: xml.Name{}, Creator: tCreator.Creator{}, Rights: []tRights.Rights{}, Encoding: tEncoding.Encoding{}, Source: "", Relation: []tRelation.Relation{}, Miscellaneous: tMiscellaneous.Miscellaneous{}},
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
					Space: "",
					Local: "identification",
				},
				Creator: tCreator.Creator{
					Composer: tCreator.CreatorElement{
						XMLName: xml.Name{
							Space: "",
							Local: "creator",
						},
						Type:    enum.CreatorType.Composer,
						Creator: "作曲者:composs",
					},
					Arranger: tCreator.CreatorElement{
						XMLName: xml.Name{
							Space: "",
							Local: "creator",
						},
						Type:    enum.CreatorType.Arranger,
						Creator: "編曲者:arran",
					},
					Lyricist: tCreator.CreatorElement{
						XMLName: xml.Name{
							Space: "",
							Local: "creator",
						},
						Type:    enum.CreatorType.Lyricist,
						Creator: "作詞者:lilly",
					},
					Poet: tCreator.CreatorElement{
						XMLName: xml.Name{
							Space: "",
							Local: "creator",
						},
						Type:    enum.CreatorType.Poet,
						Creator: "作詩者:poet",
					},
					Translator: tCreator.CreatorElement{
						XMLName: xml.Name{
							Space: "",
							Local: "creator",
						},
						Type:    enum.CreatorType.Translator,
						Creator: "訳詞者:trann",
					},
				},
				Rights: []tRights.Rights{
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "",
						CopyRightsType: enum.CopyRightsType.All,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "music",
						CopyRightsType: enum.CopyRightsType.Music,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "arrangement",
						CopyRightsType: enum.CopyRightsType.Arrangement,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "words",
						CopyRightsType: enum.CopyRightsType.Words,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "translation",
						CopyRightsType: enum.CopyRightsType.Translation,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "parody",
						CopyRightsType: enum.CopyRightsType.Parody,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "rights",
						},
						Type:           "???",
						CopyRightsType: enum.CopyRightsType.Other,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
				},
				Encoding: tEncoding.Encoding{
					XMLName: xml.Name{
						Space: "",
						Local: "encoding",
					},
					Encoder: encoder.EncoderL{
						{
							XMLName: xml.Name{
								Space: "",
								Local: "encoder",
							},
							Type:           "",
							CopyRightsType: enum.CopyRightsType.All,
							Encoder:        "Hogeran A.Fugas",
						},
					},
					EncodingDate:        time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC),
					Software:            "MuseScore 4.3.2",
					EncodingDescription: "MusicXML example",
					Supports: []supports.Supports{
						{
							XMLName: xml.Name{
								Space: "",
								Local: "supports",
							},
							Element:   "accidental",
							Type:      enum.YesNo.Yes,
							Attribute: "",
							Value:     "",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "supports",
							},
							Element:   "beam",
							Type:      enum.YesNo.Yes,
							Attribute: "",
							Value:     "",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "supports",
							},
							Element:   "print",
							Type:      enum.YesNo.No,
							Attribute: "new-page",
							Value:     "",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "supports",
							},
							Element:   "print",
							Type:      enum.YesNo.No,
							Attribute: "new-system",
							Value:     "",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "supports",
							},
							Element:   "stem",
							Type:      enum.YesNo.Yes,
							Attribute: "",
							Value:     "",
						},
					},
				},
				Source: "Based on A. Abbb Abbb of 1000, republished by Dover in 1000.",
				Relation: []tRelation.Relation{
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "",
						CopyRightsType: enum.CopyRightsType.All,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "music",
						CopyRightsType: enum.CopyRightsType.Music,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "arrangement",
						CopyRightsType: enum.CopyRightsType.Arrangement,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "words",
						CopyRightsType: enum.CopyRightsType.Words,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "translation",
						CopyRightsType: enum.CopyRightsType.Translation,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "parody",
						CopyRightsType: enum.CopyRightsType.Parody,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "relation",
						},
						Type:           "???",
						CopyRightsType: enum.CopyRightsType.Other,
						Relation:       "urn:ISBN:0-000-00000-0",
					},
				},
				Miscellaneous: tMiscellaneous.Miscellaneous{
					XMLName: xml.Name{
						Space: "",
						Local: "miscellaneous",
					},
					MiscellaneousField: []tMiscellaneousfield.MiscellaneousField{
						{
							XMLName: xml.Name{
								Space: "",
								Local: "miscellaneous-field",
							},
							Name:  "difficulty-level",
							Value: "3",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "miscellaneous-field",
							},
							Name:  "target-ages",
							Value: "older than 12",
						},
						{
							XMLName: xml.Name{
								Space: "",
								Local: "miscellaneous-field",
							},
							Name:  "parts-num",
							Value: "4",
						},
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			i := &Identification{
				XMLName:       tt.fields.XMLName,
				Creator:       tt.fields.Creator,
				Rights:        tt.fields.Rights,
				Encoding:      tt.fields.Encoding,
				Source:        tt.fields.Source,
				Relation:      tt.fields.Relation,
				Miscellaneous: tt.fields.Miscellaneous,
			}
			if err := i.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Identification.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*i, tt.wantObj); diff != "" {
				t.Errorf("Identification.UnmarshalXML() value is mismatch (-*i +tt.wantObj):%s\n", diff)
			}
		})
	}
}
