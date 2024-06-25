package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToCreatorType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *CreatorTypeEnum
		wantErr bool
	}{
		{
			name: "composer",
			args: args{
				t: "composer",
			},
			want:    &CreatorType.Composer,
			wantErr: false,
		},
		{
			name: "arranger",
			args: args{
				t: "arranger",
			},
			want:    &CreatorType.Arranger,
			wantErr: false,
		},
		{
			name: "lyricist",
			args: args{
				t: "lyricist",
			},
			want:    &CreatorType.Lyricist,
			wantErr: false,
		},
		{
			name: "poet",
			args: args{
				t: "poet",
			},
			want:    &CreatorType.Poet,
			wantErr: false,
		},
		{
			name: "translator",
			args: args{
				t: "translator",
			},
			want:    &CreatorType.Translator,
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
			got, err := ToCreatorType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCreatorType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCreatorType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllCreatorTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []CreatorTypeEnum
	}{
		{
			name: "",
			want: []CreatorTypeEnum{CreatorType.Composer, CreatorType.Arranger, CreatorType.Lyricist, CreatorType.Poet, CreatorType.Translator},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllCreatorTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllCreatorTypeEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCreatorTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj CreatorTypeEnum
	}
	tests := []struct {
		name string
		e    *CreatorTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &CreatorType.Composer,
			args: args{
				obj: CreatorType.Translator,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &CreatorType.Composer,
			args: args{
				obj: CreatorType.Composer,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CreatorTypeEnum.Equals() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCreatorTypeEnum_In(t *testing.T) {
	type args struct {
		objs []CreatorTypeEnum
	}
	tests := []struct {
		name string
		e    *CreatorTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &CreatorType.Composer,
			args: args{
				objs: []CreatorTypeEnum{CreatorType.Arranger, CreatorType.Lyricist},
			},
			want: false,
		},
		{
			name: "in",
			e:    &CreatorType.Composer,
			args: args{
				objs: []CreatorTypeEnum{CreatorType.Composer, CreatorType.Arranger, CreatorType.Lyricist},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CreatorTypeEnum.In() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCreatorTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *CreatorTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &CreatorType.Composer,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CreatorTypeEnum.Ordinal() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCreatorTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *CreatorTypeEnum
		want string
	}{
		{
			name: "composer",
			e:    &CreatorType.Composer,
			want: "composer",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CreatorTypeEnum.String() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCreatorTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *CreatorTypeEnum
		want *string
	}{
		{
			name: "composer",
			e:    &CreatorType.Composer,
			want: testutil.ToStringPtr("composer"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CreatorTypeEnum.StringPtr() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}
