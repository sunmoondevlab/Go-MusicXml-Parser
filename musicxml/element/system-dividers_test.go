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

func TestSystemDividers_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName      xml.Name
		LeftDivider  *LeftDivider
		RightDivider *RightDivider
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
		wantObj SystemDividers
	}{
		{
			name:   "system-dividers invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-dividers>
	<left-divider/>
	<right-divider/>
</system-dividers>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "system-divider",
					},
				},
			},
			wantErr: true,
			wantObj: SystemDividers{},
		},
		{
			name:   "system-dividers omit left-divider",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-dividers>
	<right-divider/>
</system-dividers>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SystemDividers{},
		},
		{
			name:   "system-dividers omit right-divider",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-dividers>
	<left-divider/>
</system-dividers>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SystemDividers{},
		},
		{
			name:   "system-dividers",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-dividers>
	<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
	<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
</system-dividers>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDividers{
				XMLName: xml.Name{
					Local: "system-dividers",
				},
				LeftDivider: &LeftDivider{
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
				RightDivider: &RightDivider{
					XMLName: xml.Name{Local: "right-divider"},
					Color: &datatype.Color{
						WebColorCode: "#FFFFEF",
						Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
					},
					DefaultX: &datatype.Tenths{
						Val: testutil.ToDecimalPtr("701.01"),
					},
					DefaultY: &datatype.Tenths{
						Val: testutil.ToDecimalPtr("13.01"),
					},
					FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
					FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
					FontStyle:   &enum.FontStyle.Normal,
					FontWeight:  &enum.FontWeight.Normal,
					Halign:      &enum.LeftCenterRight.Right,
					PrintObject: &enum.YesNo.Yes,
					RelativeX: &datatype.Tenths{
						Val: testutil.ToDecimalPtr("-0.01"),
					},
					RelativeY: &datatype.Tenths{
						Val: testutil.ToDecimalPtr("-1.01"),
					},
					Valign: &enum.Valign.Top,
				},
			},
		},
		{
			name:   "system-dividers empty",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-dividers>
	<left-divider></left-divider>
	<right-divider></right-divider>				
</system-dividers>`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDividers{
				XMLName: xml.Name{
					Local: "system-dividers",
				},
				LeftDivider: &LeftDivider{
					XMLName: xml.Name{
						Local: "left-divider",
					},
				},
				RightDivider: &RightDivider{
					XMLName: xml.Name{
						Local: "right-divider",
					},
				},
			},
		},
		{
			name:   "system-dividers empty",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemDividers{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sdO := &SystemDividers{
				XMLName:      tt.fields.XMLName,
				LeftDivider:  tt.fields.LeftDivider,
				RightDivider: tt.fields.RightDivider,
			}
			if err := sdO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SystemDividers.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(*sdO, tt.wantObj, opt); diff != "" {
				t.Errorf("SystemDividers.UnmarshalXML() value is mismatch (-sdO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestSystemDividers_MarshalXML(t *testing.T) {
	type args struct {
		sdO *SystemDividers
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "system-dividers omit left-divider",
			args: args{
				sdO: &SystemDividers{
					XMLName: xml.Name{
						Local: "system-dividers",
					},
					RightDivider: &RightDivider{
						XMLName: xml.Name{
							Local: "right-divider",
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "system-dividers omit right-divider",
			args: args{
				sdO: &SystemDividers{
					XMLName: xml.Name{
						Local: "system-dividers",
					},
					LeftDivider: &LeftDivider{
						XMLName: xml.Name{
							Local: "left-divider",
						},
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "system-dividers",
			args: args{
				sdO: &SystemDividers{
					XMLName: xml.Name{
						Local: "system-dividers",
					},
					LeftDivider: &LeftDivider{
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
					RightDivider: &RightDivider{
						XMLName: xml.Name{Local: "right-divider"},
						Color: &datatype.Color{
							WebColorCode: "#FFFFEF",
							Val:          color.RGBA{R: 255, G: 255, B: 239, A: 255},
						},
						DefaultX: &datatype.Tenths{
							Val: testutil.ToDecimalPtr("701.01"),
						},
						DefaultY: &datatype.Tenths{
							Val: testutil.ToDecimalPtr("13.01"),
						},
						FontFamily:  &datatype.FontFamily{"Edwin", "Leland", "MS P Mincho"},
						FontSize:    &datatype.FontSize{CssFontSize: &enum.CssFontSize.Large},
						FontStyle:   &enum.FontStyle.Normal,
						FontWeight:  &enum.FontWeight.Normal,
						Halign:      &enum.LeftCenterRight.Right,
						PrintObject: &enum.YesNo.Yes,
						RelativeX: &datatype.Tenths{
							Val: testutil.ToDecimalPtr("-0.01"),
						},
						RelativeY: &datatype.Tenths{
							Val: testutil.ToDecimalPtr("-1.01"),
						},
						Valign: &enum.Valign.Top,
					},
				},
			},
			wantErr: false,
			wantXml: `<system-dividers>
	<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
	<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
</system-dividers>`,
		},
		{
			name: "system-dividers",
			args: args{
				sdO: &SystemDividers{
					XMLName: xml.Name{
						Local: "system-dividers",
					},
					LeftDivider: &LeftDivider{
						XMLName: xml.Name{Local: "left-divider"},
					},
					RightDivider: &RightDivider{
						XMLName: xml.Name{Local: "right-divider"},
					},
				},
			},
			wantErr: false,
			wantXml: `<system-dividers>
	<left-divider></left-divider>
	<right-divider></right-divider>
</system-dividers>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.sdO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemDividers.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("SystemDividers.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
