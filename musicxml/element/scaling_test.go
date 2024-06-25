package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestScaling_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		Millimeters *Millimeters
		Tenths      *Tenths
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
		wantObj Scaling
	}{
		{
			name:   "scaling invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<scaling>
	<millimeters>10.3921</millimeters>
	<tenths>10</tenths>
</scaling>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "scalin",
					},
				},
			},
			wantErr: true,
		},
		{
			name:   "scaling",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<scaling>
	<millimeters>10.3921</millimeters>
	<tenths>10</tenths>
</scaling>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Scaling{
				XMLName: xml.Name{
					Local: "scaling",
				},
				Millimeters: &Millimeters{
					XMLName: xml.Name{
						Local: "millimeters",
					},
					Millimeters: *datatype.NewMillimetersFixedPoint(103921, -4),
				},
				Tenths: &Tenths{
					XMLName: xml.Name{
						Local: "tenths",
					},
					Tenths: *datatype.NewTenthsFixedPoint(1, 1),
				},
			},
		},
		{
			name:   "scaling omit millimeters",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<scaling><tenths>10</tenths>
				</scaling>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Scaling{},
		},
		{
			name:   "scaling omit tenths",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<scaling><millimeters>10.3921</millimeters></scaling>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Scaling{},
		},
		{
			name:   "scaling empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Scaling{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sO := &Scaling{
				XMLName:     tt.fields.XMLName,
				Millimeters: tt.fields.Millimeters,
				Tenths:      tt.fields.Tenths,
			}
			if err := sO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Scaling.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*sO, tt.wantObj); diff != "" {
				t.Errorf("Scaling.UnmarshalXML() value is mismatch (-sO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestScaling_MarshalXML(t *testing.T) {
	type args struct {
		sO *Scaling
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "scaling",
			args: args{
				sO: &Scaling{
					XMLName: xml.Name{
						Local: "scaling",
					},
					Millimeters: &Millimeters{
						XMLName: xml.Name{
							Local: "millimeters",
						},
						Millimeters: *datatype.NewMillimetersFixedPoint(103921, -4),
					},
					Tenths: &Tenths{
						XMLName: xml.Name{
							Local: "tenths",
						},
						Tenths: *datatype.NewTenthsFromInt(50),
					},
				},
			},
			wantErr: false,
			wantXml: `<scaling>
	<millimeters>10.3921</millimeters>
	<tenths>50</tenths>
</scaling>`,
		},
		{
			name: "scaling omit millimeters",
			args: args{
				sO: &Scaling{
					XMLName: xml.Name{
						Local: "scaling",
					},
					Tenths: &Tenths{
						XMLName: xml.Name{
							Local: "tenths",
						},
						Tenths: *datatype.NewTenthsFromInt(10),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "scaling omit tenths",
			args: args{
				sO: &Scaling{
					XMLName: xml.Name{
						Local: "scaling",
					},
					Millimeters: &Millimeters{
						XMLName: xml.Name{
							Local: "millimeters",
						},
						Millimeters: *datatype.NewMillimetersFixedPoint(103921, -4),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name:    "scaling empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.sO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Scaling.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Scaling.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
