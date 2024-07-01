package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/shopspring/decimal"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNonNegativeDecimal_String(t *testing.T) {
	type fields struct {
		NonNegativeDecimal *decimal.Decimal
		Stringer           fmt.Stringer
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
				NonNegativeDecimal: testutil.DecimalPtr(decimal.New(1, 1)),
			},
			want: "10",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &NonNegativeDecimal{
				Val:      tt.fields.NonNegativeDecimal,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NonNegativeDecimal.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNonNegativeDecimal_StringPtr(t *testing.T) {
	type fields struct {
		NonNegativeDecimal *decimal.Decimal
		Stringer           fmt.Stringer
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
				NonNegativeDecimal: testutil.DecimalPtr(decimal.New(1, 1)),
			},
			want: testutil.ToStringPtr("10"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &NonNegativeDecimal{
				Val:      tt.fields.NonNegativeDecimal,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NonNegativeDecimal.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNonNegativeDecimal_Value(t *testing.T) {
	type fields struct {
		NonNegativeDecimal *decimal.Decimal
		Stringer           fmt.Stringer
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
				NonNegativeDecimal: testutil.DecimalPtr(decimal.New(2, 3)),
				Stringer:           nil,
			},
			want: testutil.DecimalPtr(decimal.New(2, 3)),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tr := &NonNegativeDecimal{
				Val:      tt.fields.NonNegativeDecimal,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				tr = nil
			}
			got := tr.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NonNegativeDecimal.Value() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewNonNegativeDecimalFromString(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *NonNegativeDecimal
		wantErr bool
	}{
		{
			name: "NewNonNegativeDecimalFromString empty",
			args: args{
				value: " ",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewNonNegativeDecimalFromString invalid",
			args: args{
				value: "111,48",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewNonNegativeDecimalFromString negative",
			args: args{
				value: "-111.48",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewNonNegativeDecimalFromString",
			args: args{
				value: "111.48",
			},
			want: &NonNegativeDecimal{
				Val: testutil.DecimalPtr(decimal.New(11148, -2)),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewNonNegativeDecimalFromString(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewNonNegativeDecimalFromString() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewNonNegativeDecimalFromString() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewNonNegativeDecimalFixedPoint(t *testing.T) {
	type args struct {
		value int64
		exp   int32
	}
	tests := []struct {
		name string
		args args
		want *NonNegativeDecimal
	}{
		{
			name: "NewNonNegativeDecimalFixedPoint negative",
			args: args{
				value: -1287,
				exp:   -3,
			},
			want: nil,
		},
		{
			name: "NewNonNegativeDecimalFixedPoint",
			args: args{
				value: 1287,
				exp:   -3,
			},
			want: &NonNegativeDecimal{
				Val: testutil.DecimalPtr(decimal.New(1287, -3)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNonNegativeDecimalFixedPoint(tt.args.value, tt.args.exp)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewNonNegativeDecimalFixedPoint() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewNonNegativeDecimalFromInt(t *testing.T) {
	type args struct {
		value int64
	}
	tests := []struct {
		name string
		args args
		want *NonNegativeDecimal
	}{
		{
			name: "NewNonNegativeDecimalFromInt negative",
			args: args{
				value: -184,
			},
			want: nil,
		},
		{
			name: "NewNonNegativeDecimalFromInt",
			args: args{
				value: 184,
			},
			want: &NonNegativeDecimal{
				Val: testutil.DecimalPtr(decimal.NewFromInt(184)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNonNegativeDecimalFromInt(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewNonNegativeDecimalFromInt() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewNonNegativeDecimalFromUint64(t *testing.T) {
	type args struct {
		value uint64
	}
	tests := []struct {
		name string
		args args
		want *NonNegativeDecimal
	}{
		{
			name: "NewNonNegativeDecimalFromUint64",
			args: args{
				value: 11113,
			},
			want: &NonNegativeDecimal{
				Val: testutil.DecimalPtr(decimal.NewFromUint64(11113)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNonNegativeDecimalFromUint64(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewNonNegativeDecimalFromUint64() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewNonNegativeDecimalFromFloat(t *testing.T) {
	type args struct {
		value float64
	}
	tests := []struct {
		name string
		args args
		want *NonNegativeDecimal
	}{
		{
			name: "NewNonNegativeDecimalFromFloat negative",
			args: args{
				value: -12.18,
			},
			want: nil,
		},
		{
			name: "NewNonNegativeDecimalFromFloat",
			args: args{
				value: 12.18,
			},
			want: &NonNegativeDecimal{
				Val: testutil.DecimalPtr(decimal.NewFromFloat(12.18)),
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewNonNegativeDecimalFromFloat(tt.args.value)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewNonNegativeDecimalFromFloat() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
