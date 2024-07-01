package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToCssFontSize(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *CssFontSizeEnum
		wantErr bool
	}{
		{
			name: "xx-small",
			args: args{
				t: "xx-small"},
			want:    &CssFontSize.XXSmall,
			wantErr: false,
		},
		{
			name: "x-small",
			args: args{
				t: "x-small"},
			want:    &CssFontSize.XSmall,
			wantErr: false,
		},
		{
			name: "small",
			args: args{
				t: "small"},
			want:    &CssFontSize.Small,
			wantErr: false,
		},
		{
			name: "medium",
			args: args{
				t: "medium"},
			want:    &CssFontSize.Medium,
			wantErr: false,
		},
		{
			name: "large",
			args: args{
				t: "large"},
			want:    &CssFontSize.Large,
			wantErr: false,
		},
		{
			name: "x-large",
			args: args{
				t: "x-large"},
			want:    &CssFontSize.XLarge,
			wantErr: false,
		},
		{
			name: "xx-large",
			args: args{
				t: "xx-large"},
			want:    &CssFontSize.XXLarge,
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
			got, err := ToCssFontSize(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCssFontSize() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToCssFontSize() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllCssFontSizeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []CssFontSizeEnum
	}{
		{
			name: "",
			want: []CssFontSizeEnum{CssFontSize.XXSmall, CssFontSize.XSmall, CssFontSize.Small, CssFontSize.Medium, CssFontSize.Large, CssFontSize.XLarge, CssFontSize.XXLarge},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllCssFontSizeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllCssFontSizeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCssFontSizeEnum_Equals(t *testing.T) {
	type args struct {
		obj CssFontSizeEnum
	}
	tests := []struct {
		name string
		e    *CssFontSizeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &CssFontSize.XSmall,
			args: args{
				obj: CssFontSize.Small,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &CssFontSize.XSmall,
			args: args{
				obj: CssFontSize.XSmall,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CssFontSizeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCssFontSizeEnum_In(t *testing.T) {
	type args struct {
		objs []CssFontSizeEnum
	}
	tests := []struct {
		name string
		e    *CssFontSizeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &CssFontSize.XXSmall,
			args: args{
				objs: []CssFontSizeEnum{CssFontSize.XSmall, CssFontSize.Small},
			},
			want: false,
		},
		{
			name: "in",
			e:    &CssFontSize.XXSmall,
			args: args{
				objs: []CssFontSizeEnum{CssFontSize.XXSmall, CssFontSize.XSmall, CssFontSize.Small},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CssFontSizeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCssFontSizeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *CssFontSizeEnum
		want string
	}{
		{
			name: "ord",
			e:    &CssFontSize.XXSmall,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CssFontSizeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCssFontSizeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *CssFontSizeEnum
		want string
	}{
		{
			name: "xx-small",
			e:    &CssFontSize.XXSmall,
			want: "xx-small",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CssFontSizeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCssFontSizeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *CssFontSizeEnum
		want *string
	}{
		{
			name: "xx-small",
			e:    &CssFontSize.XXSmall,
			want: testutil.ToStringPtr("xx-small"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CssFontSizeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
