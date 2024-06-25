package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToLeftCenterRight(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *LeftCenterRightEnum
		wantErr bool
	}{
		{
			name: "left",
			args: args{
				t: "left"},
			want:    &LeftCenterRight.Left,
			wantErr: false,
		},
		{
			name: "center",
			args: args{
				t: "center"},
			want:    &LeftCenterRight.Center,
			wantErr: false,
		},
		{
			name: "right",
			args: args{
				t: "right"},
			want:    &LeftCenterRight.Right,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "0"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLeftCenterRight(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLeftCenterRight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToLeftCenterRight() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllLeftCenterRightEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []LeftCenterRightEnum
	}{
		{
			name: "",
			want: []LeftCenterRightEnum{
				LeftCenterRight.Left,
				LeftCenterRight.Center,
				LeftCenterRight.Right,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllLeftCenterRightEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllLeftCenterRightEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLeftCenterRightEnum_Equals(t *testing.T) {
	type args struct {
		obj LeftCenterRightEnum
	}
	tests := []struct {
		name string
		e    *LeftCenterRightEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &LeftCenterRight.Center,
			args: args{
				obj: LeftCenterRight.Right,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &LeftCenterRight.Center,
			args: args{
				obj: LeftCenterRight.Center,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LeftCenterRightEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLeftCenterRightEnum_In(t *testing.T) {
	type args struct {
		objs []LeftCenterRightEnum
	}
	tests := []struct {
		name string
		e    *LeftCenterRightEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &LeftCenterRight.Left,
			args: args{
				objs: []LeftCenterRightEnum{LeftCenterRight.Center, LeftCenterRight.Right},
			},
			want: false,
		},
		{
			name: "in",
			e:    &LeftCenterRight.Left,
			args: args{
				objs: []LeftCenterRightEnum{LeftCenterRight.Left, LeftCenterRight.Center, LeftCenterRight.Right},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LeftCenterRightEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLeftCenterRightEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *LeftCenterRightEnum
		want string
	}{
		{
			name: "ord",
			e:    &LeftCenterRight.Left,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LeftCenterRightEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLeftCenterRightEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *LeftCenterRightEnum
		want string
	}{
		{
			name: "left",
			e:    &LeftCenterRight.Left,
			want: "left",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LeftCenterRightEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLeftCenterRightEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *LeftCenterRightEnum
		want *string
	}{
		{
			name: "left",
			e:    &LeftCenterRight.Left,
			want: testutil.ToStringPtr("left"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LeftCenterRightEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
