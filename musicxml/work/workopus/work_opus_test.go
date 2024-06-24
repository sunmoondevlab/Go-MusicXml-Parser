package workopus

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/attrgrp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestWorkOpus_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		XLink   string
		Href    string
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
		wantObj WorkOpus
	}{
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
			wantObj: WorkOpus{
				XMLName: xml.Name{
					Local: "opus",
				},
				XLink:   "http://www.w3.org/1999/xlink",
				Href:    "opus/winterreise.musicxml",
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
			wantObj: WorkOpus{
				XMLName: xml.Name{
					Local: "opus",
				},
				XLink: "http://www.w3.org/1999/xlink",
				Href:  "opus/winterreise.musicxml",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			woo := &WorkOpus{
				XMLName: tt.fields.XMLName,
				XLink:   tt.fields.XLink,
				Href:    tt.fields.Href,
				Type:    tt.fields.Type,
				Role:    tt.fields.Role,
				Title:   tt.fields.Title,
				Show:    tt.fields.Show,
				Actuate: tt.fields.Actuate,
			}
			if err := woo.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("WorkOpus.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(woo, &tt.wantObj); diff != "" {
				t.Errorf("WorkOpus.UnmarshalXML() value is mismatch (-woo +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_MarshalXML(t *testing.T) {
	type args struct {
		woo *WorkOpus
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantXml string
	}{
		{
			name: "work>opus",
			args: args{
				woo: &WorkOpus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Href:    "opus/winterreise.musicxml",
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
			name: "work>opus ommit empty",
			args: args{
				woo: &WorkOpus{
					XMLName: xml.Name{
						Local: "opus",
					},
					Href:  "opus/winterreise.musicxml",
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
			b, err := xml.MarshalIndent(tt.args.woo, "", "  ")
			if (err != nil) != tt.wantErr {
				t.Errorf("WorkOpus.MarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			ox := string(b)
			if diff := cmp.Diff(ox, tt.wantXml); diff != "" {
				t.Errorf("WorkOpus.MarshalXML() value is mismatch (-ox +tt.wantXml):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetXMLName(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   xml.Name
	}{
		{
			name: "get",
			fields: fields{
				XMLName: xml.Name{
					Local: "opus",
				},
			},
			want: xml.Name{
				Local: "opus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetXMLName()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetXMLName() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetXLink(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get",
			fields: fields{
				XLink: "xlink",
			},
			want: "xlink",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetXLink()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetXLink() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetHref(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get",
			fields: fields{
				Href: "href",
			},
			want: "href",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetHref()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetHref() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetType(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   *enum.XlinkTypeEnum
	}{
		{
			name: "get",
			fields: fields{
				Type: &enum.XlinkType.Simple,
			},
			want: &enum.XlinkType.Simple,
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetType()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetType() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetRole(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Role: testutil.ToStringPtr("role"),
			},
			want: testutil.ToStringPtr("role"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetRole()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetRole() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetTitle(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Title: testutil.ToStringPtr("title"),
			},
			want: testutil.ToStringPtr("title"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetTitle()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetTitle() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetShow(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   *enum.XlinkShowEnum
	}{
		{
			name: "get",
			fields: fields{
				Show: &enum.XlinkShow.New,
			},
			want: &enum.XlinkShow.New,
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetShow()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetShow() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_GetActuate(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	tests := []struct {
		name   string
		fields fields
		want   *enum.XlinkActuateEnum
	}{
		{
			name: "get",
			fields: fields{
				Actuate: &enum.XlinkActuate.OnRequest,
			},
			want: &enum.XlinkActuate.OnRequest,
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			got := l.GetActuate()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpus.GetActuate() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetXMLName(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xn xml.Name
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   xml.Name
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xn: xml.Name{Local: "opus"}},
			want: xml.Name{
				Local: "opus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetXMLName(tt.args.xn)
			if diff := cmp.Diff(l.XMLName, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetXMLName() value is mismatch (-l.XMLName +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetXLink(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xl: "xlink"},
			want:   "xlink",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetXLink(tt.args.xl)
			if diff := cmp.Diff(l.XLink, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetXLink() value is mismatch (-l.Xlink +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetHref(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		h string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{h: "href"},
			want:   "href",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetHref(tt.args.h)
			if diff := cmp.Diff(l.Href, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetHref() value is mismatch (-l.Href +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetType(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xt *enum.XlinkTypeEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *enum.XlinkTypeEnum
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xt: &enum.XlinkType.Simple},
			want:   &enum.XlinkType.Simple,
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetType(tt.args.xt)
			if diff := cmp.Diff(l.Type, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetType() value is mismatch (-l.Type +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetRole(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		r *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{r: testutil.ToStringPtr("role")},
			want:   testutil.ToStringPtr("role"),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetRole(tt.args.r)
			if diff := cmp.Diff(l.Role, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetRole() value is mismatch (-l.Role +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetTitle(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		t *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{t: testutil.ToStringPtr("title")},
			want:   testutil.ToStringPtr("title"),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetTitle(tt.args.t)
			if diff := cmp.Diff(l.Title, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetTitle() value is mismatch (-l.Title +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetShow(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xs *enum.XlinkShowEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *enum.XlinkShowEnum
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xs: &enum.XlinkShow.New},
			want:   &enum.XlinkShow.New,
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetShow(tt.args.xs)
			if diff := cmp.Diff(l.Show, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetShow() value is mismatch (-l.Show +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetActuate(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           *enum.XlinkTypeEnum
		Role           *string
		Title          *string
		Show           *enum.XlinkShowEnum
		Actuate        *enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xa *enum.XlinkActuateEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *enum.XlinkActuateEnum
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xa: &enum.XlinkActuate.OnRequest},
			want:   &enum.XlinkActuate.OnRequest,
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpus{
				XMLName:        tt.fields.XMLName,
				XLink:          tt.fields.XLink,
				Href:           tt.fields.Href,
				Type:           tt.fields.Type,
				Role:           tt.fields.Role,
				Title:          tt.fields.Title,
				Show:           tt.fields.Show,
				Actuate:        tt.fields.Actuate,
				LinkAttributes: tt.fields.LinkAttributes,
			}
			l.SetActuate(tt.args.xa)
			if diff := cmp.Diff(l.Actuate, tt.want); diff != "" {
				t.Errorf("WorkOpus.SetActuate() value is mismatch (-l.Actuate +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetXMLName(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   xml.Name
	}{
		{
			name: "get",
			fields: fields{
				XMLName: xml.Name{
					Local: "opus",
				},
			},
			want: xml.Name{
				Local: "opus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetXMLName()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetXMLName() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetXLink(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get",
			fields: fields{
				XLink: "xlink",
			},
			want: "xlink",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetXLink()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetXLink() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetHref(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "get",
			fields: fields{
				Href: "href",
			},
			want: "href",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetHref()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetHref() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetType(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Type: testutil.ToStringPtr("simple"),
			},
			want: testutil.ToStringPtr("simple"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetType()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetType() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetRole(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Role: testutil.ToStringPtr("role"),
			},
			want: testutil.ToStringPtr("role"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetRole()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetRole() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetTitle(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Title: testutil.ToStringPtr("title"),
			},
			want: testutil.ToStringPtr("title"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetTitle()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetTitle() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetShow(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Show: testutil.ToStringPtr("new"),
			},
			want: testutil.ToStringPtr("new"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetShow()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetShow() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_GetActuate(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
	}{
		{
			name: "get",
			fields: fields{
				Actuate: testutil.ToStringPtr("onRequest"),
			},
			want: testutil.ToStringPtr("onRequest"),
		},
		{
			name:   "get nil",
			fields: fields{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			l := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			got := l.GetActuate()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.GetActuate() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetXMLName(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		xn xml.Name
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   xml.Name
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xn: xml.Name{Local: "opus"}},
			want: xml.Name{
				Local: "opus",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetXMLName(tt.args.xn)
			if diff := cmp.Diff(lr.XMLName, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetXMLName() value is mismatch (-lr.XMLName +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetXLink(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		xl string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xl: "xlink"},
			want:   "xlink",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetXLink(tt.args.xl)
			if diff := cmp.Diff(lr.XLink, tt.want); diff != "" {
				t.Errorf("WorkOpuRaws.SetXLink() value is mismatch (-lr.Xlink +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetHref(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		h string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{h: "href"},
			want:   "href",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetHref(tt.args.h)
			if diff := cmp.Diff(lr.Href, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetHref() value is mismatch (-lr.Href +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetType(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		xt *enum.XlinkTypeEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xt: &enum.XlinkType.Simple},
			want:   enum.XlinkType.Simple.StringPtr(),
		},
		{
			name:   "set",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetType(tt.args.xt)
			if diff := cmp.Diff(lr.Type, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetType() value is mismatch (-lr.Type +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetRole(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		r *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{r: testutil.ToStringPtr("role")},
			want:   testutil.ToStringPtr("role"),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetRole(tt.args.r)
			if diff := cmp.Diff(lr.Role, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetRole() value is mismatch (-lr.Role +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetTitle(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		t *string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{t: testutil.ToStringPtr("title")},
			want:   testutil.ToStringPtr("title"),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetTitle(tt.args.t)
			if diff := cmp.Diff(lr.Title, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetTitle() value is mismatch (-lr.Title +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetShow(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		xs *enum.XlinkShowEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xs: &enum.XlinkShow.New},
			want:   enum.XlinkShow.New.StringPtr(),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetShow(tt.args.xs)
			if diff := cmp.Diff(lr.Show, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetShow() value is mismatch (-lr.Show +tt.want):%s\n", diff)
			}
		})
	}
}

func TestWorkOpusRaw_SetActuate(t *testing.T) {
	type fields struct {
		XMLName           xml.Name
		XLink             string
		Href              string
		Type              *string
		Role              *string
		Title             *string
		Show              *string
		Actuate           *string
		LinkAttributesRaw attrgrp.LinkAttributesRaw
	}
	type args struct {
		xa *enum.XlinkActuateEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   *string
	}{
		{
			name:   "set",
			fields: fields{},
			args:   args{xa: &enum.XlinkActuate.OnRequest},
			want:   enum.XlinkActuate.OnRequest.StringPtr(),
		},
		{
			name:   "set nil",
			fields: fields{},
			args:   args{},
			want:   nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			lr := &WorkOpusRaw{
				XMLName:           tt.fields.XMLName,
				XLink:             tt.fields.XLink,
				Href:              tt.fields.Href,
				Type:              tt.fields.Type,
				Role:              tt.fields.Role,
				Title:             tt.fields.Title,
				Show:              tt.fields.Show,
				Actuate:           tt.fields.Actuate,
				LinkAttributesRaw: tt.fields.LinkAttributesRaw,
			}
			lr.SetActuate(tt.args.xa)
			if diff := cmp.Diff(lr.Actuate, tt.want); diff != "" {
				t.Errorf("WorkOpusRaw.SetActuate() value is mismatch (-lr.Actuate +tt.want):%s\n", diff)
			}
		})
	}
}
