package miscellaneousfield

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestMiscellaneousFieldL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		mfl     *MiscellaneousFieldL
		args    args
		wantErr bool
		wantObj MiscellaneousFieldL
	}{
		{
			name: "miscellaneous-field invalid decode",

			mfl: &MiscellaneousFieldL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field>
`))),
				start: xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "miscellaneous-fiel",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: []MiscellaneousField{},
		},
		{
			name: "miscellaneous-field",

			mfl: &MiscellaneousFieldL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field>
`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []MiscellaneousField{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mfl.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("MiscellaneousFieldL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.mfl, tt.wantObj); diff != "" {
				t.Errorf("MiscellaneousFieldL.UnmarshalXML() value is mismatch (-*tt.mfl +tt.wantObj):%s\n", diff)
			}
		})
	}
}
