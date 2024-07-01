package enum

import (
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestToLineWidthType(t *testing.T) {
	type args struct {
		t string
	}
	tests := []struct {
		name    string
		args    args
		want    *LineWidthTypeEnum
		wantErr bool
	}{
		{
			name: "beam",
			args: args{
				t: "beam"},
			want:    &LineWidthType.Beam,
			wantErr: false,
		},
		{
			name: "bracket",
			args: args{
				t: "bracket"},
			want:    &LineWidthType.Bracket,
			wantErr: false,
		},
		{
			name: "dashes",
			args: args{
				t: "dashes"},
			want:    &LineWidthType.Dashes,
			wantErr: false,
		},
		{
			name: "enclosure",
			args: args{
				t: "enclosure"},
			want:    &LineWidthType.Enclosure,
			wantErr: false,
		},
		{
			name: "ending",
			args: args{
				t: "ending"},
			want:    &LineWidthType.Ending,
			wantErr: false,
		},
		{
			name: "extend",
			args: args{
				t: "extend"},
			want:    &LineWidthType.Extend,
			wantErr: false,
		},
		{
			name: "dashes",
			args: args{
				t: "dashes"},
			want:    &LineWidthType.Dashes,
			wantErr: false,
		},
		{
			name: "heavy barline",
			args: args{
				t: "heavy barline"},
			want:    &LineWidthType.HeavyBarline,
			wantErr: false,
		},
		{
			name: "dashes",
			args: args{
				t: "dashes"},
			want:    &LineWidthType.Dashes,
			wantErr: false,
		},
		{
			name: "leger",
			args: args{
				t: "leger"},
			want:    &LineWidthType.Leger,
			wantErr: false,
		},
		{
			name: "light barline",
			args: args{
				t: "light barline"},
			want:    &LineWidthType.LightBarline,
			wantErr: false,
		},
		{
			name: "dashes",
			args: args{
				t: "dashes"},
			want:    &LineWidthType.Dashes,
			wantErr: false,
		},
		{
			name: "octave shift",
			args: args{
				t: "octave shift"},
			want:    &LineWidthType.OctaveShift,
			wantErr: false,
		},
		{
			name: "pedal",
			args: args{
				t: "pedal"},
			want:    &LineWidthType.Pedal,
			wantErr: false,
		},
		{
			name: "slur middle",
			args: args{
				t: "slur middle"},
			want:    &LineWidthType.SlurMiddle,
			wantErr: false,
		},
		{
			name: "slur tip",
			args: args{
				t: "slur tip"},
			want:    &LineWidthType.SlurTip,
			wantErr: false,
		},
		{
			name: "staff",
			args: args{
				t: "staff"},
			want:    &LineWidthType.Staff,
			wantErr: false,
		},
		{
			name: "stem",
			args: args{
				t: "stem"},
			want:    &LineWidthType.Stem,
			wantErr: false,
		},
		{
			name: "tie middle",
			args: args{
				t: "tie middle"},
			want:    &LineWidthType.TieMiddle,
			wantErr: false,
		},
		{
			name: "tie tip",
			args: args{
				t: "tie tip"},
			want:    &LineWidthType.TieTip,
			wantErr: false,
		},
		{
			name: "tuplet bracket",
			args: args{
				t: "tuplet bracket"},
			want:    &LineWidthType.TupletBracket,
			wantErr: false,
		},
		{
			name: "wedge",
			args: args{
				t: "wedge"},
			want:    &LineWidthType.Wedge,
			wantErr: false,
		},
		{
			name: "invalid",
			args: args{
				t: "ra"},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ToLineWidthType(tt.args.t)
			if (err != nil) != tt.wantErr {
				t.Errorf("ToLineWidthType() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("ToLineWidthType() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestAllLineWidthTypeEnumValues(t *testing.T) {
	tests := []struct {
		name string
		want []LineWidthTypeEnum
	}{
		{
			name: "",
			want: []LineWidthTypeEnum{LineWidthType.Beam,
				LineWidthType.Bracket,
				LineWidthType.Dashes,
				LineWidthType.Enclosure,
				LineWidthType.Ending,
				LineWidthType.Extend,
				LineWidthType.HeavyBarline,
				LineWidthType.Leger,
				LineWidthType.LightBarline,
				LineWidthType.OctaveShift,
				LineWidthType.Pedal,
				LineWidthType.SlurMiddle,
				LineWidthType.SlurTip,
				LineWidthType.Staff,
				LineWidthType.Stem,
				LineWidthType.TieMiddle,
				LineWidthType.TieTip,
				LineWidthType.TupletBracket,
				LineWidthType.Wedge,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := AllLineWidthTypeEnumValues()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("AllLineWidthTypeEnumValues() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLineWidthTypeEnum_Equals(t *testing.T) {
	type args struct {
		obj LineWidthTypeEnum
	}
	tests := []struct {
		name string
		e    *LineWidthTypeEnum
		args args
		want bool
	}{
		{
			name: "ne",
			e:    &LineWidthType.Beam,
			args: args{
				obj: LineWidthType.Bracket,
			},
			want: false,
		},
		{
			name: "eq",
			e:    &LineWidthType.Bracket,
			args: args{
				obj: LineWidthType.Bracket,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Equals(tt.args.obj)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LineWidthTypeEnum.Equals() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLineWidthTypeEnum_In(t *testing.T) {
	type args struct {
		objs []LineWidthTypeEnum
	}
	tests := []struct {
		name string
		e    *LineWidthTypeEnum
		args args
		want bool
	}{
		{
			name: "ni",
			e:    &LineWidthType.Bracket,
			args: args{
				objs: []LineWidthTypeEnum{LineWidthType.Dashes, LineWidthType.Enclosure},
			},
			want: false,
		},
		{
			name: "in",
			e:    &LineWidthType.Beam,
			args: args{
				objs: []LineWidthTypeEnum{LineWidthType.Beam, LineWidthType.Dashes, LineWidthType.Enclosure},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.In(tt.args.objs...)
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LineWidthTypeEnum.In() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLineWidthTypeEnum_Ordinal(t *testing.T) {
	tests := []struct {
		name string
		e    *LineWidthTypeEnum
		want string
	}{
		{
			name: "ord",
			e:    &LineWidthType.Beam,
			want: "0",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.Ordinal()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LineWidthTypeEnum.Ordinal() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLineWidthTypeEnum_String(t *testing.T) {
	tests := []struct {
		name string
		e    *LineWidthTypeEnum
		want string
	}{
		{
			name: "beam",
			e:    &LineWidthType.Beam,
			want: "beam",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LineWidthTypeEnum.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestLineWidthTypeEnum_StringPtr(t *testing.T) {
	tests := []struct {
		name string
		e    *LineWidthTypeEnum
		want *string
	}{
		{
			name: "beam",
			e:    &LineWidthType.Beam,
			want: testutil.ToStringPtr("beam"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.e.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("LineWidthTypeEnum.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}
