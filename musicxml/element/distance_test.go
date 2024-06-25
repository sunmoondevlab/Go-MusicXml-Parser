package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestDistanceL_UnmarshalXML(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
	}
	tests := []struct {
		name    string
		dL      *DistanceL
		args    args
		wantErr bool
		wantObj DistanceL
	}{
		{
			name: "distance invalid decode",
			dL:   &DistanceL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<distance type="beam">10</distance>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "distanc",
					},
				},
			},
			wantErr: true,
			wantObj: DistanceL{},
		},
		{
			name: "distance omit type",
			dL:   &DistanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<distance >10</distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: DistanceL{},
		},
		{
			name: "distance invalid type",
			dL:   &DistanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<distance type="bea">G10</distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: DistanceL{},
		},
		{
			name: "distance invalid decimal",
			dL:   &DistanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<distance type="beam">G10</distance>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: DistanceL{},
		},
		{
			name: "distance",
			dL:   &DistanceL{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<distance type="hyphen">10</distance>
<distance type="beam">1</distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: DistanceL{
				{
					XMLName: xml.Name{
						Local: "distance",
					},
					Type:     &enum.DistanceType.Hyphen,
					Distance: *datatype.NewTenthsFixedPoint(1, 1),
				},
				{
					XMLName: xml.Name{
						Local: "distance",
					},
					Type:     &enum.DistanceType.Beam,
					Distance: *datatype.NewTenthsFixedPoint(1, 0),
				},
			},
		},
		{
			name: "distance empty value",
			dL:   &DistanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<distance type="beam"></distance>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: DistanceL{
				{
					XMLName: xml.Name{
						Local: "distance",
					},
					Type: &enum.DistanceType.Beam,
				},
			},
		},
		{
			name: "distance empty",
			dL:   &DistanceL{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: DistanceL{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.dL.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("DistanceL.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.dL, tt.wantObj); diff != "" {
				t.Errorf("DistanceL.UnmarshalXML() value is mismatch (-dL, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestDistanceL_MarshalXML(t *testing.T) {
	type args struct {
		dL *DistanceL
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "distance omit type",
			args: args{
				dL: &DistanceL{
					{
						XMLName: xml.Name{
							Local: "distance",
						},
						Distance: *datatype.NewTenthsFixedPoint(1, 1),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "distance",
			args: args{
				dL: &DistanceL{
					{
						XMLName: xml.Name{
							Local: "distance",
						},
						Type:     &enum.DistanceType.Beam,
						Distance: *datatype.NewTenthsFixedPoint(11, 0),
					},
					{
						XMLName: xml.Name{
							Local: "distance",
						},
						Type:     &enum.DistanceType.Hyphen,
						Distance: *datatype.NewTenthsFixedPoint(1, 1),
					},
				},
			},
			wantErr: false,
			wantXml: `<distance type="beam">11</distance>
<distance type="hyphen">10</distance>`,
		},
		{
			name: "distance empty",
			args: args{
				dL: &DistanceL{},
			},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.dL, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Distance.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Distance.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
