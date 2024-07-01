package valuetext

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestNonNegativeDecimal_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		nnd     *NonNegativeDecimal
		args    args
		wantErr bool
		want    NonNegativeDecimal
	}{
		{
			name: "empty",
			nnd:  (*NonNegativeDecimal)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("")),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			nnd:  (*NonNegativeDecimal)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("G103.11")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "invalid negative",
			nnd:  (*NonNegativeDecimal)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("-103.11")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			nnd:  (*NonNegativeDecimal)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("105.01")),
			},
			wantErr: false,
			want:    (NonNegativeDecimal)("105.01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.nnd.UnmarshalText(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("NonNegativeDecimal.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.nnd, tt.want); diff != "" {
				t.Errorf("NonNegativeDecimal.UnmarshalText() value is mismatch (-nnd, +want):%s\n", diff)
			}
		})
	}
}

func TestNonNegativeDecimal_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		nnd  NonNegativeDecimal
		want datatype.NonNegativeDecimal
	}{
		{
			name: "empty",
			nnd:  "",
			want: datatype.NonNegativeDecimal{},
		},
		{
			name: "invalid",
			nnd:  (NonNegativeDecimal)("G105.01"),
			want: datatype.NonNegativeDecimal{},
		},
		{
			name: "invalid negative",
			nnd:  (NonNegativeDecimal)("-105.01"),
			want: datatype.NonNegativeDecimal{},
		},
		{
			name: "valid",
			nnd:  (NonNegativeDecimal)("105.01"),
			want: *datatype.NewNonNegativeDecimalFixedPoint(10501, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.nnd.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NonNegativeDecimal.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
