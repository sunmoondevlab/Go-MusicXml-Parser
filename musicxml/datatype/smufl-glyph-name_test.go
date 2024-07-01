package datatype

import (
	"fmt"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestSmuflGlyphName_String(t *testing.T) {
	type fields struct {
		Val      *string
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
				Val: testutil.ToStringPtr("fClefFrench"),
			},
			want: "fClefFrench",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgn := &SmuflGlyphName{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				sgn = nil
			}
			got := sgn.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestSmuflGlyphName_StringPtr(t *testing.T) {
	type fields struct {
		Val      *string
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
				Val: testutil.ToStringPtr("fClefFrench"),
			},
			want: testutil.ToStringPtr("fClefFrench"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sgn := &SmuflGlyphName{
				Val:      tt.fields.Val,
				Stringer: tt.fields.Stringer,
			}
			if tt.isNil {
				sgn = nil
			}
			got := sgn.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("SmuflGlyphName.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewSmuflGlyphName(t *testing.T) {
	type args struct {
		value string
	}
	tests := []struct {
		name    string
		args    args
		want    *SmuflGlyphName
		wantErr bool
	}{
		{
			name: "NewSmuflGlyphName empty",
			args: args{
				value: " ",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "NewSmuflGlyphName invalid",
			args: args{
				value: "aaaaa",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "NewSmuflGlyphName",
			args: args{
				value: "fClefSmall",
			},
			want: &SmuflGlyphName{
				Val: testutil.ToStringPtr("fClefSmall"),
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewSmuflGlyphName(tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewSmuflGlyphName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("NewSmuflGlyphName() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
