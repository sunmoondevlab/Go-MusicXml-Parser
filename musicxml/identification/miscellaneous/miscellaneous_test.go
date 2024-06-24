package miscellaneous

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	tMiscellaneousfield "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/identification/miscellaneous/miscellaneousfield"
)

func TestMiscellaneous_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName            xml.Name
		MiscellaneousField *tMiscellaneousfield.MiscellaneousFieldL
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
			name:   "miscellaneous invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<miscellaneous><miscellaneous-field name="difficulty-level">3</miscellaneous-field>
    <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
    <miscellaneous-field name="parts-num">4</miscellaneous-field></miscellaneous>
`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "miscellaneou",
					},
					Attr: []xml.Attr{},
				},
			},
			wantErr: true,
			wantObj: Miscellaneous{},
		},
		{
			name:   "miscellaneous",
			fields: fields{},
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
					Local: "miscellaneous",
				},
				MiscellaneousField: &tMiscellaneousfield.MiscellaneousFieldL{
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
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mO := &Miscellaneous{
				XMLName:            tt.fields.XMLName,
				MiscellaneousField: tt.fields.MiscellaneousField,
			}
			if err := mO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Miscellaneous.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*mO, tt.wantObj); diff != "" {
				t.Errorf("Miscellaneous.UnmarshalXML() value is mismatch (-*mO +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestMiscellaneous_MarshalXML(t *testing.T) {
	type args struct {
		mO *Miscellaneous
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "miscellaneous",
			args: args{
				mO: &Miscellaneous{
					XMLName: xml.Name{
						Local: "miscellaneous",
					},
					MiscellaneousField: &tMiscellaneousfield.MiscellaneousFieldL{
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
			},
			wantErr: false,
			wantXml: `<miscellaneous>
  <miscellaneous-field name="difficulty-level">3</miscellaneous-field>
  <miscellaneous-field name="target-ages">older than 12</miscellaneous-field>
  <miscellaneous-field name="parts-num">4</miscellaneous-field>
</miscellaneous>`,
		},
		{
			name: "miscellaneous omit miscelaneous-field",
			args: args{
				mO: &Miscellaneous{
					XMLName: xml.Name{
						Local: "miscellaneous",
					},
				},
			},
			wantErr: false,
			wantXml: `<miscellaneous></miscellaneous>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.mO, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("Miscellaneous.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Miscellaneous.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}
