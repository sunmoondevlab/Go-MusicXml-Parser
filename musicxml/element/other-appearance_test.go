package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestOtherAppearanceL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		gL      *OtherAppearanceL
		args    args
		wantErr bool
		wantObj OtherAppearanceL
	}{
		{
			name: "other-appearance invalid decode",
			gL:   &OtherAppearanceL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<other-appearance type="misc">other</other-appearance>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "other-appearanc",
					},
				},
			},
			wantErr: true,
			wantObj: OtherAppearanceL{},
		},
		{
			name: "other-appearance omit type",
			gL:   &OtherAppearanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<other-appearance >other</other-appearance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: OtherAppearanceL{},
		},
		{
			name: "other-appearance",
			gL:   &OtherAppearanceL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<other-appearance type="misc">other</other-appearance>
<other-appearance type="other">misc</other-appearance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: OtherAppearanceL{
				{
					XMLName: xml.Name{
						Local: "other-appearance",
					},
					Type:            testutil.ToStringPtr("misc"),
					OtherAppearance: "other",
				},
				{
					XMLName: xml.Name{
						Local: "other-appearance",
					},
					Type:            testutil.ToStringPtr("other"),
					OtherAppearance: "misc",
				},
			},
		},
		{
			name: "other-appearance empty value",
			gL:   &OtherAppearanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<other-appearance type="other"></other-appearance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: OtherAppearanceL{
				{
					XMLName: xml.Name{
						Local: "other-appearance",
					},
					Type: testutil.ToStringPtr("other"),
				},
			},
		},
		{
			name: "other-appearance empty",
			gL:   &OtherAppearanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: OtherAppearanceL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.gL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("OtherAppearanceL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.gL, tt.wantObj); diff != "" {
				t.Errorf("OtherAppearanceL.UnmarshalXML() value is mismatch (-gL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestOtherAppearanceL_MarshalXML(t *testing.T) {
	type args struct {
		gL *OtherAppearanceL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "other-appearance omit type",
			args: args{
				gL: &OtherAppearanceL{
					{
						XMLName: xml.Name{
							Local: "other-appearance",
						},
						OtherAppearance: "other",
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "other-appearance",
			args: args{
				gL: &OtherAppearanceL{
					{
						XMLName: xml.Name{
							Local: "other-appearance",
						},
						Type:            testutil.ToStringPtr("misc"),
						OtherAppearance: "other",
					},
					{
						XMLName: xml.Name{
							Local: "other-appearance",
						},
						Type:            testutil.ToStringPtr("other"),
						OtherAppearance: "misc",
					},
				},
			},
			wantErr: false,
			wantXml: `<other-appearance type="misc">other</other-appearance>
<other-appearance type="other">misc</other-appearance>`,
		},
		{
			name: "other-appearance empty",
			args: args{
				gL: &OtherAppearanceL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.gL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("OtherAppearance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("OtherAppearance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
