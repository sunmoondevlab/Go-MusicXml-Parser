package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToXlinkShow(t *testing.T) {
	type args struct {
		t  string
		en string
	}
	tests := []struct {
		name    string
		args    args
		want    *XlinkShowEnum
		wantErr bool
	}{
		{
			name: "new",
			args: args{
				t:  "new",
				en: "opus",
			},
			want:    &XlinkShow.New,
			wantErr: false,
		},
		{
			name: "replace",
			args: args{
				t:  "replace",
				en: "opus",
			},
			want:    &XlinkShow.Replace,
			wantErr: false,
		},
		{
			name: "embed",
			args: args{
				t:  "embed",
				en: "opus",
			},
			want:    &XlinkShow.Embed,
			wantErr: false,
		},
		{
			name: "other",
			args: args{
				t:  "other",
				en: "opus",
			},
			want:    &XlinkShow.Other,
			wantErr: false,
		},
		{
			name: "none",
			args: args{
				t:  "none",
				en: "opus",
			},
			want:    &XlinkShow.None,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t:  "t",
				en: "opus",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToXlinkShow(tt.args.t, tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXlinkShow() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToXlinkShow() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllXlinkShowEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []XlinkShowEnum
	}{
		{
			name: "",
			want: []XlinkShowEnum{XlinkShow.New, XlinkShow.Replace, XlinkShow.Embed, XlinkShow.Other, XlinkShow.None},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllXlinkShowEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllXlinkShowEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestXlinkShowEnum_Equals(t *testing.T) {
	type args struct {
		obj XlinkShowEnum
	}
	tests := []struct {
		name string
		e    *XlinkShowEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &XlinkShow.New,
			args: args{
				obj: XlinkShow.None,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &XlinkShow.New,
			args: args{
				obj: XlinkShow.New,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equals(tt.args.obj); got != tt.want {
				t.Errorf("XlinkShowEnum.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkShowEnum_In(t *testing.T) {
	type args struct {
		objs []XlinkShowEnum
	}
	tests := []struct {
		name string
		e    *XlinkShowEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &XlinkShow.New,
			args: args{
				objs: []XlinkShowEnum{XlinkShow.Replace, XlinkShow.Embed},
			},
			want: false,
		},
		{
			name: "in",
			e:    &XlinkShow.New,
			args: args{
				objs: []XlinkShowEnum{XlinkShow.New, XlinkShow.Replace, XlinkShow.Embed},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.In(tt.args.objs...); got != tt.want {
				t.Errorf("XlinkShowEnum.In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkShowEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkShowEnum
		want string
	}{
		{
			name: "ord",
			e:    &XlinkShow.New,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Ordinal(); got != tt.want {
				t.Errorf("XlinkShowEnum.Ordinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkShowEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkShowEnum
		want string
	}{
		{
			name: "new",
			e:    &XlinkShow.New,
			want: "new",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("XlinkShowEnum.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
