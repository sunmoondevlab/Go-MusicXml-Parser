package valuetext

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestTenths_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		t       *Tenths
		args    args
		wantErr bool
		want    Tenths
	}{
		{
			name: "empty",
			t:    (*Tenths)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("")),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			t:    (*Tenths)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("G103.11")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			t:    (*Tenths)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("105.01")),
			},
			wantErr: false,
			want:    (Tenths)("105.01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.t.UnmarshalText(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Tenths.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.t, tt.want); diff != "" {
				t.Errorf("Tenths.UnmarshalText() value is mismatch (-t, +want):%s\n", diff)
			}
		})
	}
}

func TestTenths_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		t    Tenths
		want datatype.Tenths
	}{
		{
			name: "empty",
			t:    "",
			want: datatype.Tenths{},
		},
		{
			name: "invalid",
			t:    (Tenths)("G105.01"),
			want: datatype.Tenths{},
		},
		{
			name: "valid",
			t:    (Tenths)("105.01"),
			want: *datatype.NewTenthsFixedPoint(10501, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.t.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Tenths.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
