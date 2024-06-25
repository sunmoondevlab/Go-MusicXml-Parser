package rights

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestRightsL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		rL      *RightsL
		args    args
		wantErr bool
		wantObj RightsL
	}{
		{
			name: "rights invalid decode",
			rL:   &RightsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<rights>Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "right",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: []Rights{},
		},
		{
			name: "rights",
			rL:   &RightsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<rights>Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
    <rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>`))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []Rights{
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           nil,
					CopyRightsType: nil,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("music"),
					CopyRightsType: &enum.CopyRightsType.Music,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("arrangement"),
					CopyRightsType: &enum.CopyRightsType.Arrangement,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("words"),
					CopyRightsType: &enum.CopyRightsType.Words,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("translation"),
					CopyRightsType: &enum.CopyRightsType.Translation,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("parody"),
					CopyRightsType: &enum.CopyRightsType.Parody,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Local: "rights",
					},
					Type:           testutil.ToStringPtr("???"),
					CopyRightsType: &enum.CopyRightsType.Other,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
			},
		},
		{
			name: "rights empty",
			rL:   &RightsL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []Rights{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RightsL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.rL, tt.wantObj); diff != "" {
			t.Errorf("RightsL.UnmarshalXML() value is mismatch (-*tt.rL +tt.wantObj):%s\n", diff)
		}
	}
}

func TestRightsL_MarshalXML(t *testing.T) {
	type args struct {
		rL *RightsL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "rights",
			args: args{
				rL: &RightsL{
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           nil,
						CopyRightsType: nil,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("music"),
						CopyRightsType: &enum.CopyRightsType.Music,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("arrangement"),
						CopyRightsType: &enum.CopyRightsType.Arrangement,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("words"),
						CopyRightsType: &enum.CopyRightsType.Words,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("translation"),
						CopyRightsType: &enum.CopyRightsType.Translation,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("parody"),
						CopyRightsType: &enum.CopyRightsType.Parody,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
					{
						XMLName: xml.Name{
							Local: "rights",
						},
						Type:           testutil.ToStringPtr("???"),
						CopyRightsType: &enum.CopyRightsType.Other,
						CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
					},
				},
			},
			wantErr: false,
			wantXml: `<rights>Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="music">Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="arrangement">Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="words">Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="translation">Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="parody">Copyright © 2024 Hogeran A.Fugas</rights>
<rights type="???">Copyright © 2024 Hogeran A.Fugas</rights>`,
		},
		{
			name: "rights empty",
			args: args{
				rL: &RightsL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.rL, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("RightsL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("RightsL.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
