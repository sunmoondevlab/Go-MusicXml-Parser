package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToNoteSizeType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *NoteSizeTypeEnum
		wantErr bool
	}{
		{
			name: "cue",
			args: args{
				t: "cue"},
			want:    &NoteSizeType.Cue,
			wantErr: false,
		},
		{
			name: "grace",
			args: args{
				t: "grace"},
			want:    &NoteSizeType.Grace,
			wantErr: false,
		},
		{
			name: "grace-cue",
			args: args{
				t: "grace-cue"},
			want:    &NoteSizeType.GraceCue,
			wantErr: false,
		},
		{
			name: "large",
			args: args{
				t: "large"},
			want:    &NoteSizeType.Large,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "lar"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToNoteSizeType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToNoteSizeType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToNoteSizeType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllNoteSizeTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []NoteSizeTypeEnum
	}{
		{
			name: "",
			want: []NoteSizeTypeEnum{NoteSizeType.Cue,
				NoteSizeType.Grace,
				NoteSizeType.GraceCue,
				NoteSizeType.Large,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllNoteSizeTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllNoteSizeTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj NoteSizeTypeEnum
	}
	tests := []struct {
		name string
		e    *NoteSizeTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &NoteSizeType.Cue,
			args: args{
				obj: NoteSizeType.Grace,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &NoteSizeType.Grace,
			args: args{
				obj: NoteSizeType.Grace,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NoteSizeTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeTypeEnum_In(t *testing.T) {
	type args struct {
		objs []NoteSizeTypeEnum
	}
	tests := []struct {
		name string
		e    *NoteSizeTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &NoteSizeType.Cue,
			args: args{
				objs: []NoteSizeTypeEnum{NoteSizeType.Grace, NoteSizeType.GraceCue},
			},
			want: false,
		},
		{
			name: "in",
			e:    &NoteSizeType.Cue,
			args: args{
				objs: []NoteSizeTypeEnum{NoteSizeType.Cue, NoteSizeType.Grace, NoteSizeType.GraceCue},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NoteSizeTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *NoteSizeTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &NoteSizeType.Cue,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NoteSizeTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *NoteSizeTypeEnum
		want string
	}{
		{
			name: "cue",
			e:    &NoteSizeType.Cue,
			want: "cue",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NoteSizeTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNoteSizeTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *NoteSizeTypeEnum
		want *string
	}{
		{
			name: "cue",
			e:    &NoteSizeType.Cue,
			want: testutil.ToStringPtr("cue"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NoteSizeTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
