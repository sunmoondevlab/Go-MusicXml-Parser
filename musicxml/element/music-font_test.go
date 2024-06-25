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

func TestMusicFont_UnmarshalXML(t *testing.T) {
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
		wantObj MusicFont
	}{
		{
			name:   "music-font invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font />`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "music-fon",
					},
				},
			},
			wantErr: true,
			wantObj: MusicFont{},
		},
		{
			name:   "music-font invalid font-size css-font-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-size="xxx-large"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: MusicFont{},
		},
		{
			name:   "music-font invalid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-size="14.0G"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: MusicFont{},
		},
		{
			name:   "music-font valid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-size="14.0"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: MusicFont{
				XMLName: xml.Name{
					Local: "music-font",
				},
				FontSize: &datatype.FontSize{
					FontPointSize: testutil.ToDecimalPtr("14.0"),
				},
			},
		},
		{
			name:   "music-font invalid font-style",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-style="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: MusicFont{},
		},
		{
			name:   "music-font invalid font-weight",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-weight="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: MusicFont{},
		},
		{
			name:   "music-font",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold"/>
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: MusicFont{
				XMLName:    xml.Name{Local: "music-font"},
				FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
				FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
				FontStyle:  &enum.FontStyle.Italic,
				FontWeight: &enum.FontWeight.Bold,
			},
		},
		{
			name:   "music-font omit optional",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<music-font />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: MusicFont{
				XMLName: xml.Name{Local: "music-font"},
			},
		},
		{
			name:   "music-font empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: MusicFont{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mfO := &MusicFont{
				XMLName:    tt.fields.XMLName,
				FontFamily: tt.fields.FontFamily,
				FontSize:   tt.fields.FontSize,
				FontStyle:  tt.fields.FontStyle,
				FontWeight: tt.fields.FontWeight,
			}
			if err := mfO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("MusicFont.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*mfO, tt.wantObj, opt); diff != "" {
				t.Errorf("MusicFont.UnmarshalXML() value is mismatch (-mfO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestMusicFont_MarshalXML(t *testing.T) {
	type args struct {
		mfO *MusicFont
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "music-font",
			args: args{
				mfO: &MusicFont{
					XMLName:    xml.Name{Local: "music-font"},
					FontFamily: &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:   &datatype.FontSize{CssFontSize: &enum.CssFontSize.Small},
					FontStyle:  &enum.FontStyle.Italic,
					FontWeight: &enum.FontWeight.Bold,
				},
			},
			wantErr: false,
			wantXml: `<music-font font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold"></music-font>`,
		},
		{
			name: "music-font omit optional",
			args: args{
				mfO: &MusicFont{
					XMLName: xml.Name{
						Local: "music-font",
					},
				},
			},
			wantErr: false,
			wantXml: `<music-font></music-font>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.mfO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("MusicFont.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("MusicFont.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
