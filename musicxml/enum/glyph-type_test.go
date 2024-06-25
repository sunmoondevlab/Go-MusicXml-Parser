package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToGlyphType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *GlyphTypeEnum
		wantErr bool
	}{
		{
			name: "quarter-rest",
			args: args{
				t: "quarter-rest"},
			want:    &GlyphType.QuarterRest,
			wantErr: false,
		},
		{
			name: "g-clef-ottava-bassa",
			args: args{
				t: "g-clef-ottava-bassa"},
			want:    &GlyphType.GClefOttavaBassa,
			wantErr: false,
		},
		{
			name: "c-clef",
			args: args{
				t: "c-clef"},
			want:    &GlyphType.CClef,
			wantErr: false,
		},
		{
			name: "f-clef",
			args: args{
				t: "f-clef"},
			want:    &GlyphType.FClef,
			wantErr: false,
		},
		{
			name: "percussion-clef",
			args: args{
				t: "percussion-clef"},
			want:    &GlyphType.PercussionClef,
			wantErr: false,
		},
		{
			name: "octave-shift-up-8",
			args: args{
				t: "octave-shift-up-8"},
			want:    &GlyphType.OctaveShiftUp8,
			wantErr: false,
		},
		{
			name: "octave-shift-down-8",
			args: args{
				t: "octave-shift-down-8"},
			want:    &GlyphType.OctaveShiftDown8,
			wantErr: false,
		},
		{
			name: "octave-shift-continue-8",
			args: args{
				t: "octave-shift-continue-8"},
			want:    &GlyphType.OctaveShiftContinue8,
			wantErr: false,
		},
		{
			name: "octave-shift-up-15",
			args: args{
				t: "octave-shift-up-15"},
			want:    &GlyphType.OctaveShiftUp15,
			wantErr: false,
		},
		{
			name: "octave-shift-down-15",
			args: args{
				t: "octave-shift-down-15"},
			want:    &GlyphType.OctaveShiftDown15,
			wantErr: false,
		},
		{
			name: "octave-shift-continue-15",
			args: args{
				t: "octave-shift-continue-15"},
			want:    &GlyphType.OctaveShiftContinue15,
			wantErr: false,
		},
		{
			name: "octave-shift-up-22",
			args: args{
				t: "octave-shift-up-22"},
			want:    &GlyphType.OctaveShiftUp22,
			wantErr: false,
		},
		{
			name: "octave-shift-down-22",
			args: args{
				t: "octave-shift-down-22"},
			want:    &GlyphType.OctaveShiftDown22,
			wantErr: false,
		},
		{
			name: "octave-shift-continue-22",
			args: args{
				t: "octave-shift-continue-22"},
			want:    &GlyphType.OctaveShiftContinue22,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "f-cle"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToGlyphType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToGlyphType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToGlyphType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllGlyphTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []GlyphTypeEnum
	}{
		{
			name: "",
			want: []GlyphTypeEnum{GlyphType.QuarterRest,
				GlyphType.GClefOttavaBassa,
				GlyphType.CClef,
				GlyphType.FClef,
				GlyphType.PercussionClef,
				GlyphType.OctaveShiftUp8,
				GlyphType.OctaveShiftDown8,
				GlyphType.OctaveShiftContinue8,
				GlyphType.OctaveShiftUp15,
				GlyphType.OctaveShiftDown15,
				GlyphType.OctaveShiftContinue15,
				GlyphType.OctaveShiftUp22,
				GlyphType.OctaveShiftDown22,
				GlyphType.OctaveShiftContinue22,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllGlyphTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllGlyphTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestGlyphTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj GlyphTypeEnum
	}
	tests := []struct {
		name string
		e    *GlyphTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &GlyphType.OctaveShiftContinue15,
			args: args{
				obj: GlyphType.OctaveShiftContinue22,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &GlyphType.OctaveShiftContinue15,
			args: args{
				obj: GlyphType.OctaveShiftContinue15,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GlyphTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestGlyphTypeEnum_In(t *testing.T) {
	type args struct {
		objs []GlyphTypeEnum
	}
	tests := []struct {
		name string
		e    *GlyphTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &GlyphType.OctaveShiftUp8,
			args: args{
				objs: []GlyphTypeEnum{GlyphType.OctaveShiftUp15, GlyphType.OctaveShiftUp22},
			},
			want: false,
		},
		{
			name: "in",
			e:    &GlyphType.OctaveShiftUp8,
			args: args{
				objs: []GlyphTypeEnum{GlyphType.OctaveShiftUp8, GlyphType.OctaveShiftUp15, GlyphType.OctaveShiftUp22},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GlyphTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestGlyphTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *GlyphTypeEnum
		want string
	}{
		{
			name: "quarter-rest",
			e:    &GlyphType.QuarterRest,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GlyphTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestGlyphTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *GlyphTypeEnum
		want string
	}{
		{
			name: "c-clef",
			e:    &GlyphType.CClef,
			want: "c-clef",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GlyphTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestGlyphTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *GlyphTypeEnum
		want *string
	}{
		{
			name: "c-clef",
			e:    &GlyphType.CClef,
			want: testutil.ToStringPtr("c-clef"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("GlyphTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
