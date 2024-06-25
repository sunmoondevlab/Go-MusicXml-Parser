package valuetext

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestYyyyMmDd_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		ymd     *YyyyMmDd
		args    args
		wantErr bool
		want    YyyyMmDd
	}{
		{
			name: "empty",
			ymd:  (*YyyyMmDd)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("")),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			ymd:  (*YyyyMmDd)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("AD2018-01-01")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			ymd:  (*YyyyMmDd)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("2024-07-01")),
			},
			wantErr: false,
			want:    (YyyyMmDd)("2024-07-01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.ymd.UnmarshalText(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("YyyyMmDd.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.ymd, tt.want); diff != "" {
				t.Errorf("YyyyMmDd.UnmarshalText() value is mismatch (-ymd, +want):%s\n", diff)
			}
		})
	}
}

func TestYyyyMmDd_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		ymd  YyyyMmDd
		want datatype.YyyyMmDd
	}{
		{
			name: "empty",
			ymd:  "",
			want: datatype.YyyyMmDd{},
		},
		{
			name: "invalid",
			ymd:  (YyyyMmDd)("2024-07-011"),
			want: datatype.YyyyMmDd{},
		},
		{
			name: "valid",
			ymd:  (YyyyMmDd)("2024-07-01"),
			want: datatype.YyyyMmDd{
				Val:      testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
				Stringer: nil,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.ymd.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YyyyMmDd.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
