package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToYesNo(t *testing.T) {
	type args struct {
		t  string
		en string
		at string
	}
	tests := []struct {
		name    string
		args    args
		want    *YesNoEnum
		wantErr bool
	}{
		{
			name: "yes",
			args: args{
				t:  "yes",
				en: "support",
				at: "type",
			},
			want:    &YesNo.Yes,
			wantErr: false,
		},
		{
			name: "no",
			args: args{
				t:  "no",
				en: "support",
				at: "type",
			},
			want:    &YesNo.No,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t:  "t",
				en: "support",
				at: "type",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToYesNo(tt.args.t, tt.args.en, tt.args.at)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToYesNo() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToYesNo() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestAllYesNoEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []YesNoEnum
	}{
		{
			name: "",
			want: []YesNoEnum{YesNo.Yes, YesNo.No},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllYesNoEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllYesNoEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestYesNoEnum_Equals(t *testing.T) {
	type args struct {
		obj YesNoEnum
	}
	tests := []struct {
		name string
		e    *YesNoEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &YesNo.Yes,
			args: args{
				obj: YesNo.No,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &YesNo.Yes,
			args: args{
				obj: YesNo.Yes,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YesNoEnum.Equals() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestYesNoEnum_In(t *testing.T) {
	type args struct {
		objs []YesNoEnum
	}
	tests := []struct {
		name string
		e    *YesNoEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &YesNo.Yes,
			args: args{
				objs: []YesNoEnum{YesNo.No},
			},
			want: false,
		},
		{
			name: "in",
			e:    &YesNo.Yes,
			args: args{
				objs: []YesNoEnum{YesNo.Yes, YesNo.No},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YesNoEnum.In() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestYesNoEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *YesNoEnum
		want string
	}{
		{
			name: "ord",
			e:    &YesNo.Yes,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YesNoEnum.Ordinal() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestYesNoEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *YesNoEnum
		want string
	}{
		{
			name: "yes",
			e:    &YesNo.Yes,
			want: "yes",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YesNoEnum.String() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestYesNoEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *YesNoEnum
		want *string
	}{
		{
			name: "yes",
			e:    &YesNo.Yes,
			want: testutil.ToStringPtr("yes"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YesNoEnum.StringPtr() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}
