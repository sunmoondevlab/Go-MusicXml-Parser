package musicxml

import (
	"bytes"
	"encoding/xml"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	tEncoder "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/encoder"
	tSupports "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/encoding/supports"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestEncoding_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName             xml.Name
		Encoder             *tEncoder.EncoderL
		EncodingDate        *time.Time
		Software            *string
		EncodingDescription *string
		Supports            *tSupports.SupportsL
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
			name:   "encoding invalid decode",
			fields: fields{},
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
				start: xml.StartElement{Name: xml.Name{Local: "encodi"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: Encoding{},
		},
		{
			name:   "encoding invalid date",
			fields: fields{},
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
			name:   "encoding",
			fields: fields{},
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
					Local: "encoding",
				},
				Encoder: &tEncoder.EncoderL{
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
				Supports: &tSupports.SupportsL{
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
		}, {
			name:   "encoding omit encoder",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
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
					Local: "encoding",
				},
				EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
				Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
				EncodingDescription: testutil.ToStringPtr("MusicXML example"),
				Supports: &tSupports.SupportsL{
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
		}, {
			name:   "encoding omit encoding-date",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
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
					Local: "encoding",
				},
				Encoder: &tEncoder.EncoderL{
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Encoder: "Hogeran A.Fugas",
					},
				},
				Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
				EncodingDescription: testutil.ToStringPtr("MusicXML example"),
				Supports: &tSupports.SupportsL{
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
		}, {
			name:   "encoding omit software",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
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
					Local: "encoding",
				},
				Encoder: &tEncoder.EncoderL{
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Encoder: "Hogeran A.Fugas",
					},
				},
				EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
				EncodingDescription: testutil.ToStringPtr("MusicXML example"),
				Supports: &tSupports.SupportsL{
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
		},
		{
			name:   "encoding omit encoding-description",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
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
					Local: "encoding",
				},
				Encoder: &tEncoder.EncoderL{
					{
						XMLName: xml.Name{
							Local: "encoder",
						},
						Encoder: "Hogeran A.Fugas",
					},
				},
				EncodingDate: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
				Software:     testutil.ToStringPtr("MuseScore 4.3.2"),
				Supports: &tSupports.SupportsL{
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
		},
		{
			name:   "encoding supports",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<encoding>
      <encoder>Hogeran A.Fugas</encoder>
      <software>MuseScore 4.3.2</software>
      <encoding-date>2024-06-21</encoding-date>
      <encoding-description>MusicXML example</encoding-description>
      </encoding>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Encoding{
				XMLName: xml.Name{
					Local: "encoding",
				},
				Encoder: &tEncoder.EncoderL{
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
			},
		},
		{
			name:   "encoding empty children",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<encoding></encoding>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Encoding{
				XMLName: xml.Name{
					Local: "encoding",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			eO := &Encoding{
				XMLName:             tt.fields.XMLName,
				Encoder:             tt.fields.Encoder,
				EncodingDate:        tt.fields.EncodingDate,
				Software:            tt.fields.Software,
				EncodingDescription: tt.fields.EncodingDescription,
				Supports:            tt.fields.Supports,
			}
			if err := eO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Encoding.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*eO, tt.wantObj); diff != "" {
				t.Errorf("Encoding.UnmarshalXML() value is mismatch (-*eO +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestEncoding_MarshalXML(t *testing.T) {
	type args struct {
		eO *Encoding
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "encoding",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &tEncoder.EncoderL{
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
					Supports: &tSupports.SupportsL{
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
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoder>Hogeran A.Fugas</encoder>
  <encoding-date>2024-06-21</encoding-date>
  <software>MuseScore 4.3.2</software>
  <encoding-description>MusicXML example</encoding-description>
  <supports element="accidental" type="yes"></supports>
  <supports element="beam" type="yes"></supports>
  <supports element="print" type="no" attribute="new-page"></supports>
  <supports element="print" type="no" attribute="new-system"></supports>
  <supports element="stem" type="yes"></supports>
</encoding>`,
		},
		{
			name: "encoding omit encoder",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
					Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
					EncodingDescription: testutil.ToStringPtr("MusicXML example"),
					Supports: &tSupports.SupportsL{
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
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoding-date>2024-06-21</encoding-date>
  <software>MuseScore 4.3.2</software>
  <encoding-description>MusicXML example</encoding-description>
  <supports element="accidental" type="yes"></supports>
  <supports element="beam" type="yes"></supports>
  <supports element="print" type="no" attribute="new-page"></supports>
  <supports element="print" type="no" attribute="new-system"></supports>
  <supports element="stem" type="yes"></supports>
</encoding>`,
		},
		{
			name: "encoding omit encoding-date",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &tEncoder.EncoderL{
						{
							XMLName: xml.Name{
								Local: "encoder",
							},
							Encoder: "Hogeran A.Fugas",
						},
					},
					Software:            testutil.ToStringPtr("MuseScore 4.3.2"),
					EncodingDescription: testutil.ToStringPtr("MusicXML example"),
					Supports: &tSupports.SupportsL{
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
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoder>Hogeran A.Fugas</encoder>
  <software>MuseScore 4.3.2</software>
  <encoding-description>MusicXML example</encoding-description>
  <supports element="accidental" type="yes"></supports>
  <supports element="beam" type="yes"></supports>
  <supports element="print" type="no" attribute="new-page"></supports>
  <supports element="print" type="no" attribute="new-system"></supports>
  <supports element="stem" type="yes"></supports>
</encoding>`,
		},
		{
			name: "encoding omit software",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &tEncoder.EncoderL{
						{
							XMLName: xml.Name{
								Local: "encoder",
							},
							Encoder: "Hogeran A.Fugas",
						},
					},
					EncodingDate:        testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
					EncodingDescription: testutil.ToStringPtr("MusicXML example"),
					Supports: &tSupports.SupportsL{
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
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoder>Hogeran A.Fugas</encoder>
  <encoding-date>2024-06-21</encoding-date>
  <encoding-description>MusicXML example</encoding-description>
  <supports element="accidental" type="yes"></supports>
  <supports element="beam" type="yes"></supports>
  <supports element="print" type="no" attribute="new-page"></supports>
  <supports element="print" type="no" attribute="new-system"></supports>
  <supports element="stem" type="yes"></supports>
</encoding>`,
		},
		{
			name: "encoding omit encoding-description",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &tEncoder.EncoderL{
						{
							XMLName: xml.Name{
								Local: "encoder",
							},
							Encoder: "Hogeran A.Fugas",
						},
					},
					EncodingDate: testutil.ToTimePtr(time.Date(2024, time.June, 21, 0, 0, 0, 0, time.UTC)),
					Software:     testutil.ToStringPtr("MuseScore 4.3.2"),
					Supports: &tSupports.SupportsL{
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
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoder>Hogeran A.Fugas</encoder>
  <encoding-date>2024-06-21</encoding-date>
  <software>MuseScore 4.3.2</software>
  <supports element="accidental" type="yes"></supports>
  <supports element="beam" type="yes"></supports>
  <supports element="print" type="no" attribute="new-page"></supports>
  <supports element="print" type="no" attribute="new-system"></supports>
  <supports element="stem" type="yes"></supports>
</encoding>`,
		},
		{
			name: "encoding omit supports",
			args: args{
				eO: &Encoding{
					XMLName: xml.Name{
						Local: "encoding",
					},
					Encoder: &tEncoder.EncoderL{
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
				},
			},
			wantErr: false,
			wantXml: `<encoding>
  <encoder>Hogeran A.Fugas</encoder>
  <encoding-date>2024-06-21</encoding-date>
  <software>MuseScore 4.3.2</software>
  <encoding-description>MusicXML example</encoding-description>
</encoding>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.eO, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("Encoding.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Encoding.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
