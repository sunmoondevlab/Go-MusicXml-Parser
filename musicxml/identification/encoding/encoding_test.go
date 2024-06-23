package musicxml

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
)

func TestEncoding_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName             xml.Name
		Encoder             encoder.EncoderL
		EncodingDate        time.Time
		Software            string
		EncodingDescription string
		Supports            supports.SupportsL
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
		wantObj Encoding
	}{
		{
			name: "encoding invalid decode",
			fields: fields{
				XMLName:             xml.Name{},
				Encoder:             nil,
				EncodingDate:        time.Time{},
				Software:            "",
				EncodingDescription: "",
				Supports:            nil,
			},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      <supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
      </encoding>`))),
				start: xml.StartElement{Name: xml.Name{Space: "", Local: "encodi"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: Encoding{},
		},
		{
			name: "encoding invalid date",
			fields: fields{
				XMLName:             xml.Name{},
				Encoder:             nil,
				EncodingDate:        time.Time{},
				Software:            "",
				EncodingDescription: "",
				Supports:            nil,
			},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-1</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      <supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
      </encoding>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Encoding{},
		},
		{
			name: "encoding",
			fields: fields{
				XMLName:             xml.Name{},
				Encoder:             nil,
				EncodingDate:        time.Time{},
				Software:            "",
				EncodingDescription: "",
				Supports:            nil,
			},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      <supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
      </encoding>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Encoding{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			e := &Encoding{
				XMLName:             tt.fields.XMLName,
				Encoder:             tt.fields.Encoder,
				EncodingDate:        tt.fields.EncodingDate,
				Software:            tt.fields.Software,
				EncodingDescription: tt.fields.EncodingDescription,
				Supports:            tt.fields.Supports,
			}
			if err := e.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Encoding.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*e, tt.wantObj); diff != "" {
				t.Errorf("Encoding.UnmarshalXML() value is mismatch (-*e +tt.wantObj):%s\n", diff)
			}
		})
	}
}
