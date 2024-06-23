package rights

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestRightsL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		rl      *RightsL
		args    args
		wantErr bool
		wantObj RightsL
	}{
		{
			name: "rights invalid decode",
			rl:   &RightsL{},
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
						Space: "",
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
			rl:   &RightsL{},
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
						Space: "",
						Local: "rights",
					},
					Type:           "",
					CopyRightsType: enum.CopyRightsType.All,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "music",
					CopyRightsType: enum.CopyRightsType.Music,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "arrangement",
					CopyRightsType: enum.CopyRightsType.Arrangement,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "words",
					CopyRightsType: enum.CopyRightsType.Words,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "translation",
					CopyRightsType: enum.CopyRightsType.Translation,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "parody",
					CopyRightsType: enum.CopyRightsType.Parody,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
				{
					XMLName: xml.Name{
						Space: "",
						Local: "rights",
					},
					Type:           "???",
					CopyRightsType: enum.CopyRightsType.Other,
					CopyRight:      "Copyright © 2024 Hogeran A.Fugas",
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.rl.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("RightsL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
		if diff := cmp.Diff(*tt.rl, tt.wantObj); diff != "" {
			t.Errorf("RightsL.UnmarshalXML() value is mismatch (-*tt.rl +tt.wantObj):%s\n", diff)
		}
	}
}
