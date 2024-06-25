package valuetext

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/musicxml/datatype"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestMillimeters_UnmarshalText(t *testing.T) {
	type args struct {
		b []byte
	}
	tests := []struct {
		name    string
		m       *Millimeters
		args    args
		wantErr bool
		want    Millimeters
	}{
		{
			name: "empty",
			m:    (*Millimeters)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("")),
			},
			wantErr: false,
			want:    "",
		},
		{
			name: "invalid",
			m:    (*Millimeters)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("G103.11")),
			},
			wantErr: true,
			want:    "",
		},
		{
			name: "valid",
			m:    (*Millimeters)(testutil.ToStringPtr("")),
			args: args{
				b: ([]byte("105.01")),
			},
			wantErr: false,
			want:    (Millimeters)("105.01"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.m.UnmarshalText(tt.args.b); (err != nil) != tt.wantErr {
				t.Errorf("Millimeters.UnmarshalText() error = %v, wantErr %v", err, tt.wantErr)
			}
			if diff := cmp.Diff(*tt.m, tt.want); diff != "" {
				t.Errorf("Millimeters.UnmarshalText() value is mismatch (-m, +want):%s\n", diff)
			}
		})
	}
}

func TestMillimeters_ToEntityData(t *testing.T) {
	tests := []struct {
		name string
		m    Millimeters
		want datatype.Millimeters
	}{
		{
			name: "empty",
			m:    "",
			want: datatype.Millimeters{},
		},
		{
			name: "invalid",
			m:    (Millimeters)("G105.01"),
			want: datatype.Millimeters{},
		},
		{
			name: "valid",
			m:    (Millimeters)("105.01"),
			want: *datatype.NewMillimetersFixedPoint(10501, -2),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.m.ToEntityData()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Millimeters.ToEntityData() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
