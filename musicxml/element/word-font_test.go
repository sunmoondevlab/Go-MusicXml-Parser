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

func TestWordFont_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName    xml.Name
		FontFamily *datatype.FontFamily
		FontSize   *datatype.FontSize
		FontStyle  *enum.FontStyleEnum
		FontWeight *enum.FontWeightEnum
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
		wantObj WordFont
	}{
		{
			name:   "word-font invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font />`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "word-fon",
					},
				},
			},
			wantErr: true,
			wantObj: WordFont{},
		},
		{
			name:   "word-font invalid font-size css-font-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-size="xxx-large"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: WordFont{},
		},
		{
			name:   "word-font invalid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-size="14.0G"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: WordFont{},
		},
		{
			name:   "word-font valid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-size="14.0"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: WordFont{
				XMLName: xml.Name{
					Local: "word-font",
				},
				FontSize: &datatype.FontSize{
					FontPointSize: testutil.ToDecimalPtr("14.0"),
				},
			},
		},
		{
			name:   "word-font invalid font-style",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-style="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: WordFont{},
		},
		{
			name:   "word-font invalid font-weight",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-weight="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: WordFont{},
		},
		{
			name:   "word-font",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: WordFont{
				XMLName:    xml.Name{Local: "word-font"},
				FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
				FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
				FontStyle:  &enum.FontStyle.Italic,
				FontWeight: &enum.FontWeight.Bold,
			},
		},
		{
			name:   "word-font omit optional",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<word-font />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: WordFont{
				XMLName: xml.Name{Local: "word-font"},
			},
		},
		{
			name:   "word-font empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: WordFont{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wfO := &WordFont{
				XMLName:    tt.fields.XMLName,
				FontFamily: tt.fields.FontFamily,
				FontSize:   tt.fields.FontSize,
				FontStyle:  tt.fields.FontStyle,
				FontWeight: tt.fields.FontWeight,
			}
			if err := wfO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("WordFont.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*wfO, tt.wantObj, opt); diff != "" {
				t.Errorf("WordFont.UnmarshalXML() value is mismatch (-wfO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestWordFont_MarshalXML(t *testing.T) {
	type args struct {
		wfO *WordFont
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "word-font",
			args: args{
				wfO: &WordFont{
					XMLName:    xml.Name{Local: "word-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
					FontStyle:  &enum.FontStyle.Italic,
					FontWeight: &enum.FontWeight.Bold,
				},
			},
			wantErr: false,
			wantXml: `<word-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold"></word-font>`,
		},
		{
			name: "word-font omit optional",
			args: args{
				wfO: &WordFont{
					XMLName: xml.Name{
						Local: "word-font",
					},
				},
			},
			wantErr: false,
			wantXml: `<word-font></word-font>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.wfO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("WordFont.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("WordFont.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
