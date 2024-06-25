package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestStaffNumber_String(t *testing.T) {
	type fields struct {
		Value    *uint64
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
				Value: testutil.ToUint64Ptr(11),
			},
			want: "11",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := &StaffNumber{
				Val:      tt.fields.Value,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				sn = nil
			}
			got := sn.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("StaffNumber.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestStaffNumber_StringPtr(t *testing.T) {
	type fields struct {
		Uint     *uint64
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
				Uint: testutil.ToUint64Ptr(11),
			},
			want: testutil.ToStringPtr("11"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := &StaffNumber{
				Val:      tt.fields.Uint,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				sn = nil
			}
			got := sn.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("StaffNumber.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestStaffNumber_Value(t *testing.T) {
	type fields struct {
		Uint     *uint64
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *uint64
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
				Uint: testutil.ToUint64Ptr(11),
			},
			want: testutil.ToUint64Ptr(11),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sn := &StaffNumber{
				Val:      tt.fields.Uint,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				sn = nil
			}
			got := sn.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("StaffNumber.Value() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewStaffNumber(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *StaffNumber
		wantErr bool
	}{
		{
			name:    "NewStaffNumber empty",
			args:    args{},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewStaffNumber invalid",
			args: args{
				value: "-1",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewStaffNumber valid",
			args: args{
				value: "111",
			},
			want: &StaffNumber{
				Val: testutil.ToUint64Ptr(111),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewStaffNumber(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewStaffNumber() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewStaffNumber() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
