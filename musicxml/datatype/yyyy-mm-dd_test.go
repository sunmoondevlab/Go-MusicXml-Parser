package datatype

import (
	"fmt"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestYyyyMmDd_String(t *testing.T) {
	type fields struct {
		Val      *time.Time
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		isNil  bool
	}{
		{
			name:   "string",
			fields: fields{},
			want:   "",
			isNil:  true,
		},
		{
			name:   "string",
			fields: fields{},
			want:   "",
		},
		{
			name: "string",
			fields: fields{
				Val: testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
			},
			want: "2024-07-01",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymd := &YyyyMmDd{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				ymd = nil
			}
			got := ymd.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YyyyMmDd.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestYyyyMmDd_StringPtr(t *testing.T) {
	type fields struct {
		Val      *time.Time
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		isNil  bool
	}{
		{
			name:   "string",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name:   "string",
			fields: fields{},
			want:   nil,
		},
		{
			name: "string",
			fields: fields{
				Val: testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
			},
			want: testutil.ToStringPtr("2024-07-01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymd := &YyyyMmDd{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				ymd = nil
			}
			got := ymd.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YyyyMmDd.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestYyyyMmDd_Value(t *testing.T) {
	type fields struct {
		Val      *time.Time
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *time.Time
		isNil  bool
	}{
		{
			name:   "value",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name:   "value",
			fields: fields{},
			want:   nil,
		},
		{
			name: "value",
			fields: fields{
				Val: testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
			},
			want: testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ymd := &YyyyMmDd{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				ymd = nil
			}
			got := ymd.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("YyyyMmDd.Value() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewYyyyMmDd(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *YyyyMmDd
		wantErr bool
	}{
		{
			name:    "NewYyyyMmDd empty",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewYyyyMmDd invalid",
			args: args{
				value: "A2001-2-11",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewYyyyMmDd valid",
			args: args{
				value: "2024-07-01",
			},
			want: &YyyyMmDd{
				Val: testutil.ToTimePtr(time.Date(2024, 7, 1, 0, 0, 0, 0, time.UTC)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewYyyyMmDd(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewYyyyMmDd() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewYyyyMmDd() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
