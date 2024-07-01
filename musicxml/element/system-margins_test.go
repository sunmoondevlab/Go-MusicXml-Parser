package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
)

func TestSystemMargins_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName     xml.Name
		LeftMargin  *LeftMargin
		RightMargin *RightMargin
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
		wantObj SystemMargins
	}{
		{
			name:   "system-margins invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
</system-margins>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "system-margin",
					},
				},
			},
			wantErr: true,
			wantObj: SystemMargins{},
		},
		{
			name:   "system-margins omit left-margin",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-margins>
	<right-margin>85.7253</right-margin>
</system-margins>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SystemMargins{},
		},
		{
			name:   "system-margins omit right-margin",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-margins>
	<left-margin>85.7252</left-margin>
</system-margins>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: SystemMargins{},
		},
		{
			name:   "system-margins",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<system-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
</system-margins>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemMargins{
				XMLName: xml.Name{
					Local: "system-margins",
				},
				LeftMargin: &LeftMargin{
					XMLName: xml.Name{
						Local: "left-margin",
					},
					LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
				},
				RightMargin: &RightMargin{
					XMLName: xml.Name{
						Local: "right-margin",
					},
					RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
				},
			},
		},
		{
			name:   "system-margins empty",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(``))), start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: SystemMargins{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			smO := &SystemMargins{
				XMLName:     tt.fields.XMLName,
				LeftMargin:  tt.fields.LeftMargin,
				RightMargin: tt.fields.RightMargin,
			}
			if err := smO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("SystemMargins.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*smO, tt.wantObj); diff != "" {
				t.Errorf("SystemMargins.UnmarshalXML() value is mismatch (-smO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestSystemMargins_MarshalXML(t *testing.T) {
	type args struct {
		smO *SystemMargins
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "system-margins omit left-margin",
			args: args{
				smO: &SystemMargins{
					XMLName: xml.Name{
						Local: "system-margins",
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "system-margins omit right-margin",
			args: args{
				smO: &SystemMargins{
					XMLName: xml.Name{
						Local: "system-margins",
					},
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "system-margins",
			args: args{
				smO: &SystemMargins{
					XMLName: xml.Name{
						Local: "system-margins",
					},
					LeftMargin: &LeftMargin{
						XMLName: xml.Name{
							Local: "left-margin",
						},
						LeftMargin: *datatype.NewTenthsFixedPoint(857252, -4),
					},
					RightMargin: &RightMargin{
						XMLName: xml.Name{
							Local: "right-margin",
						},
						RightMargin: *datatype.NewTenthsFixedPoint(857253, -4),
					},
				},
			},
			wantErr: false,
			wantXml: `<system-margins>
	<left-margin>85.7252</left-margin>
	<right-margin>85.7253</right-margin>
</system-margins>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.smO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("SystemMargins.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("SystemMargins.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
