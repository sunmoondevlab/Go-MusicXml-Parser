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

func TestLeftDivider_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		Color       *datatype.Color
		DefaultX    *datatype.Tenths
		DefaultY    *datatype.Tenths
		FontFamily  *datatype.FontFamily
		FontSize    *datatype.FontSize
		FontStyle   *enum.FontStyleEnum
		FontWeight  *enum.FontWeightEnum
		Halign      *enum.LeftCenterRightEnum
		PrintObject *enum.YesNoEnum
		RelativeX   *datatype.Tenths
		RelativeY   *datatype.Tenths
		Valign      *enum.ValignEnum
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
		wantObj LeftDivider
	}{
		{
			name:   "left-divider invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider color="#FFFFFF" print-object="yes"/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "left-divide",
					},
				},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid color",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider color="#FFFFE"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid default-x",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider default-x="101,01"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid default-y",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider default-y="811,11"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid font-size css-font-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider font-size="xxx-large"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider font-size="14.0G"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider valid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider font-size="14.0"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LeftDivider{
				XMLName: xml.Name{
					Local: "left-divider",
				},
				FontSize: &datatype.FontSize{
					FontPointSize: testutil.ToDecimalPtr("14.0"),
				},
			},
		},
		{
			name:   "left-divider invalid font-style",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider font-style="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid font-weight",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider font-weight="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid halign",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider halign="lef"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid print-object",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider print-object="ye"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid default-x",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-relative relative-x="101,01"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid relative-y",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider relative-y="811,11"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider invalid valign",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider valign="to"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: LeftDivider{},
		},
		{
			name:   "left-divider",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" print-object="yes" font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"/>
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: LeftDivider{
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
		},
		{
			name:   "left-divider omit optional",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-divider />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: LeftDivider{
				XMLName: xml.Name{Local: "left-divider"},
			},
		},
		{
			name:   "left-divider empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: LeftDivider{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ldO := &LeftDivider{
				XMLName:     tt.fields.XMLName,
				Color:       tt.fields.Color,
				DefaultX:    tt.fields.DefaultX,
				DefaultY:    tt.fields.DefaultY,
				FontFamily:  tt.fields.FontFamily,
				FontSize:    tt.fields.FontSize,
				FontStyle:   tt.fields.FontStyle,
				FontWeight:  tt.fields.FontWeight,
				Halign:      tt.fields.Halign,
				PrintObject: tt.fields.PrintObject,
				RelativeX:   tt.fields.RelativeX,
				RelativeY:   tt.fields.RelativeY,
				Valign:      tt.fields.Valign,
			}
			if err := ldO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("LeftDivider.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*ldO, tt.wantObj, opt); diff != "" {
				t.Errorf("LeftDivider.UnmarshalXML() value is mismatch (-ldO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestLeftDivider_MarshalXML(t *testing.T) {
	type args struct {
		ldO *LeftDivider
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "left-divider",
			args: args{
				ldO: &LeftDivider{
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
			},
			wantErr: false,
			wantXml: `<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>`,
		},
		{
			name: "left-divider omit optional",
			args: args{
				ldO: &LeftDivider{
					XMLName: xml.Name{
						Local: "left-divider",
					},
				},
			},
			wantErr: false,
			wantXml: `<left-divider></left-divider>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.ldO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("LeftDivider.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("LeftDivider.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
