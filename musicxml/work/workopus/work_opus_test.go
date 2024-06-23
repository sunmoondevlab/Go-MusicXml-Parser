package workopus

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/attrgrp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestWorkOpus_UnmarshalXML(t *testing.T) {
	type fields struct {
		XMLName xml.Name
		XLink   string
		Href    string
		Type    enum.XlinkTypeEnum
		Role    string
		Title   string
		Show    enum.XlinkShowEnum
		Actuate enum.XlinkActuateEnum
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
				XMLName: xml.Name{Space: "", Local: "opus"},
				XLink:   "",
				Href:    "",
				Type:    "",
				Role:    "",
				Title:   "",
				Show:    "",
				Actuate: "",
			},
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
			},
			wantErr: false,
			wantObj: WorkOpus{
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
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			wo := &WorkOpus{
				XMLName: tt.fields.XMLName,
				XLink:   tt.fields.XLink,
				Href:    tt.fields.Href,
				Type:    tt.fields.Type,
				Role:    tt.fields.Role,
				Title:   tt.fields.Title,
				Show:    tt.fields.Show,
				Actuate: tt.fields.Actuate,
			}
			if err := wo.UnmarshalXML(tt.args.d, tt.args.start); (err != nil) != tt.wantErr {
				t.Errorf("WorkOpus.UnmarshalXML() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(wo, &tt.wantObj); diff != "" {
				t.Errorf("WorkOpus.UnmarshalXML() value is mismatch (-wo +tt.wantObj):%s\n", diff)
			}
		})
	}
}

func TestWorkOpus_SetXMLName(t *testing.T) {
	type fields struct {
		XMLName        xml.Name
		XLink          string
		Href           string
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
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
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{xn: xml.Name{Space: "", Local: "opus"}},
			want: xml.Name{
				Space: "",
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
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
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
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
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xt enum.XlinkTypeEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   enum.XlinkTypeEnum
	}{
		{
			name:   "set",
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{xt: enum.XlinkType.Simple},
			want:   enum.XlinkType.Simple,
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		r string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{r: "role"},
			want:   "role",
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		t string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   string
	}{
		{
			name:   "set",
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{t: "title"},
			want:   "title",
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xs enum.XlinkShowEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   enum.XlinkShowEnum
	}{
		{
			name:   "set",
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{xs: enum.XlinkShow.New},
			want:   enum.XlinkShow.New,
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
		Type           enum.XlinkTypeEnum
		Role           string
		Title          string
		Show           enum.XlinkShowEnum
		Actuate        enum.XlinkActuateEnum
		LinkAttributes attrgrp.LinkAttributes
	}
	type args struct {
		xa enum.XlinkActuateEnum
	}
	tests := []struct {
		name   string
		fields fields
		args   args
		want   enum.XlinkActuateEnum
	}{
		{
			name:   "set",
			fields: fields{XMLName: xml.Name{}, XLink: "", Href: "", Type: "", Role: "", Title: "", Show: "", Actuate: "", LinkAttributes: nil},
			args:   args{xa: enum.XlinkActuate.OnRequest},
			want:   enum.XlinkActuate.OnRequest,
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
					Space: "",
					Local: "opus",
				},
				XLink:             "",
				Href:              "",
				Type:              "",
				Role:              "",
				Title:             "",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
			},
			want: xml.Name{
				Space: "",
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "xlink",
				Href:              "",
				Type:              "",
				Role:              "",
				Title:             "",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "href",
				Type:              "",
				Role:              "",
				Title:             "",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "",
				Type:              "simple",
				Role:              "",
				Title:             "",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
			},
			want: "simple",
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "",
				Type:              "",
				Role:              "role",
				Title:             "",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
			},
			want: "role",
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "",
				Type:              "",
				Role:              "",
				Title:             "title",
				Show:              "",
				Actuate:           "",
				LinkAttributesRaw: nil,
			},
			want: "title",
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "",
				Type:              "",
				Role:              "",
				Title:             "",
				Show:              "new",
				Actuate:           "",
				LinkAttributesRaw: nil,
			},
			want: "new",
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
		Type              string
		Role              string
		Title             string
		Show              string
		Actuate           string
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
				XMLName: xml.Name{
					Space: "",
					Local: "",
				},
				XLink:             "",
				Href:              "",
				Type:              "",
				Role:              "",
				Title:             "",
				Show:              "",
				Actuate:           "onRequest",
				LinkAttributesRaw: nil,
			},
			want: "onRequest",
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
