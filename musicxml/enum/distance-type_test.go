package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToDistanceType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *DistanceTypeEnum
		wantErr bool
	}{
		{
			name: "beam",
			args: args{
				t: "beam"},
			want:    &DistanceType.Beam,
			wantErr: false,
		},
		{
			name: "hyphen",
			args: args{
				t: "hyphen"},
			want:    &DistanceType.Hyphen,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "bea"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToDistanceType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToDistanceType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToDistanceType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllDistanceTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []DistanceTypeEnum
	}{
		{
			name: "",
			want: []DistanceTypeEnum{DistanceType.Beam,
				DistanceType.Hyphen,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllDistanceTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllDistanceTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestDistanceTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj DistanceTypeEnum
	}
	tests := []struct {
		name string
		e    *DistanceTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &DistanceType.Beam,
			args: args{
				obj: DistanceType.Hyphen,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &DistanceType.Beam,
			args: args{
				obj: DistanceType.Beam,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DistanceTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestDistanceTypeEnum_In(t *testing.T) {
	type args struct {
		objs []DistanceTypeEnum
	}
	tests := []struct {
		name string
		e    *DistanceTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &DistanceType.Beam,
			args: args{
				objs: []DistanceTypeEnum{DistanceType.Hyphen},
			},
			want: false,
		},
		{
			name: "in",
			e:    &DistanceType.Beam,
			args: args{
				objs: []DistanceTypeEnum{DistanceType.Beam, DistanceType.Hyphen},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DistanceTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestDistanceTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *DistanceTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &DistanceType.Beam,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DistanceTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestDistanceTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *DistanceTypeEnum
		want string
	}{
		{
			name: "hyphen",
			e:    &DistanceType.Hyphen,
			want: "hyphen",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DistanceTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestDistanceTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *DistanceTypeEnum
		want *string
	}{
		{
			name: "hyphen",
			e:    &DistanceType.Hyphen,
			want: testutil.ToStringPtr("hyphen"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("DistanceTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
