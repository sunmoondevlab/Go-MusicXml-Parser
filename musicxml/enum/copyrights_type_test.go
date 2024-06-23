package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
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
			name: "",
			args: args{
				t: "",
			},
			want:    &CopyRightsType.All,
			wantErr: false,
		},
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
				CopyRightsType.All,
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
			if got := tt.e.Equals(tt.args.obj); got != tt.want {
				t.Errorf("CopyRightsTypeEnum.Equals() = %v, want %v", got, tt.want)
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
			if got := tt.e.In(tt.args.objs...); got != tt.want {
				t.Errorf("CopyRightsTypeEnum.In() = %v, want %v", got, tt.want)
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
			e:    &CopyRightsType.All,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.e.Ordinal(); got != tt.want {
				t.Errorf("CopyRightsTypeEnum.Ordinal() = %v, want %v", got, tt.want)
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
			if got := tt.e.String(); got != tt.want {
				t.Errorf("CopyRightsTypeEnum.String() = %v, want %v", got, tt.want)
			}
		})
	}
}
