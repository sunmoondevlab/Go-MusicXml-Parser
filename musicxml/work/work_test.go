package work

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	tWorkOpus "github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/work/workopus"
)

func TestWork_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName    xml.Name
		WorkTitle  string
		WorkNumber string
		WorkOpus   tWorkOpus.WorkOpus
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
		wantObj Work
	}{
		{
			name:   "work invalid decode",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    <work-number>Work Number</work-number>
    <work-title>無題のスコア</work-title>
    <opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>
    </work>`))),
				start: xml.StartElement{Name: xml.Name{Space: "", Local: "wor"}, Attr: []xml.Attr{}},
			},
			wantErr: true,
			wantObj: Work{},
		},
		{
			name:   "work",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    <work-number>Work Number</work-number>
    <work-title>無題のスコア</work-title>
    <opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>
    </work>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{
				XMLName: xml.Name{
					Space: "",
					Local: "work",
				},
				WorkTitle:  "無題のスコア",
				WorkNumber: "Work Number",
				WorkOpus: tWorkOpus.WorkOpus{
					XMLName: xml.Name{
						Space: "",
						Local: "opus",
					},
					XLink:   "http://www.w3.org/1999/xlink",
					Href:    "opus/winterreise.musicxml",
					Type:    enum.XlinkType.Simple,
					Role:    "role",
					Title:   "winterreise",
					Show:    enum.XlinkShow.New,
					Actuate: enum.XlinkActuate.None,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &Work{
				XMLName:    tt.fields.XMLName,
				WorkTitle:  tt.fields.WorkTitle,
				WorkNumber: tt.fields.WorkNumber,
				WorkOpus:   tt.fields.WorkOpus,
			}
			if err := w.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Work.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(w, &tt.wantObj); diff != "" {
				t.Errorf("Work.UnmarshalXML() value is mismatch (-w +tt.wantObj):%s\n", diff)
			}
		})
	}
}
