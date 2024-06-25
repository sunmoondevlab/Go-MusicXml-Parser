package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/enum"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestFontSize_String(t *testing.T) {
	type fields struct {
		FontPointSize *decimal.Decimal
		CssFontSize   *enum.CssFontSizeEnum
		Stringer      fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		isNil  bool
	}{
		{
			name:   "nil",
			fields: fields{},
			want:   "",
			isNil:  true,
		},
		{
			name:   "empty",
			fields: fields{},
			want:   "",
		},
		{
			name: "css font size",
			fields: fields{
				CssFontSize: &enum.CssFontSize.Medium,
			},
			want: "medium",
		},
		{
			name: "font point size",
			fields: fields{
				FontPointSize: testutil.ToDecimalPtr("43.44"),
			},
			want: "43.44",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FontSize{
				FontPointSize: tt.fields.FontPointSize,
				CssFontSize:   tt.fields.CssFontSize,
				Stringer:      tt.fields.Stringer,
			}
			if tt.isNil {
				fs = nil
			}
			got := fs.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontSize.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontSize_StringPtr(t *testing.T) {
	type fields struct {
		FontPointSize *decimal.Decimal
		CssFontSize   *enum.CssFontSizeEnum
		Stringer      fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		isNil  bool
	}{
		{
			name:   "nil",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name:   "empty",
			fields: fields{},
			want:   nil,
		},
		{
			name: "css font size",
			fields: fields{
				CssFontSize: &enum.CssFontSize.Medium,
			},
			want: testutil.ToStringPtr("medium"),
		},
		{
			name: "font point size",
			fields: fields{
				FontPointSize: testutil.ToDecimalPtr("43.44"),
			},
			want: testutil.ToStringPtr("43.44"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := &FontSize{
				FontPointSize: tt.fields.FontPointSize,
				CssFontSize:   tt.fields.CssFontSize,
				Stringer:      tt.fields.Stringer,
			}
			if tt.isNil {
				fs = nil
			}
			got := fs.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontSize.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewFontSize(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *FontSize
		wantErr bool
	}{
		{
			name:    "nil",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "css-font-size xx-small",
			args: args{
				t: "xx-small"},
			want: &FontSize{
				CssFontSize: &enum.CssFontSize.XXSmall,
			},
			wantErr: false,
		},
		{
			name: "decimal",
			args: args{
				t: "40.4"},
			want: &FontSize{
				FontPointSize: testutil.ToDecimalPtr("40.4"),
			},
			wantErr: false,
		},
		{
			name: "css-font-size invalid",
			args: args{
				t: "xx-smalld"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFontSize(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFontSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewFontSize() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
