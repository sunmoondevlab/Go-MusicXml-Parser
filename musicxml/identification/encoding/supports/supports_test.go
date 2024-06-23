package supports

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestSupportsL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		sl      *SupportsL
		args    args
		wantErr bool
		wantObj SupportsL
	}{
		{
			name: "supports invalid decode",
			sl:   &SupportsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<supports element="accidental" type="yes"/>
		      <supports element="beam" type="yes"/>
		      <supports element="print" attribute="new-page" type="no"/>
		      <supports element="print" attribute="new-system" type="no"/>
		      <supports element="stem" type="yes"/>`))),
				start: xml.StartElement{Name: xml.Name{Space: "", Local: "support"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: []Supports{},
		},
		{
			name: "supports invalid type",
			sl:   &SupportsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<supports element="accidental" type="yes"/>
      <supports element="beam" type="yes"/>
      <supports element="print" attribute="new-page" type="yo"/>
      <supports element="print" attribute="new-system" type="no"/>
      <supports element="stem" type="yes"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: []Supports{
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
			},
		},
		{
			name: "supports",
			sl:   &SupportsL{},
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
			wantObj: []Supports{
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
				{
					XMLName: xml.Name{
						Space: "",
						Local: "supports",
					},
					Element:   "bend",
					Type:      enum.YesNo.Yes,
					Attribute: "fast-beat",
					Value:     "100",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.sl.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SupportsL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.sl, tt.wantObj); diff != "" {
				t.Errorf("SupportsL.UnmarshalXML() value is mismatch (-*tt.sl +tt.wantObj):%s\n", diff)
			}
		})
	}
}
