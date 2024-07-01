package element

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestOpus_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		Xlink   *string
		Href    *string
		Type    *enum.XlinkTypeEnum
		Role    *string
		Title   *string
		Show    *enum.XlinkShowEnum
		Actuate *enum.XlinkActuateEnum
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
		wantObj Opus
	}{
		{
			name: "opus invalid decode",
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simpl" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Local: "opu",
					},
				},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus omit xlink",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus omit href",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus invalid type",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simpl" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus invalid show",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="ne" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus invalid actuate",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="non"/>`))),
				start: xml.StartElement{},
			},
			wantErr: true,
			wantObj: Opus{},
		},
		{
			name: "opus",
			fields: fields{
				XMLName: xml.Name{Local: "opus"},
			},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Opus{
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
		{
			name: "opus omit optional",
			fields: fields{
				XMLName: xml.Name{Local: "opus"},
			},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" />`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Opus{
				XMLName: xml.Name{
					Local: "opus",
				},
				Xlink: testutil.ToStringPtr("http://www.w3.org/1999/xlink"),
				Href:  testutil.ToStringPtr("opus/winterreise.musicxml"),
			},
		},
		{
			name:   "opus empty",
			fields: fields{},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(``))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: Opus{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			oO := &Opus{
				XMLName: tt.fields.XMLName,
				Xlink:   tt.fields.Xlink,
				Href:    tt.fields.Href,
				Type:    tt.fields.Type,
				Role:    tt.fields.Role,
				Title:   tt.fields.Title,
				Show:    tt.fields.Show,
				Actuate: tt.fields.Actuate,
			}
			if err := oO.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("Opus.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(oO, &tt.wantObj); diff != "" {
				t.Errorf("Opus.UnmarshalXML() value is mismatch (-oO, +wantObj):%s\n", diff)
			}
		})
	}
}

func TestOpus_MarshalXML(t *testing.T) {
	type args struct {
		woo *Opus
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "opus omit href",
			args: args{
				woo: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Type:    &enum.XlinkType.Simple,
					Role:    testutil.ToStringPtr("role"),
					Title:   testutil.ToStringPtr("winterreise"),
					Show:    &enum.XlinkShow.New,
					Actuate: &enum.XlinkActuate.None,
				},
			},
			wantErr: true,
			wantXml: ``,
		},
		{
			name: "opus",
			args: args{
				woo: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Href:    testutil.ToStringPtr("opus/winterreise.musicxml"),
					Type:    &enum.XlinkType.Simple,
					Role:    testutil.ToStringPtr("role"),
					Title:   testutil.ToStringPtr("winterreise"),
					Show:    &enum.XlinkShow.New,
					Actuate: &enum.XlinkActuate.None,
				},
			},
			wantErr: false,
			wantXml: `<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"></opus>`,
		},
		{
			name: "opus ommit empty",
			args: args{
				woo: &Opus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Href:  testutil.ToStringPtr("opus/winterreise.musicxml"),
					Role:  testutil.ToStringPtr("role"),
					Title: testutil.ToStringPtr("winterreise"),
				},
			},
			wantErr: false,
			wantXml: `<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:role="role" xlink:title="winterreise"></opus>`,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			b, err := xml.MarshalIndent(tt.args.woo, "", "\t")
			if (err != nil) != tt.wantErr {
				t.Errorf("Opus.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("Opus.MarshalXML() value is mismatch (-ox, +wantXml):%s\n", diff)
			}
		})
	}
}
