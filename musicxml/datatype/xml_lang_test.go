package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestXmlLang_String(t *testing.T) {
	type fields struct {
		Val      *string
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		isNil  bool
	}{
		{
			name:   "string",
			fields: fields{},
			want:   "",
			isNil:  true,
		},
		{
			name:   "string",
			fields: fields{},
			want:   "",
		},
		{
			name: "string",
			fields: fields{
				Val: testutil.ToStringPtr("ja-JP"),
			},
			want: "ja-JP",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xl := &XmlLang{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				xl = nil
			}
			got := xl.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XmlLang.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestXmlLang_StringPtr(t *testing.T) {
	type fields struct {
		Val      *string
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		isNil  bool
	}{
		{
			name:   "string",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name:   "string",
			fields: fields{},
			want:   nil,
		},
		{
			name: "string",
			fields: fields{
				Val: testutil.ToStringPtr("ja-JP"),
			},
			want: testutil.ToStringPtr("ja-JP"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			xl := &XmlLang{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				xl = nil
			}
			got := xl.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XmlLang.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewXmlLang(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *XmlLang
		wantErr bool
	}{
		{
			name: "NewXmlLang empty",
			args: args{
				value: " ",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewXmlLang invalid",
			args: args{
				value: "aaaaa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewXmlLang",
			args: args{
				value: "ja-JP",
			},
			want: &XmlLang{
				Val: testutil.ToStringPtr("ja-JP"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewXmlLang(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewXmlLang() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewXmlLang() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
