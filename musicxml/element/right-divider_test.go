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

func TestRightDivider_UnmarshalXML(t *testing.T) {
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
		wantObj RightDivider
	}{
		{
			name:   "right-divider invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider color="#FFFFFF" print-object="yes"/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "right-divide",
					},
				},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid color",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider color="#FFFFE"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid default-x",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider default-x="101,01"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid default-y",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider default-y="811,11"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid font-size css-font-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider font-size="xxx-large"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider font-size="14.0G"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider valid font-size font-point-size",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider font-size="14.0"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: RightDivider{
				XMLName: xml.Name{
					Local: "right-divider",
				},
				FontSize: &datatype.FontSize{
					FontPointSize: testutil.ToDecimalPtr("14.0"),
				},
			},
		},
		{
			name:   "right-divider invalid font-style",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider font-style="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid font-weight",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider font-weight="norma"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid halign",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider halign="lef"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid print-object",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider print-object="ye"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid default-x",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<left-relative relative-x="101,01"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid relative-y",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider relative-y="811,11"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider invalid valign",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider valign="to"/>
`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: RightDivider{},
		},
		{
			name:   "right-divider",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider color="#FFFFEF" default-x="11.01" default-y="12.01" print-object="yes" font-family="Edwin, Leland, MS P Mincho " font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"/>
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: RightDivider{
				XMLName: xml.Name{Local: "right-divider"},
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
			name:   "right-divider omit optional",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<right-divider />
`))),
				start: xml.StartElement{},
			},

			wantErr: false,
			wantObj: RightDivider{
				XMLName: xml.Name{Local: "right-divider"},
			},
		},
		{
			name:   "right-divider empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: RightDivider{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			rdO := &RightDivider{
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
			if err := rdO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RightDivider.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*rdO, tt.wantObj, opt); diff != "" {
				t.Errorf("RightDivider.UnmarshalXML() value is mismatch (-rdO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestRightDivider_MarshalXML(t *testing.T) {
	type args struct {
		rdO *RightDivider
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "right-divider",
			args: args{
				rdO: &RightDivider{
					XMLName: xml.Name{Local: "right-divider"},
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
			wantXml: `<right-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></right-divider>`,
		},
		{
			name: "right-divider omit optional",
			args: args{
				rdO: &RightDivider{
					XMLName: xml.Name{
						Local: "right-divider",
					},
				},
			},
			wantErr: false,
			wantXml: `<right-divider></right-divider>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.rdO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("RightDivider.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("RightDivider.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
