package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestSupportsL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		sL      *SupportsL
		args    args
		wantErr bool
		wantObj SupportsL
	}{
		{
			name: "supports invalid decode",
			sL:   &SupportsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<supports element="accidental" type="yes"/>
		      <supports element="beam" type="yes"/>
		      <supports element="print" attribute="new-page" type="no"/>
		      <supports element="print" attribute="new-system" type="no"/>
		      <supports element="stem" type="yes"/>`))),
				start: xml.StartElement{Name: xml.Name{Local: "support"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: SupportsL{},
		},
		{
			name: "supports invalid type",
			sL:   &SupportsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="yo"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SupportsL{
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
			},
		},
		{
			name: "supports",
			sL:   &SupportsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="no"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>
	  <supports element="bend" type="yes" attribute="fast-beat" value="100"/>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SupportsL{
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
				{
					XMLName: xml.Name{
						Local: "supports",
					},
					Element:   "bend",
					Type:      enum.YesNo.Yes,
					Attribute: testutil.ToStringPtr("fast-beat"),
					Value:     testutil.ToStringPtr("100"),
				},
			},
		},
		{
			name: "supports empty",
			sL:   &SupportsL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SupportsL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SupportsL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.sL, tt.wantObj); diff != "" {
				t.Errorf("SupportsL.UnmarshalXML() value is mismatch (-sL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestSupportsL_MarshalXML(t *testing.T) {
	type args struct {
		sL *SupportsL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "supports",
			args: args{
				sL: &SupportsL{
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
					{
						XMLName: xml.Name{
							Local: "supports",
						},
						Element:   "bend",
						Type:      enum.YesNo.Yes,
						Attribute: testutil.ToStringPtr("fast-beat"),
						Value:     testutil.ToStringPtr("100"),
					},
				},
			},
			wantErr: false,
			wantXml: `<supports element="accidental" type="yes"></supports>
<supports element="beam" type="yes"></supports>
<supports element="print" type="no" attribute="new-page"></supports>
<supports element="print" type="no" attribute="new-system"></supports>
<supports element="stem" type="yes"></supports>
<supports element="bend" type="yes" attribute="fast-beat" value="100"></supports>`,
		},
		{
			name: "supports empty",
			args: args{
				sL: &SupportsL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.sL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("SupportsL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("SupportsL.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
