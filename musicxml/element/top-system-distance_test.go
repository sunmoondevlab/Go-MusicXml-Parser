package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestTopSystemDistance_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		TopSystemDistance datatype.Tenths
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
		wantObj TopSystemDistance
	}{
		{
			name:   "top-system-distance invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<top-system-distance>10</top-system-distance>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "top-system-distanc",
					},
				},
			},
			wantErr: true,
			wantObj: TopSystemDistance{},
		},
		{
			name:   "top-system-distance invalid decimal",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-system-distance>G10</top-system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: TopSystemDistance{},
		},
		{
			name:   "top-system-distance",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-system-distance>10</top-system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopSystemDistance{
				XMLName: xml.Name{
					Local: "top-system-distance",
				},
				TopSystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
			},
		},
		{
			name:   "top-system-distance empty value",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<top-system-distance></top-system-distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopSystemDistance{
				XMLName: xml.Name{
					Local: "top-system-distance",
				},
			},
		},
		{
			name:   "top-system-distance empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: TopSystemDistance{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tsdO := &TopSystemDistance{
				XMLName:           tt.fields.XMLName,
				TopSystemDistance: tt.fields.TopSystemDistance,
			}
			if err := tsdO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("TopSystemDistance.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tsdO, tt.wantObj); diff != "" {
				t.Errorf("TopSystemDistance.UnmarshalXML() value is mismatch (-tsdO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestTopSystemDistance_MarshalXML(t *testing.T) {
	type args struct {
		tsdO *TopSystemDistance
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "top-system-distance",
			args: args{
				tsdO: &TopSystemDistance{
					XMLName: xml.Name{
						Local: "top-system-distance",
					},
					TopSystemDistance: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
			wantErr: false,
			wantXml: `<top-system-distance>10</top-system-distance>`,
		},
		{
			name:    "top-system-distance empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.tsdO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("TopSystemDistance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("TopSystemDistance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
