package datatype

import (
	"fmt"
	"image/color"
	"testing"

	"github.com/google/go-cmp/cmp"
	"github.com/sunmoondevlab/Go-MusicXml-Parser/testutil"
)

func TestColor_String(t *testing.T) {
	type fields struct {
		WebColorCode string
		isARGB       bool
		isUpper      bool
		Color        color.Color
		Stringer     fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   string
		isNil  bool
	}{
		{
			name:   "String",
			fields: fields{},
			want:   "",
			isNil:  true,
		},
		{
			name: "String",
			fields: fields{
				WebColorCode: "#FFFFFFFF",
				isARGB:       true,
				isUpper:      true,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: "#FFFFFFFF",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				WebColorCode: tt.fields.WebColorCode,
				isARGB:       tt.fields.isARGB,
				isUpper:      tt.fields.isUpper,
				Val:          tt.fields.Color,
				Stringer:     tt.fields.Stringer,
			}
			if tt.isNil {
				c = nil
			}
			got := c.String()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Color.String() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestColor_StringPtr(t *testing.T) {
	type fields struct {
		WebColorCode string
		isARGB       bool
		isUpper      bool
		Color        color.Color
		Stringer     fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   *string
		isNil  bool
	}{
		{
			name:   "String",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name: "String",
			fields: fields{
				WebColorCode: "#FFFFFFFF",
				isARGB:       true,
				isUpper:      true,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: testutil.ToStringPtr("#FFFFFFFF"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				WebColorCode: tt.fields.WebColorCode,
				isARGB:       tt.fields.isARGB,
				isUpper:      tt.fields.isUpper,
				Val:          tt.fields.Color,
				Stringer:     tt.fields.Stringer,
			}
			if tt.isNil {
				c = nil
			}
			got := c.StringPtr()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Color.StringPtr() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestColor_IsARGB(t *testing.T) {
	type fields struct {
		WebColorCode string
		isARGB       bool
		isUpper      bool
		Color        color.Color
		Stringer     fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		isNil  bool
	}{
		{
			name:   "rgb",
			fields: fields{},
			want:   false,
			isNil:  true,
		},
		{
			name: "rgb",
			fields: fields{
				WebColorCode: "#FFFFFF",
				isARGB:       false,
				isUpper:      true,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: false,
		},
		{
			name: "argb",
			fields: fields{
				WebColorCode: "#FFFFFFFF",
				isARGB:       true,
				isUpper:      true,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				WebColorCode: tt.fields.WebColorCode,
				isARGB:       tt.fields.isARGB,
				isUpper:      tt.fields.isUpper,
				Val:          tt.fields.Color,
				Stringer:     tt.fields.Stringer,
			}
			if tt.isNil {
				c = nil
			}
			got := c.IsARGB()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Color.IsARGB() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestColor_IsUpper(t *testing.T) {
	type fields struct {
		WebColorCode string
		isARGB       bool
		isUpper      bool
		Color        color.Color
		Stringer     fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   bool
		isNil  bool
	}{
		{
			name:   "upper",
			fields: fields{},
			want:   false,
			isNil:  true,
		},
		{
			name: "upper",
			fields: fields{
				WebColorCode: "#FFFFFF",
				isARGB:       false,
				isUpper:      true,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: true,
		},
		{
			name: "lower",
			fields: fields{
				WebColorCode: "#ffffff",
				isARGB:       false,
				isUpper:      false,
				Color: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				WebColorCode: tt.fields.WebColorCode,
				isARGB:       tt.fields.isARGB,
				isUpper:      tt.fields.isUpper,
				Val:          tt.fields.Color,
				Stringer:     tt.fields.Stringer,
			}
			if tt.isNil {
				c = nil
			}
			got := c.IsUpper()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Color.IsUpper() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestColor_Value(t *testing.T) {
	type fields struct {
		WebColorCode string
		isARGB       bool
		isUpper      bool
		Color        color.Color
		Stringer     fmt.Stringer
	}
	tests := []struct {
		name   string
		fields fields
		want   color.Color
		isNil  bool
	}{
		{
			name:   "ColorValue",
			fields: fields{},
			want:   nil,
			isNil:  true,
		},
		{
			name: "ColorValue",
			fields: fields{
				WebColorCode: "#0F0F0F0F",
				isARGB:       true,
				isUpper:      true,
				Color: color.RGBA{
					R: 15,
					G: 15,
					B: 15,
					A: 15,
				},
			},
			want: color.RGBA{
				R: 15,
				G: 15,
				B: 15,
				A: 15,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Color{
				WebColorCode: tt.fields.WebColorCode,
				isARGB:       tt.fields.isARGB,
				isUpper:      tt.fields.isUpper,
				Val:          tt.fields.Color,
				Stringer:     tt.fields.Stringer,
			}
			if tt.isNil {
				c = nil
			}
			got := c.Value()
			if diff := cmp.Diff(got, tt.want); diff != "" {
				t.Errorf("Color.Value() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestValidateColorCodeRequireUpper(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc: "#00000G",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc: "#FFFFF",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc: "#FFFFFFF",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc: "#afafaf",
			},
			wantErr: true,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc: "#000000",
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc: "#09AF90",
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc: "#FFFFFF",
			},
			wantErr: false,
		},

		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc: "#0000000G",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc: "#FFFFFFF",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc: "#FFFFFFFFF",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc: "#afafafaf",
			},
			wantErr: true,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc: "#00000000",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc: "#0F09AF90",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc: "#FFFFFFFF",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateColorCodeRequireUpper(tt.args.wcc); (err != nil) != tt.wantErr {
				t.Errorf("ValidateColorCodeRequireUpper() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateColorCodeRequireLower(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc: "#00000g",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc: "#fffff",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc: "#fffffff",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc: "#AFAFAF",
			},
			wantErr: true,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc: "#000000",
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc: "#09af90",
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc: "#ffffff",
			},
			wantErr: false,
		},

		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc: "#0000000g",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc: "#fffffff",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc: "#fffffffff",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc: "#AFAFAFAF",
			},
			wantErr: true,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc: "#00000000",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc: "#0009af90",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc: "#ffffffff",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateColorCodeRequireLower(tt.args.wcc); (err != nil) != tt.wantErr {
				t.Errorf("ValidateColorCodeRequireLower() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestValidateColorCodeCaseInsensitive(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc: "#00G00g",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc: "#ffFff",
			},
			wantErr: true,
		},
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc: "#ffFfFff",
			},
			wantErr: true,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc: "#000000",
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc: "#09af90",
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc: "#ffffff",
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesinensitive upper",
			args: args{
				wcc: "#000000",
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive upper",
			args: args{
				wcc: "#09AF90",
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive upper",
			args: args{
				wcc: "#FFFFFF",
			},
			wantErr: false,
		},

		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc: "#0000G00g",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc: "#fFffFff",
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc: "#fFffFfFff",
			},
			wantErr: true,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc: "#00000000",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc: "#0099af90",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc: "#ffffffff",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesinensitive upper",
			args: args{
				wcc: "#00000000",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive upper",
			args: args{
				wcc: "#0F09AF90",
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive upper",
			args: args{
				wcc: "#FFFFFFFF",
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := ValidateColorCodeCaseInsensitive(tt.args.wcc); (err != nil) != tt.wantErr {
				t.Errorf("ValidateColorCodeCaseInsensitive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNewColorRequireUpper(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name    string
		args    args
		want    *Color
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				wcc: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "upper rgb invalid",
			args: args{
				wcc: "#FFFFFFF",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "upper rgb white",
			args: args{
				wcc: "#FFFFFF",
			},
			want: &Color{
				WebColorCode: "#FFFFFF",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower rgb white",
			args: args{
				wcc: "#ffffff",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "upper argb white",
			args: args{
				wcc: "#FFFFFFFF",
			},
			want: &Color{
				WebColorCode: "#FFFFFFFF",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower argb white",
			args: args{
				wcc: "#ffffffff",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "upper rgb",
			args: args{
				wcc: "#EEDDCC",
			},
			want: &Color{
				WebColorCode: "#EEDDCC",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper argb",
			args: args{
				wcc: "#FFEEDDCC",
			},
			want: &Color{
				WebColorCode: "#FFEEDDCC",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColorRequireUpper(tt.args.wcc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColorRequireUpper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColorRequireUpper() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewColorRequireLower(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name    string
		args    args
		want    *Color
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				wcc: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "lower rgb invalid",
			args: args{
				wcc: "#fffffff",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "upper rgb white",
			args: args{
				wcc: "#FFFFFF",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "lower rgb white",
			args: args{
				wcc: "#ffffff",
			},
			want: &Color{
				WebColorCode: "#ffffff",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper argb white",
			args: args{
				wcc: "#FFFFFFFF",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "lower argb white",
			args: args{
				wcc: "#ffffffff",
			},
			want: &Color{
				WebColorCode: "#ffffffff",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper rgb",
			args: args{
				wcc: "#EEDDCC",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "lower rgb",
			args: args{
				wcc: "#eeddcc",
			},
			want: &Color{
				WebColorCode: "#eeddcc",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper argb",
			args: args{
				wcc: "#FFEEDDCC",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "lower argb",
			args: args{
				wcc: "#ffeeddcc",
			},
			want: &Color{
				WebColorCode: "#ffeeddcc",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColorRequireLower(tt.args.wcc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColorRequireLower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColorRequireLower() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewColorForceUpper(t *testing.T) {
	type args struct {
		wcc             string
		isCaceSensitive bool
		isUpper         bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Color
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				wcc: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "rower rgb invalid",
			args: args{
				wcc:             "#fffffff",
				isCaceSensitive: true,
				isUpper:         true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "rower rgba invalid",
			args: args{
				wcc:             "#fffffffff",
				isCaceSensitive: true,
				isUpper:         true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "to upper rgb",
			args: args{
				wcc:             "#eeddcc",
				isCaceSensitive: false,
				isUpper:         true,
			},
			want: &Color{
				WebColorCode: "#EEDDCC",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to upper argb",
			args: args{
				wcc:             "#ffeeddcc",
				isCaceSensitive: false,
				isUpper:         true,
			},
			want: &Color{
				WebColorCode: "#FFEEDDCC",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColorForceUpper(tt.args.wcc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColorForceUpper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColorForceUpper() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewColorForceLower(t *testing.T) {
	type args struct {
		wcc             string
		isCaceSensitive bool
		isUpper         bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Color
		wantErr bool
	}{
		{
			name: "nil",
			args: args{
				wcc: "",
			},
			want:    nil,
			wantErr: false,
		},
		{
			name: "lower rgb invalid",
			args: args{
				wcc:             "#fffffff",
				isCaceSensitive: true,
				isUpper:         true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "lower argb invalid",
			args: args{
				wcc:             "#fffffffff",
				isCaceSensitive: true,
				isUpper:         true,
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "to lower rgb",
			args: args{
				wcc:             "#EEDDCC",
				isCaceSensitive: false,
				isUpper:         false,
			},
			want: &Color{
				WebColorCode: "#eeddcc",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to lower argb",
			args: args{
				wcc:             "#FFEEDDCC",
				isCaceSensitive: false,
				isUpper:         false,
			},
			want: &Color{
				WebColorCode: "#ffeeddcc",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := NewColorForceLower(tt.args.wcc)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColorForceLower() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColorForceLower() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func TestNewColorFromARGB(t *testing.T) {
	type args struct {
		c       color.RGBA
		isARGB  bool
		isUpper bool
	}
	tests := []struct {
		name string
		args args
		want *Color
	}{
		{
			name: "rgb  upper",
			args: args{
				c: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 4,
				},
				isARGB:  false,
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#010203",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
			},
		},
		{
			name: "rgba  upper",
			args: args{
				c: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
				isARGB:  true,
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#FF010203",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
			},
		},
		{
			name: "rgb  lower",
			args: args{
				c: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 4,
				},
				isARGB:  false,
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#010203",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
			},
		},
		{
			name: "rgba  lower",
			args: args{
				c: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
				isARGB:  true,
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#ff010203",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 1,
					G: 2,
					B: 3,
					A: 255,
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := NewColorFromARGB(tt.args.c, tt.args.isARGB, tt.args.isUpper)
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColorFromARGB() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func Test_validateColorCode(t *testing.T) {
	type args struct {
		wcc             string
		isCaceSensitive bool
		isUpper         bool
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc:             "#00000G",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc:             "#FFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc:             "#FFFFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive upper",
			args: args{
				wcc:             "#afafaf",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc:             "#000000",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc:             "#09AF90",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive upper",
			args: args{
				wcc:             "#FFFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc:             "#00000g",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc:             "#fffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc:             "#fffffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB casesensitive lower",
			args: args{
				wcc:             "#AFAFAF",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc:             "#000000",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc:             "#09af90",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesensitive lower",
			args: args{
				wcc:             "#ffffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc:             "#00G00g",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc:             "#ffFff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid RGB caseinsensitive",
			args: args{
				wcc:             "#ffFfFff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc:             "#000000",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc:             "#09af90",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive lower",
			args: args{
				wcc:             "#ffffff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB casesinensitive upper",
			args: args{
				wcc:             "#000000",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive upper",
			args: args{
				wcc:             "#09AF90",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid RGB caseinsensitive upper",
			args: args{
				wcc:             "#FFFFFF",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},

		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc:             "#0000000G",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc:             "#FFFFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc:             "#FFFFFFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive upper",
			args: args{
				wcc:             "#afafafaf",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: true,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc:             "#00000000",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc:             "#0F09AF90",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive upper",
			args: args{
				wcc:             "#FFFFFFFF",
				isCaceSensitive: true,
				isUpper:         true,
			},
			wantErr: false,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc:             "#0000000g",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc:             "#fffffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc:             "#fffffffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB casesensitive lower",
			args: args{
				wcc:             "#AFAFAFAF",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc:             "#00000000",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc:             "#0009af90",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesensitive lower",
			args: args{
				wcc:             "#ffffffff",
				isCaceSensitive: true,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc:             "#0000G00g",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc:             "#fFffFff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "invalid ARGB caseinsensitive",
			args: args{
				wcc:             "#fFffFfFff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: true,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc:             "#00000000",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc:             "#0099af90",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive lower",
			args: args{
				wcc:             "#ffffffff",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB casesinensitive upper",
			args: args{
				wcc:             "#00000000",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive upper",
			args: args{
				wcc:             "#0F09AF90",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
		{
			name: "valid ARGB caseinsensitive upper",
			args: args{
				wcc:             "#FFFFFFFF",
				isCaceSensitive: false,
				isUpper:         false,
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := validateColorCode(tt.args.wcc, tt.args.isCaceSensitive, tt.args.isUpper); (err != nil) != tt.wantErr {
				t.Errorf("ValidateColorCode() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Tes_newColor(t *testing.T) {
	type args struct {
		wcc     string
		isUpper bool
	}
	tests := []struct {
		name    string
		args    args
		want    *Color
		wantErr bool
	}{
		{
			name: "upper rgb white",
			args: args{
				wcc:     "#FFFFFF",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#FFFFFF",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower rgb white",
			args: args{
				wcc:     "#ffffff",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#ffffff",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper argb white",
			args: args{
				wcc:     "#FFFFFFFF",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#FFFFFFFF",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower argb white",
			args: args{
				wcc:     "#ffffffff",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#ffffffff",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 255,
					G: 255,
					B: 255,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper rgb",
			args: args{
				wcc:     "#EEDDCC",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#EEDDCC",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower rgb",
			args: args{
				wcc:     "#eeddcc",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#eeddcc",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "upper argb",
			args: args{
				wcc:     "#FFEEDDCC",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#FFEEDDCC",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "lower argb",
			args: args{
				wcc:     "#ffeeddcc",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#ffeeddcc",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to upper rgb",
			args: args{
				wcc:     "#eeddcc",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#EEDDCC",
				isARGB:       false,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to lower rgb",
			args: args{
				wcc:     "#EEDDCC",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#eeddcc",
				isARGB:       false,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to upper argb",
			args: args{
				wcc:     "#ffeeddcc",
				isUpper: true,
			},
			want: &Color{
				WebColorCode: "#FFEEDDCC",
				isARGB:       true,
				isUpper:      true,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
		{
			name: "to lower argb",
			args: args{
				wcc:     "#FFEEDDCC",
				isUpper: false,
			},
			want: &Color{
				WebColorCode: "#ffeeddcc",
				isARGB:       true,
				isUpper:      false,
				Val: color.RGBA{
					R: 238,
					G: 221,
					B: 204,
					A: 255,
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := newColor(tt.args.wcc, tt.args.isUpper)
			if (err != nil) != tt.wantErr {
				t.Errorf("NewColor() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			opt := cmp.AllowUnexported(Color{})
			if diff := cmp.Diff(got, tt.want, opt); diff != "" {
				t.Errorf("NewColor() value is mismatch (-got, +want):%s\n", diff)
			}
		})
	}
}

func Test_convertColorCodeToRGB(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name string
		args args
		r    uint8
		g    uint8
		b    uint8
	}{

		{
			name: "upper white",
			args: args{
				wcc: "#FFFFFF",
			},
			r: 255,
			g: 255,
			b: 255,
		},
		{
			name: "lower white",
			args: args{
				wcc: "#ffffff",
			},
			r: 255,
			g: 255,
			b: 255,
		},
		{
			name: "upper",
			args: args{
				wcc: "#112233",
			},
			r: 17,
			g: 34,
			b: 51,
		},
		{
			name: "lower",
			args: args{
				wcc: "#ffeedd",
			},
			r: 255,
			g: 238,
			b: 221,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b := convertColorCodeToRGB(tt.args.wcc)
			if r != tt.r {
				t.Errorf("convertColorCodeToRGB() r = %v, want %v", r, tt.r)
			}
			if g != tt.g {
				t.Errorf("convertColorCodeToRGB() g = %v, want %v", g, tt.g)
			}
			if b != tt.b {
				t.Errorf("convertColorCodeToRGB() b = %v, want %v", b, tt.b)
			}
		})
	}
}

func Test_convertColorCodeToARGB(t *testing.T) {
	type args struct {
		wcc string
	}
	tests := []struct {
		name string
		args args
		r    uint8
		g    uint8
		b    uint8
		a    uint8
	}{
		{
			name: "upper white",
			args: args{
				wcc: "#FFFFFFFF",
			},
			r: 255,
			g: 255,
			b: 255,
			a: 255,
		},
		{
			name: "lower white",
			args: args{
				wcc: "#ffffffff",
			},
			r: 255,
			g: 255,
			b: 255,
			a: 255,
		},
		{
			name: "upper",
			args: args{
				wcc: "#11223344",
			},
			r: 34,
			g: 51,
			b: 68,
			a: 17,
		},
		{
			name: "lower",
			args: args{
				wcc: "#ffeeddcc",
			},
			r: 238,
			g: 221,
			b: 204,
			a: 255,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, g, b, a := convertColorCodeToARGB(tt.args.wcc)
			if r != tt.r {
				t.Errorf("convertColorCodeToARGB() r = %v, want %v", r, tt.r)
			}
			if g != tt.g {
				t.Errorf("convertColorCodeToARGB() g = %v, want %v", g, tt.g)
			}
			if b != tt.b {
				t.Errorf("convertColorCodeToARGB() b = %v, want %v", b, tt.b)
			}
			if a != tt.a {
				t.Errorf("convertColorCodeToARGB() a = %v, want %v", b, tt.a)
			}
		})
	}
}
