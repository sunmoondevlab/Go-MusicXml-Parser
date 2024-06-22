package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToXlinkActuate(t *testing.T) {
	type args struct {
		t  string
		en string
	}
	tests := []struct {
		name    string
		args    args
		want    *XlinkActuateEnum
		wantErr bool
	}{
		{
			name: "onRequest",
			args: args{
				t:  "onRequest",
				en: "opus",
			},
			want:    &XlinkActuate.OnRequest,
			wantErr: false,
		},
		{
			name: "onLoad",
			args: args{
				t:  "onLoad",
				en: "opus",
			},
			want:    &XlinkActuate.OnLoad,
			wantErr: false,
		},
		{
			name: "other",
			args: args{
				t:  "other",
				en: "opus",
			},
			want:    &XlinkActuate.Other,
			wantErr: false,
		},
		{
			name: "none",
			args: args{
				t:  "none",
				en: "opus",
			},
			want:    &XlinkActuate.None,
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
			got, err := ToXlinkActuate(tt.args.t, tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXlinkActuate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToXlinkActuate() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllXlinkActuateEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []XlinkActuateEnum
	}{
		{
			name: "",
			want: []XlinkActuateEnum{XlinkActuate.OnRequest, XlinkActuate.OnLoad, XlinkActuate.Other, XlinkActuate.None},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllXlinkActuateEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllXlinkActuateEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestXlinkActuateEnum_Equals(t *testing.T) {
	type args struct {
		obj XlinkActuateEnum
	}
	tests := []struct {
		name string
		e    *XlinkActuateEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &XlinkActuate.OnRequest,
			args: args{
				obj: XlinkActuate.None,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &XlinkActuate.OnRequest,
			args: args{
				obj: XlinkActuate.OnRequest,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equals(tt.args.obj); got != tt.want {
				t.Errorf("XlinkActuateEnum.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkActuateEnum_In(t *testing.T) {
	type args struct {
		objs []XlinkActuateEnum
	}
	tests := []struct {
		name string
		e    *XlinkActuateEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &XlinkActuate.OnRequest,
			args: args{
				objs: []XlinkActuateEnum{XlinkActuate.OnLoad, XlinkActuate.None},
			},
			want: false,
		},
		{
			name: "in",
			e:    &XlinkActuate.OnRequest,
			args: args{
				objs: []XlinkActuateEnum{XlinkActuate.OnRequest, XlinkActuate.OnLoad, XlinkActuate.None},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.In(tt.args.objs...); got != tt.want {
				t.Errorf("XlinkActuateEnum.In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkActuateEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkActuateEnum
		want string
	}{
		{
			name: "ord",
			e:    &XlinkActuate.OnRequest,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Ordinal(); got != tt.want {
				t.Errorf("XlinkActuateEnum.Ordinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkActuateEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkActuateEnum
		want string
	}{
		{
			name: "onRequest",
			e:    &XlinkActuate.OnRequest,
			want: "onRequest",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("XlinkActuateEnum.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
