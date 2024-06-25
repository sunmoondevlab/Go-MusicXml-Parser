package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestConcertScore_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName xml.Name
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
		wantObj ConcertScore
	}{
		{
			name:   "concert-score invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<concert-score/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "concert-scor",
					},
				},
			},
			wantErr: true,
			wantObj: ConcertScore{},
		},
		{
			name:   "concert-score",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<concert-score/>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: ConcertScore{
				XMLName: xml.Name{
					Local: "concert-score",
				},
			},
		},
		{
			name:   "concert-score empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: ConcertScore{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			csO := &ConcertScore{
				XMLName: tt.fields.XMLName,
			}
			if err := csO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("ConcertScore.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*csO, tt.wantObj); diff != "" {
				t.Errorf("ConcertScore.UnmarshalXML() value is mismatch (-csO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestConcertScore_MarshalXML(t *testing.T) {
	type args struct {
		csO *ConcertScore
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "concert-score",
			args: args{
				csO: &ConcertScore{
					XMLName: xml.Name{
						Local: "concert-score",
					},
				},
			},
			wantErr: false,
			wantXml: `<concert-score></concert-score>`,
		},
		{
			name:    "concert-score empty",
			args:    args{},
			wantErr: false,
			wantXml: ``,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.csO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("ConcertScore.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("ConcertScore.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
