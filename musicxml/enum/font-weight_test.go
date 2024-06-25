package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNewFontWeight(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *FontWeightEnum
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				t: "normal"},
			want:    &FontWeight.Normal,
			wantErr: false,
		},
		{
			name: "bold",
			args: args{
				t: "bold"},
			want:    &FontWeight.Bold,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "la"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewFontWeight(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFontWeight() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewFontWeight() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllFontWeightEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []FontWeightEnum
	}{
		{
			name: "",
			want: []FontWeightEnum{
				FontWeight.Normal,
				FontWeight.Bold,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllFontWeightEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllFontWeightEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontWeightEnum_Equals(t *testing.T) {
	type args struct {
		obj FontWeightEnum
	}
	tests := []struct {
		name string
		e    *FontWeightEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &FontWeight.Bold,
			args: args{
				obj: FontWeight.Normal,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &FontWeight.Bold,
			args: args{
				obj: FontWeight.Bold,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontWeightEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontWeightEnum_In(t *testing.T) {
	type args struct {
		objs []FontWeightEnum
	}
	tests := []struct {
		name string
		e    *FontWeightEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &FontWeight.Normal,
			args: args{
				objs: []FontWeightEnum{FontWeight.Bold},
			},
			want: false,
		},
		{
			name: "in",
			e:    &FontWeight.Normal,
			args: args{
				objs: []FontWeightEnum{FontWeight.Normal, FontWeight.Bold},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontWeightEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontWeightEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *FontWeightEnum
		want string
	}{
		{
			name: "ord",
			e:    &FontWeight.Normal,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontWeightEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontWeightEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *FontWeightEnum
		want string
	}{
		{
			name: "normal",
			e:    &FontWeight.Normal,
			want: "normal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontWeightEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontWeightEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *FontWeightEnum
		want *string
	}{
		{
			name: "normal",
			e:    &FontWeight.Normal,
			want: testutil.ToStringPtr("normal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontWeightEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
