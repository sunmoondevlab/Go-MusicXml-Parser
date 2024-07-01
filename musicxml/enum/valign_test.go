package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToValign(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *ValignEnum
		wantErr bool
	}{
		{
			name: "top",
			args: args{
				t: "top"},
			want:    &Valign.Top,
			wantErr: false,
		},
		{
			name: "middle",
			args: args{
				t: "middle"},
			want:    &Valign.Middle,
			wantErr: false,
		},
		{
			name: "bottom",
			args: args{
				t: "bottom"},
			want:    &Valign.Bottom,
			wantErr: false,
		},
		{
			name: "baseline",
			args: args{
				t: "baseline"},
			want:    &Valign.Baseline,
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
			got, err := ToValign(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToValign() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToValign() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllValignEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []ValignEnum
	}{
		{
			name: "",
			want: []ValignEnum{
				Valign.Top,
				Valign.Middle,
				Valign.Bottom,
				Valign.Baseline,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllValignEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllValignEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValignEnum_Equals(t *testing.T) {
	type args struct {
		obj ValignEnum
	}
	tests := []struct {
		name string
		e    *ValignEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &Valign.Middle,
			args: args{
				obj: Valign.Bottom,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &Valign.Middle,
			args: args{
				obj: Valign.Middle,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ValignEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValignEnum_In(t *testing.T) {
	type args struct {
		objs []ValignEnum
	}
	tests := []struct {
		name string
		e    *ValignEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &Valign.Top,
			args: args{
				objs: []ValignEnum{Valign.Middle, Valign.Bottom},
			},
			want: false,
		},
		{
			name: "in",
			e:    &Valign.Top,
			args: args{
				objs: []ValignEnum{Valign.Top, Valign.Middle, Valign.Bottom},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ValignEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValignEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *ValignEnum
		want string
	}{
		{
			name: "ord",
			e:    &Valign.Top,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ValignEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValignEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *ValignEnum
		want string
	}{
		{
			name: "top",
			e:    &Valign.Top,
			want: "top",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ValignEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValignEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *ValignEnum
		want *string
	}{
		{
			name: "top",
			e:    &Valign.Top,
			want: testutil.ToStringPtr("top"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ValignEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
