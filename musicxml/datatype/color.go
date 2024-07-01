package datatype

import (
	"fmt"
	"image/color"
	"regexp"
	"strconv"
	"strings"
)

type Color struct {
	WebColorCode string
	isARGB       bool
	isUpper      bool
	Val          color.Color
	fmt.Stringer
}

func (c *Color) String() string {
	if c == nil {
		return ""
	}
	return c.WebColorCode
}

func (c *Color) StringPtr() *string {
	if c == nil {
		return nil
	}
	return &c.WebColorCode
}

func (c *Color) IsARGB() bool {
	if c == nil {
		return false
	}
	return c.isARGB
}

func (c *Color) IsUpper() bool {
	if c == nil {
		return false
	}
	return c.isUpper
}

func (c *Color) Value() color.Color {
	if c == nil {
		return nil
	}
	return c.Val
}

func ValidateColorCodeRequireUpper(wcc string) error {
	return validateColorCode(wcc, true, true)
}

func ValidateColorCodeRequireLower(wcc string) error {
	return validateColorCode(wcc, true, false)
}

func ValidateColorCodeCaseInsensitive(wcc string) error {
	return validateColorCode(wcc, false, false)
}

func NewColorRequireUpper(wcc string) (*Color, error) {
	if strings.TrimSpace(wcc) == "" {
		return nil, nil
	}
	if err := ValidateColorCodeRequireUpper(wcc); err != nil {
		return nil, err
	}
	return newColor(wcc, true)
}
func NewColorRequireLower(wcc string) (*Color, error) {
	if strings.TrimSpace(wcc) == "" {
		return nil, nil
	}
	if err := ValidateColorCodeRequireLower(wcc); err != nil {
		return nil, err
	}
	return newColor(wcc, false)
}
func NewColorForceUpper(wcc string) (*Color, error) {
	if strings.TrimSpace(wcc) == "" {
		return nil, nil
	}
	if err := ValidateColorCodeCaseInsensitive(wcc); err != nil {
		return nil, err
	}
	return newColor(wcc, true)
}
func NewColorForceLower(wcc string) (*Color, error) {
	if strings.TrimSpace(wcc) == "" {
		return nil, nil
	}
	if err := ValidateColorCodeCaseInsensitive(wcc); err != nil {
		return nil, err
	}
	return newColor(wcc, false)
}

func NewColorFromARGB(c color.RGBA, isARGB bool, isUpper bool) *Color {
	wcc := ""

	if isARGB {
		wcc = fmt.Sprintf("#%02x%02x%02x%02x", c.A, c.R, c.G, c.B)
	} else {
		c.A = 255
		wcc = fmt.Sprintf("#%02x%02x%02x", c.R, c.G, c.B)
	}
	if isUpper {
		wcc = strings.ToUpper(wcc)
	}
	return &Color{
		WebColorCode: wcc,
		isARGB:       isARGB,
		isUpper:      isUpper,
		Val:          c,
	}
}

func validateColorCode(wcc string, isCaceSensitive bool, isUpper bool) error {
	var re *regexp.Regexp
	rep := ""

	if isCaceSensitive && isUpper {
		rep = "^#([0-9A-F]{6}|[0-9A-F]{8})$"
		re = regexp.MustCompile(rep)
	} else if isCaceSensitive && !isUpper {
		rep = "^#([0-9a-f]{6}|[0-9a-f]{8})$"
		re = regexp.MustCompile(rep)
	} else {
		rep = "^#([0-9A-Fa-f]{6}|[0-9A-Fa-f]{8})$"
		re = regexp.MustCompile(rep)
	}
	m := re.MatchString(wcc)
	if !m {
		return fmt.Errorf("invalid web color code, web color code=%s pattern=%s", wcc, rep)
	}
	return nil
}

func newColor(wcc string, isUpper bool) (*Color, error) {
	if isUpper {
		wcc = strings.ToUpper(wcc)
	} else {
		wcc = strings.ToLower(wcc)
	}
	if len(wcc) == 9 {
		r, g, b, a := convertColorCodeToARGB(wcc)
		return &Color{
			WebColorCode: wcc,
			isARGB:       true,
			isUpper:      isUpper,
			Val: color.RGBA{
				R: r,
				G: g,
				B: b,
				A: a,
			},
		}, nil
	} else {
		r, g, b := convertColorCodeToRGB(wcc)
		return &Color{
			WebColorCode: wcc,
			isARGB:       false,
			isUpper:      isUpper,
			Val: color.RGBA{
				R: r,
				G: g,
				B: b,
				A: 255,
			},
		}, nil
	}
}

func convertColorCodeToRGB(wcc string) (uint8, uint8, uint8) {
	rep := "^#(?P<r>[0-9A-Fa-f]{2})(?P<g>[0-9A-Fa-f]{2})(?P<b>[0-9A-Fa-f]{2})$"
	re := regexp.MustCompile(rep)

	matches := re.FindAllStringSubmatch(wcc, -1)

	r, _ := strconv.ParseUint(matches[0][re.SubexpIndex("r")], 16, 8)
	g, _ := strconv.ParseUint(matches[0][re.SubexpIndex("g")], 16, 8)
	b, _ := strconv.ParseUint(matches[0][re.SubexpIndex("b")], 16, 8)
	return uint8(r), uint8(g), uint8(b)
}

func convertColorCodeToARGB(wcc string) (uint8, uint8, uint8, uint8) {
	rep := "^#(?P<a>[0-9A-Fa-f]{2})(?P<r>[0-9A-Fa-f]{2})(?P<g>[0-9A-Fa-f]{2})(?P<b>[0-9A-Fa-f]{2})$"
	re := regexp.MustCompile(rep)

	matches := re.FindAllStringSubmatch(wcc, -1)

	r, _ := strconv.ParseUint(matches[0][re.SubexpIndex("r")], 16, 8)
	g, _ := strconv.ParseUint(matches[0][re.SubexpIndex("g")], 16, 8)
	b, _ := strconv.ParseUint(matches[0][re.SubexpIndex("b")], 16, 8)
	a, _ := strconv.ParseUint(matches[0][re.SubexpIndex("a")], 16, 8)
	return uint8(r), uint8(g), uint8(b), uint8(a)
}
