package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestTenths_String(t *testing.T) {
	type fields struct {
		Val      *decimal.Decimal
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
				Val: testutil.DecimalPtr(decimal.New(1, 1)),
			},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tenths{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tenths.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestTenths_StringPtr(t *testing.T) {
	type fields struct {
		Val      *decimal.Decimal
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
				Val: testutil.DecimalPtr(decimal.New(1, 1)),
			},
			want: testutil.ToStringPtr("10"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tenths{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tenths.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestTenths_Value(t *testing.T) {
	type fields struct {
		Val      *decimal.Decimal
		Stringer fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *decimal.Decimal
		isNil  bool
	}{
		{
			name:   "value",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name: "value",
			fields: fields{
				Val:      testutil.DecimalPtr(decimal.New(2, 3)),
				Stringer: nil,
			},
			want: testutil.DecimalPtr(decimal.New(2, 3)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &Tenths{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tenths.DecimalValue() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewTenthsFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *Tenths
		wantErr bool
	}{
		{
			name: "NewTenthsFromString empty",
			args: args{
				value: " ",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewTenthsFromString invalid",
			args: args{
				value: "111,48",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewTenthsFromString",
			args: args{
				value: "111.48",
			},
			want: &Tenths{
				Val: testutil.DecimalPtr(decimal.New(11148, -2)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewTenthsFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewTenthsFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewTenthsFromString() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewTenthsFixedPoint(t *testing.T) {
	type args struct {
		value int64
		exp   int32
	}
	tests := []struct {
		name string
		args args
		want *Tenths
	}{
		{
			name: "NewTenthsFixedPoint",
			args: args{
				value: 1287,
				exp:   -3,
			},
			want: &Tenths{
				Val: testutil.DecimalPtr(decimal.New(1287, -3)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTenthsFixedPoint(tt.args.value, tt.args.exp)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewTenthsFixedPoint() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewTenthsFromInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want *Tenths
	}{
		{
			name: "NewTenthsFromInt",
			args: args{
				value: 184,
			},
			want: &Tenths{
				Val: testutil.DecimalPtr(decimal.NewFromInt(184)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTenthsFromInt(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewTenthsFromInt() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewTenthsFromUint64(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want *Tenths
	}{
		{
			name: "NewTenthsFromUint64",
			args: args{
				value: 11113,
			},
			want: &Tenths{
				Val: testutil.DecimalPtr(decimal.NewFromUint64(11113)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTenthsFromUint64(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewTenthsFromUint64() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewTenthsFromFloat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *Tenths
	}{
		{
			name: "NewTenthsFromFloat",
			args: args{
				value: 12.18,
			},
			want: &Tenths{
				Val: testutil.DecimalPtr(decimal.NewFromFloat(12.18)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewTenthsFromFloat(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewTenthsFromFloat() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
