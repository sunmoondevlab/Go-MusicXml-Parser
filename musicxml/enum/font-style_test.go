package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNewFontStyle(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *FontStyleEnum
		wantErr bool
	}{
		{
			name: "normal",
			args: args{
				t: "normal"},
			want:    &FontStyle.Normal,
			wantErr: false,
		},
		{
			name: "italic",
			args: args{
				t: "italic"},
			want:    &FontStyle.Italic,
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
			got, err := NewFontStyle(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewFontStyle() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewFontStyle() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllFontStyleEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []FontStyleEnum
	}{
		{
			name: "",
			want: []FontStyleEnum{
				FontStyle.Normal,
				FontStyle.Italic,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllFontStyleEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllFontStyleEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontStyleEnum_Equals(t *testing.T) {
	type args struct {
		obj FontStyleEnum
	}
	tests := []struct {
		name string
		e    *FontStyleEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &FontStyle.Italic,
			args: args{
				obj: FontStyle.Normal,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &FontStyle.Italic,
			args: args{
				obj: FontStyle.Italic,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontStyleEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontStyleEnum_In(t *testing.T) {
	type args struct {
		objs []FontStyleEnum
	}
	tests := []struct {
		name string
		e    *FontStyleEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &FontStyle.Normal,
			args: args{
				objs: []FontStyleEnum{FontStyle.Italic},
			},
			want: false,
		},
		{
			name: "in",
			e:    &FontStyle.Normal,
			args: args{
				objs: []FontStyleEnum{FontStyle.Normal, FontStyle.Italic},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontStyleEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontStyleEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *FontStyleEnum
		want string
	}{
		{
			name: "ord",
			e:    &FontStyle.Normal,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontStyleEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontStyleEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *FontStyleEnum
		want string
	}{
		{
			name: "normal",
			e:    &FontStyle.Normal,
			want: "normal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontStyleEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestFontStyleEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *FontStyleEnum
		want *string
	}{
		{
			name: "normal",
			e:    &FontStyle.Normal,
			want: testutil.ToStringPtr("normal"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("FontStyleEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
