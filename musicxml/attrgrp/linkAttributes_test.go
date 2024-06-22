package attrgrp_test

import (
	"bytes"
	"encoding/xml"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/attrgrp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
)

func TestUnmarshalLinkAttributes(t *testing.T) {
	type args struct {
		d     *xml.Decoder
		start xml.StartElement
		l     attrgrp.LinkAttributes
		lr    attrgrp.LinkAttributesRaw
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		wantObj attrgrp.LinkAttributes
	}{
		{
			name: "work>opus invalid decode",
			args: args{
				d: xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simpl" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{
					Name: xml.Name{
						Space: "",
						Local: "opp",
					},
					Attr: []xml.Attr{},
				},
				l:  &musicxml.WorkOpus{},
				lr: &musicxml.WorkOpusRaw{},
			},
			wantErr: true,
			wantObj: &musicxml.WorkOpus{},
		},
		{
			name: "work>opus invalid type",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simpl" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
				l:     &musicxml.WorkOpus{},
				lr:    &musicxml.WorkOpusRaw{},
			},
			wantErr: true,
			wantObj: &musicxml.WorkOpus{},
		},
		{
			name: "work>opus invalid show",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="ne" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
				l:     &musicxml.WorkOpus{},
				lr:    &musicxml.WorkOpusRaw{},
			},
			wantErr: true,
			wantObj: &musicxml.WorkOpus{},
		},
		{
			name: "work>opus invalid actuate",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlinkactuate:="non"/>`))),
				start: xml.StartElement{},
				l:     &musicxml.WorkOpus{},
				lr:    &musicxml.WorkOpusRaw{},
			},
			wantErr: true,
			wantObj: &musicxml.WorkOpus{},
		},
		{
			name: "work>opus",
			args: args{
				d:     xml.NewDecoder(bytes.NewReader([]byte(`<opus xmlns:xlink="http://www.w3.org/1999/xlink" xlink:href="opus/winterreise.musicxml" xlink:type="simple" xlink:role="role" xlink:title="winterreise" xlink:show="new" xlink:actuate="none"/>`))),
				start: xml.StartElement{},
				l:     &musicxml.WorkOpus{},
				lr:    &musicxml.WorkOpusRaw{},
			},
			wantErr: false,
			wantObj: &musicxml.WorkOpus{
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
			if err := attrgrp.UnmarshalLinkAttributes(tt.args.d, tt.args.start, tt.args.l, tt.args.lr); (err != nil) != tt.wantErr {
				t.Errorf("UnmarshalLinkAttributes() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(tt.args.l, tt.wantObj); diff != "" {
				t.Errorf("WorkOpus.UnmarshalLinkAttributes() value is mismatch (-tt.args.l +tt.wantObj):%s\n", diff)
			}
		})
	}
}
