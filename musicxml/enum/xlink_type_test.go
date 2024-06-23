package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
)

func TestToXlinkType(t *testing.T) {
	type args struct {
		t  string
		en string
	}
	tests := []struct {
		name    string
		args    args
		want    *XlinkTypeEnum
		wantErr bool
	}{
		{
			name: "simple",
			args: args{
				t:  "simple",
				en: "opus",
			},
			want:    &XlinkType.Simple,
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
			got, err := ToXlinkType(tt.args.t, tt.args.en)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToXlinkType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToXlinkType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllXlinkTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []XlinkTypeEnum
	}{
		{
			name: "",
			want: []XlinkTypeEnum{XlinkType.Simple},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllXlinkTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllXlinkTypeEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestXlinkTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj XlinkTypeEnum
	}
	tests := []struct {
		name string
		e    *XlinkTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &XlinkType.Simple,
			args: args{
				obj: "",
			},
			want: false,
		},
		{
			name: "eq",
			e:    &XlinkType.Simple,
			args: args{
				obj: XlinkType.Simple,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Equals(tt.args.obj); got != tt.want {
				t.Errorf("XlinkTypeEnum.Equals() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkTypeEnum_In(t *testing.T) {
	type args struct {
		objs []XlinkTypeEnum
	}
	tests := []struct {
		name string
		e    *XlinkTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &XlinkType.Simple,
			args: args{
				objs: []XlinkTypeEnum{},
			},
			want: false,
		},
		{
			name: "in",
			e:    &XlinkType.Simple,
			args: args{
				objs: []XlinkTypeEnum{XlinkType.Simple},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.In(tt.args.objs...); got != tt.want {
				t.Errorf("XlinkTypeEnum.In() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &XlinkType.Simple,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Ordinal(); got != tt.want {
				t.Errorf("XlinkTypeEnum.Ordinal() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestXlinkTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *XlinkTypeEnum
		want string
	}{
		{
			name: "simple",
			e:    &XlinkType.Simple,
			want: "simple",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.String(); got != tt.want {
				t.Errorf("XlinkTypeEnum.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
