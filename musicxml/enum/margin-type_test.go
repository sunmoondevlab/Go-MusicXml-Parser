package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToMarginType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *MarginTypeEnum
		wantErr bool
	}{
		{
			name: "both",
			args: args{
				t: "both",
			},
			want:    &MarginType.Both,
			wantErr: false,
		},
		{
			name: "even",
			args: args{
				t: "even",
			},
			want:    &MarginType.Even,
			wantErr: false,
		},
		{
			name: "odd",
			args: args{
				t: "odd",
			},
			want:    &MarginType.Odd,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "t",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToMarginType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToMarginType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToMarginType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllMarginTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []MarginTypeEnum
	}{
		{
			name: "",
			want: []MarginTypeEnum{
				MarginType.Both,
				MarginType.Even,
				MarginType.Odd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllMarginTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllMarginTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMarginTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj MarginTypeEnum
	}
	tests := []struct {
		name string
		e    *MarginTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &MarginType.Even,
			args: args{
				obj: MarginType.Odd,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &MarginType.Even,
			args: args{
				obj: MarginType.Even,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("MarginTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMarginTypeEnum_In(t *testing.T) {
	type args struct {
		objs []MarginTypeEnum
	}
	tests := []struct {
		name string
		e    *MarginTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &MarginType.Both,
			args: args{
				objs: []MarginTypeEnum{MarginType.Even, MarginType.Odd},
			},
			want: false,
		},
		{
			name: "in",
			e:    &MarginType.Both,
			args: args{
				objs: []MarginTypeEnum{MarginType.Both, MarginType.Even, MarginType.Odd},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("MarginTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMarginTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *MarginTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &MarginType.Both,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("MarginTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMarginTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *MarginTypeEnum
		want string
	}{
		{
			name: "both",
			e:    &MarginType.Both,
			want: "both",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("MarginTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMarginTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *MarginTypeEnum
		want *string
	}{
		{
			name: "both",
			e:    &MarginType.Both,
			want: testutil.ToStringPtr("both"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("MarginTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
