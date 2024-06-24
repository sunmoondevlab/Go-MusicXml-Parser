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
		mfL     *MiscellaneousFieldL
		args    args
		wantErr bool
		wantObj MiscellaneousFieldL
	}{
		{
			name: "miscellaneous-field invalid decode",

			mfL: &MiscellaneousFieldL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field>
`))),
				start: xml.StartElement{
					Name: xml.Name{
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

			mfL: &MiscellaneousFieldL{},
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
						Local: "miscellaneous-field",
					},
					Name:               "difficulty-level",
					MiscellaneousField: "3",
				},
				{
					XMLName: xml.Name{
						Local: "miscellaneous-field",
					},
					Name:               "target-ages",
					MiscellaneousField: "older than 12",
				},
				{
					XMLName: xml.Name{
						Local: "miscellaneous-field",
					},
					Name:               "parts-num",
					MiscellaneousField: "4",
				},
			},
		},
		{
			name: "miscellaneous-field",

			mfL: &MiscellaneousFieldL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: []MiscellaneousField{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.mfL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("MiscellaneousFieldL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.mfL, tt.wantObj); diff != "" {
				t.Errorf("MiscellaneousFieldL.UnmarshalXML() value is mismatch (-*tt.mfl +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestMiscellaneousFieldL_MarshalXML(t *testing.T) {
	type args struct {
		mfL *MiscellaneousFieldL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "miscellaneous-field",
			args: args{
				mfL: &MiscellaneousFieldL{
					{
						XMLName: xml.Name{
							Local: "miscellaneous-field",
						},
						Name:               "difficulty-level",
						MiscellaneousField: "3",
					},
					{
						XMLName: xml.Name{
							Local: "miscellaneous-field",
						},
						Name:               "target-ages",
						MiscellaneousField: "older than 12",
					},
					{
						XMLName: xml.Name{
							Local: "miscellaneous-field",
						},
						Name:               "parts-num",
						MiscellaneousField: "4",
					},
				},
			},
			wantErr: false,
			wantXml: `<miscellaneous-field name="difficulty-level">3</miscellaneous-field>
<miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
<miscellaneous-field name="parts-num">4</miscellaneous-field>`,
		},
		{
			name: "miscellaneous-field empty",
			args: args{
				mfL: &MiscellaneousFieldL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.mfL, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("MiscellaneousFieldL.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("MiscellaneousFieldL.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
