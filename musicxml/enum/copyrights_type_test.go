package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToCopyrightsType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *CopyrightsTypeEnum
		wantErr bool
	}{
		{
			name: "music",
			args: args{
				t: "music",
			},
			want:    &CopyrightsType.Music,
			wantErr: false,
		},
		{
			name: "words",
			args: args{
				t: "words",
			},
			want:    &CopyrightsType.Words,
			wantErr: false,
		},

		{
			name: "arrangement",
			args: args{
				t: "arrangement",
			},
			want:    &CopyrightsType.Arrangement,
			wantErr: false,
		},
		{
			name: "translation",
			args: args{
				t: "translation",
			},
			want:    &CopyrightsType.Translation,
			wantErr: false,
		},
		{
			name: "parody",
			args: args{
				t: "parody",
			},
			want:    &CopyrightsType.Parody,
			wantErr: false,
		},
		{
			name: "other",
			args: args{
				t: "song",
			},
			want:    &CopyrightsType.Other,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToCopyrightsType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToCopyrightsType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToCopyrightsType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllCopyrightsTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []CopyrightsTypeEnum
	}{
		{
			name: "",
			want: []CopyrightsTypeEnum{
				CopyrightsType.Music,
				CopyrightsType.Arrangement,
				CopyrightsType.Words,
				CopyrightsType.Translation,
				CopyrightsType.Parody,
				CopyrightsType.Other,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllCopyrightsTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllCopyrightsTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCopyrightsTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj CopyrightsTypeEnum
	}
	tests := []struct {
		name string
		e    *CopyrightsTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &CopyrightsType.Music,
			args: args{
				obj: CopyrightsType.Words,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &CopyrightsType.Music,
			args: args{
				obj: CopyrightsType.Music,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyrightsTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCopyrightsTypeEnum_In(t *testing.T) {
	type args struct {
		objs []CopyrightsTypeEnum
	}
	tests := []struct {
		name string
		e    *CopyrightsTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &CopyrightsType.Music,
			args: args{
				objs: []CopyrightsTypeEnum{CopyrightsType.Arrangement, CopyrightsType.Translation},
			},
			want: false,
		},
		{
			name: "in",
			e:    &CopyrightsType.Music,
			args: args{
				objs: []CopyrightsTypeEnum{CopyrightsType.Music, CopyrightsType.Arrangement, CopyrightsType.Translation},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyrightsTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCopyrightsTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyrightsTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &CopyrightsType.Music,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyrightsTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCopyrightsTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyrightsTypeEnum
		want string
	}{
		{
			name: "music",
			e:    &CopyrightsType.Music,
			want: "music",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyrightsTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestCopyrightsTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *CopyrightsTypeEnum
		want *string
	}{
		{
			name: "music",
			e:    &CopyrightsType.Music,
			want: testutil.ToStringPtr("music"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("CopyrightsTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
