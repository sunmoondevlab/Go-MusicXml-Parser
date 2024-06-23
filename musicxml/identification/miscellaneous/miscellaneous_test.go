package miscellaneous

import (
	"bytes"
	"encoding/xml"
	"testing"

	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
)

func TestMiscellaneous_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName            xml.Name
		MiscellaneousField tMiscellaneousfield.MiscellaneousFieldL
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
		wantObj Miscellaneous
	}{
		{
			name: "miscellaneous invalid decode",
			fields: fields{
				XMLName:            xml.Name{},
				MiscellaneousField: []tMiscellaneousfield.MiscellaneousField{},
			},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous><miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field></miscellaneous>
`))),
				start: xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "miscellaneou",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: Miscellaneous{},
		},
		{
			name: "miscellaneous",
			fields: fields{
				XMLName:            xml.Name{},
				MiscellaneousField: []tMiscellaneousfield.MiscellaneousField{},
			},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous><miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field></miscellaneous>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Miscellaneous{
				XMLName: xml.Name{
					Space: "",
					Local: "miscellaneous",
				},
				MiscellaneousField: []tMiscellaneousfield.MiscellaneousField{
					{
						XMLName: xml.Name{
							Space: "",
							Local: "miscellaneous-field",
						},
						Name:  "difficulty-level",
						Value: "3",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "miscellaneous-field",
						},
						Name:  "target-ages",
						Value: "older than 12",
					},
					{
						XMLName: xml.Name{
							Space: "",
							Local: "miscellaneous-field",
						},
						Name:  "parts-num",
						Value: "4",
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Miscellaneous{
				XMLName:            tt.fields.XMLName,
				MiscellaneousField: tt.fields.MiscellaneousField,
			}
			if err := m.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Miscellaneous.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
