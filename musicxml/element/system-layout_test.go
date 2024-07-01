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

func TestSystemLayout_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		SystemMargins     *SystemMargins
		SystemDistance    *SystemDistance
		TopSystemDistance *TopSystemDistance
		SystemDividers    *SystemDividers
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
		wantObj SystemLayout
	}{
		{
			name:   "system-layout invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-layout>
	<system-margins>
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
	</system-margins>
	<system-distance>10</system-distance>
	<top-system-distance>12</top-system-distance>
	<system-dividers>
		<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
		<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
	</system-dividers>
</system-layout>`))),
				start: xml.StartElement{Name: xml.Name{Local: "system-layou"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: SystemLayout{},
		},
		{
			name:   "system-layout",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-layout>
	<system-margins>
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
	</system-margins>
	<system-distance>10</system-distance>
	<top-system-distance>12</top-system-distance>
	<system-dividers>
		<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
		<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
	</system-dividers>
</system-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemLayout{
				XMLName: xml.Name{Local: "system-layout"},
				SystemMargins: &SystemMargins{
					XMLName: xml.Name{
						Local: "system-margins",
					},
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
				},
				SystemDistance: &SystemDistance{
					XMLName: xml.Name{
						Local: "system-distance",
					},
					SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
				},
				TopSystemDistance: &TopSystemDistance{
					XMLName: xml.Name{
						Local: "top-system-distance",
					},
					TopSystemDistance: *datatype.NewTenthsFixedPoint(12, 0),
				},
				SystemDividers: &SystemDividers{
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
		},
		{
			name:   "system-layout omit all children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-layout>
    </system-layout>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemLayout{
				XMLName: xml.Name{
					Local: "system-layout",
				},
			},
		},
		{
			name:   "system-layout empty",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemLayout{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			slO := &SystemLayout{
				XMLName:           tt.fields.XMLName,
				SystemMargins:     tt.fields.SystemMargins,
				SystemDistance:    tt.fields.SystemDistance,
				TopSystemDistance: tt.fields.TopSystemDistance,
				SystemDividers:    tt.fields.SystemDividers,
			}
			if err := slO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SystemLayout.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			opt := cmpopts.IgnoreUnexported(datatype.Color{})
			if diff := cmp.Diff(slO, &tt.wantObj, opt); diff != "" {
				t.Errorf("SystemLayout.UnmarshalXML() value is mismatch (-slO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestSystemLayout_MarshalXML(t *testing.T) {
	type args struct {
		slO *SystemLayout
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "system-layout",
			args: args{
				slO: &SystemLayout{
					XMLName: xml.Name{Local: "system-layout"},
					SystemMargins: &SystemMargins{
						XMLName: xml.Name{
							Local: "system-margins",
						},
						LeftMargin: &LeftMargin{
							XMLName: xml.Name{
								Local: "left-margin",
							},
							LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
						},
						RightMargin: &RightMargin{
							XMLName: xml.Name{
								Local: "right-margin",
							},
							RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
						},
					},
					SystemDistance: &SystemDistance{
						XMLName: xml.Name{
							Local: "system-distance",
						},
						SystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
					},
					TopSystemDistance: &TopSystemDistance{
						XMLName: xml.Name{
							Local: "top-system-distance",
						},
						TopSystemDistance: *datatype.NewTenthsFixedPoint(12, 0),
					},
					SystemDividers: &SystemDividers{
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
			},
			wantErr: false,
			wantXml: `<system-layout>
	<system-margins>
		<left-margin>85.7252</left-margin>
		<right-margin>85.7253</right-margin>
	</system-margins>
	<system-distance>10</system-distance>
	<top-system-distance>12</top-system-distance>
	<system-dividers>
		<left-divider color="#FFFFEF" default-x="11.01" default-y="12.01" font-family="Edwin,Leland,MS P Mincho" font-size="small" font-style="italic" font-weight="bold" halign="left" print-object="yes" relative-x="0.01" relative-y="1.01" valign="top"></left-divider>
		<right-divider color="#FFFFEF" default-x="701.01" default-y="13.01" font-family="Edwin,Leland,MS P Mincho" font-size="large" font-style="normal" font-weight="normal" halign="right" print-object="yes" relative-x="-0.01" relative-y="-1.01" valign="top"></right-divider>
	</system-dividers>
</system-layout>`,
		},
		{
			name: "system-layout omit all children",
			args: args{
				slO: &SystemLayout{
					XMLName: xml.Name{
						Local: "system-layout",
					},
				},
			},
			wantErr: false,
			wantXml: `<system-layout></system-layout>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.slO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemLayout.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("SystemLayout.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
