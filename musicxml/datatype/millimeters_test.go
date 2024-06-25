package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestMillimeters_String(t *testing.T) {
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
			m := &Millimeters{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				m = nil
			}
			got := m.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Millimeters.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMillimeters_StringPtr(t *testing.T) {
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
			m := &Millimeters{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				m = nil
			}
			got := m.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Millimeters.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestMillimeters_Value(t *testing.T) {
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
			m := &Millimeters{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				m = nil
			}
			got := m.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Millimeters.Value() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewMillimetersFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *Millimeters
		wantErr bool
	}{
		{
			name: "NewMillimetersFromString empty",
			args: args{
				value: " ",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewMillimetersFromString invalid",
			args: args{
				value: "111,48",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewMillimetersFromString",
			args: args{
				value: "111.48",
			},
			want: &Millimeters{
				Val: testutil.DecimalPtr(decimal.New(11148, -2)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewMillimetersFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewMillimetersFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMillimetersFromString() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewMillimetersFixedPoint(t *testing.T) {
	type args struct {
		value int64
		exp   int32
	}
	tests := []struct {
		name string
		args args
		want *Millimeters
	}{
		{
			name: "NewMillimetersFixedPoint",
			args: args{
				value: 1287,
				exp:   -3,
			},
			want: &Millimeters{
				Val: testutil.DecimalPtr(decimal.New(1287, -3)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMillimetersFixedPoint(tt.args.value, tt.args.exp)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMillimetersFixedPoint() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewMillimetersFromInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want *Millimeters
	}{
		{
			name: "NewMillimetersFromInt",
			args: args{
				value: 184,
			},
			want: &Millimeters{
				Val: testutil.DecimalPtr(decimal.NewFromInt(184)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMillimetersFromInt(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMillimetersFromInt() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewMillimetersFromUint64(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want *Millimeters
	}{
		{
			name: "NewMillimetersFromUint64",
			args: args{
				value: 11113,
			},
			want: &Millimeters{
				Val: testutil.DecimalPtr(decimal.NewFromUint64(11113)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMillimetersFromUint64(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMillimetersFromUint64() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewMillimetersFromFloat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *Millimeters
	}{
		{
			name: "NewMillimetersFromFloat",
			args: args{
				value: 12.18,
			},
			want: &Millimeters{
				Val: testutil.DecimalPtr(decimal.NewFromFloat(12.18)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewMillimetersFromFloat(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewMillimetersFromFloat() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
