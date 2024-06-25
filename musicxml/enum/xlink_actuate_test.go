package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToXlinkActuate(t *testing.T) {
	type args struct {
		t string
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
				t: "onRequest"},
			want:    &XlinkActuate.OnRequest,
			wantErr: false,
		},
		{
			name: "onLoad",
			args: args{
				t: "onLoad"},
			want:    &XlinkActuate.OnLoad,
			wantErr: false,
		},
		{
			name: "other",
			args: args{
				t: "other"},
			want:    &XlinkActuate.Other,
			wantErr: false,
		},
		{
			name: "none",
			args: args{
				t: "none"},
			want:    &XlinkActuate.None,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "t"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToXlinkActuate(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXlinkActuate() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToXlinkActuate() value is mismatch (-got, +want):%s\n", diff)
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
			want: []XlinkActuateEnum{
				XlinkActuate.OnRequest,
				XlinkActuate.OnLoad,
				XlinkActuate.Other,
				XlinkActuate.None,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllXlinkActuateEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllXlinkActuateEnumValues() value is mismatch (-got, +want):%s\n", diff)
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
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XlinkActuateEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
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
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XlinkActuateEnum.In() value is mismatch (-got, +want):%s\n", diff)
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
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XlinkActuateEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
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
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XlinkActuateEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestXlinkActuateEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkActuateEnum
		want *string
	}{
		{
			name: "onRequest",
			e:    &XlinkActuate.OnRequest,
			want: testutil.ToStringPtr("onRequest"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("XlinkActuateEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
