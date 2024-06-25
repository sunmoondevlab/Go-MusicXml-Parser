package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToCopyRightsType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *CopyRightsTypeEnum
		wantErr bool
	}{
		{
			name: "music",
			args: args{
				t: "music",
			},
			want:    &CopyRightsType.Music,
			wantErr: false,
		},
		{
			name: "words",
			args: args{
				t: "words",
			},
			want:    &CopyRightsType.Words,
			wantErr: false,
		},

		{
			name: "arrangement",
			args: args{
				t: "arrangement",
			},
			want:    &CopyRightsType.Arrangement,
			wantErr: false,
		},
		{
			name: "translation",
			args: args{
				t: "translation",
			},
			want:    &CopyRightsType.Translation,
			wantErr: false,
		},
		{
			name: "parody",
			args: args{
				t: "parody",
			},
			want:    &CopyRightsType.Parody,
			wantErr: false,
		},
		{
			name: "other",
			args: args{
				t: "song",
			},
			want:    &CopyRightsType.Other,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCopyRightsType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCopyRightsType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("ToCopyRightsType() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAllCopyRightsTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []CopyRightsTypeEnum
	}{
		{
			name: "",
			want: []CopyRightsTypeEnum{
				CopyRightsType.Music,
				CopyRightsType.Arrangement,
				CopyRightsType.Words,
				CopyRightsType.Translation,
				CopyRightsType.Parody,
				CopyRightsType.Other,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllCopyRightsTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllCopyRightsTypeEnumValues() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCopyRightsTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj CopyRightsTypeEnum
	}
	tests := []struct {
		name string
		e    *CopyRightsTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &CopyRightsType.Music,
			args: args{
				obj: CopyRightsType.Words,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &CopyRightsType.Music,
			args: args{
				obj: CopyRightsType.Music,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyRightsTypeEnum.Equals() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCopyRightsTypeEnum_In(t *testing.T) {
	type args struct {
		objs []CopyRightsTypeEnum
	}
	tests := []struct {
		name string
		e    *CopyRightsTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &CopyRightsType.Music,
			args: args{
				objs: []CopyRightsTypeEnum{CopyRightsType.Arrangement, CopyRightsType.Translation},
			},
			want: false,
		},
		{
			name: "in",
			e:    &CopyRightsType.Music,
			args: args{
				objs: []CopyRightsTypeEnum{CopyRightsType.Music, CopyRightsType.Arrangement, CopyRightsType.Translation},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyRightsTypeEnum.In() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCopyRightsTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyRightsTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &CopyRightsType.Music,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyRightsTypeEnum.Ordinal() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCopyRightsTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyRightsTypeEnum
		want string
	}{
		{
			name: "music",
			e:    &CopyRightsType.Music,
			want: "music",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyRightsTypeEnum.String() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}

func TestCopyRightsTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyRightsTypeEnum
		want *string
	}{
		{
			name: "music",
			e:    &CopyRightsType.Music,
			want: testutil.ToStringPtr("music"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyRightsTypeEnum.StringPtr() value is mismatch (-got +tt.want):%s\n", diff)
			}
		})
	}
}
