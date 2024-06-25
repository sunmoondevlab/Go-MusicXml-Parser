package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestWork_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName    xml.Name
		WorkTitle  *string
		WorkNumber *string
		Opus       *Opus
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
				start: xml.StartElement{Name: xml.Name{Local: "wor"}, Attr: []xml.Attr{}},
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
					Local: "work",
				},
				WorkTitle:  testutil.ToStringPtr("無題のスコア"),
				WorkNumber: testutil.ToStringPtr("Work Number"),
				Opus: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
					Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
					Type:    &enum.XlinkType.Simple,
					Role:    testutil.ToStringPtr("role"),
					Title:   testutil.ToStringPtr("winterreise"),
					Show:    &enum.XlinkShow.New,
					Actuate: &enum.XlinkActuate.None,
				},
			},
		},
		{
			name:   "work omit work-number",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    <work-title>無題のスコア</work-title>
    <opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>
    </work>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{
				XMLName: xml.Name{
					Local: "work",
				},
				WorkTitle: testutil.ToStringPtr("無題のスコア"),
				Opus: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
					Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
					Type:    &enum.XlinkType.Simple,
					Role:    testutil.ToStringPtr("role"),
					Title:   testutil.ToStringPtr("winterreise"),
					Show:    &enum.XlinkShow.New,
					Actuate: &enum.XlinkActuate.None,
				},
			},
		},
		{
			name:   "work omit work-title",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    <work-number>Work Number</work-number>
    <opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>
    </work>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{
				XMLName: xml.Name{
					Local: "work",
				},
				WorkNumber: testutil.ToStringPtr("Work Number"),
				Opus: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
					Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
					Type:    &enum.XlinkType.Simple,
					Role:    testutil.ToStringPtr("role"),
					Title:   testutil.ToStringPtr("winterreise"),
					Show:    &enum.XlinkShow.New,
					Actuate: &enum.XlinkActuate.None,
				},
			},
		},
		{
			name:   "work omit opus",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    <work-number>Work Number</work-number>
    <work-title>無題のスコア</work-title>
	</work>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{
				XMLName: xml.Name{
					Local: "work",
				},
				WorkTitle:  testutil.ToStringPtr("無題のスコア"),
				WorkNumber: testutil.ToStringPtr("Work Number"),
			},
		},
		{
			name:   "work omit all children",
			fields: fields{},
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<work>
    </work>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{
				XMLName: xml.Name{
					Local: "work",
				},
			},
		},
		{
			name:   "work empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Work{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wO := &Work{
				XMLName:    tt.fields.XMLName,
				WorkTitle:  tt.fields.WorkTitle,
				WorkNumber: tt.fields.WorkNumber,
				Opus:       tt.fields.Opus,
			}
			if err := wO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Work.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(wO, &tt.wantObj); diff != "" {
				t.Errorf("Work.UnmarshalXML() value is mismatch (-wO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestWork_MarshalXML(t *testing.T) {
	type args struct {
		wO *Work
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "work",
			args: args{
				wO: &Work{
					XMLName: xml.Name{
						Local: "work",
					},
					WorkTitle:  testutil.ToStringPtr("無題のスコア"),
					WorkNumber: testutil.ToStringPtr("Work Number"),
					Opus: &Opus{
						XMLName: xml.Name{
							Local: "opus",
						},
						Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
						Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
						Type:    &enum.XlinkType.Simple,
						Role:    testutil.ToStringPtr("role"),
						Title:   testutil.ToStringPtr("winterreise"),
						Show:    &enum.XlinkShow.New,
						Actuate: &enum.XlinkActuate.None,
					},
				},
			},
			wantErr: false,
			wantXml: `<work>
	<work-title>無題のスコア</work-title>
	<work-number>Work Number</work-number>
	<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"></opus>
</work>`,
		},
		{
			name: "work ommit work-title",
			args: args{
				wO: &Work{
					XMLName: xml.Name{
						Local: "work",
					},
					WorkNumber: testutil.ToStringPtr("Work Number"),
					Opus: &Opus{
						XMLName: xml.Name{
							Local: "opus",
						},
						Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
						Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
						Type:    &enum.XlinkType.Simple,
						Role:    testutil.ToStringPtr("role"),
						Title:   testutil.ToStringPtr("winterreise"),
						Show:    &enum.XlinkShow.New,
						Actuate: &enum.XlinkActuate.None,
					},
				},
			},
			wantErr: false,
			wantXml: `<work>
	<work-number>Work Number</work-number>
	<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"></opus>
</work>`,
		},
		{
			name: "work omit work-number",
			args: args{
				wO: &Work{
					XMLName: xml.Name{
						Local: "work",
					},
					WorkTitle: testutil.ToStringPtr("無題のスコア"),
					Opus: &Opus{
						XMLName: xml.Name{
							Local: "opus",
						},
						Xlink:   testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
						Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
						Type:    &enum.XlinkType.Simple,
						Role:    testutil.ToStringPtr("role"),
						Title:   testutil.ToStringPtr("winterreise"),
						Show:    &enum.XlinkShow.New,
						Actuate: &enum.XlinkActuate.None,
					},
				},
			},
			wantErr: false,
			wantXml: `<work>
	<work-title>無題のスコア</work-title>
	<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"></opus>
</work>`,
		},
		{
			name: "work omit opus",
			args: args{
				wO: &Work{
					XMLName: xml.Name{
						Local: "work",
					},
					WorkTitle:  testutil.ToStringPtr("無題のスコア"),
					WorkNumber: testutil.ToStringPtr("Work Number"),
				},
			},
			wantErr: false,
			wantXml: `<work>
	<work-title>無題のスコア</work-title>
	<work-number>Work Number</work-number>
</work>`,
		},
		{
			name: "work omit all children",
			args: args{
				wO: &Work{
					XMLName: xml.Name{
						Local: "work",
					},
				},
			},
			wantErr: false,
			wantXml: `<work></work>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.wO, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Work.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Work.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
