package gamut

import (
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// A Color is a color including its name and reference URL
type Color struct {
	Name      string
	Color     color.Color
	Reference string
}

// Colors is a slice of colors
type Colors []Color

// Hex returns the color encoded by a hex-string, e.g. "#ABCDEF"
func Hex(s string) color.Color {
	c, _ := colorful.Hex(s)
	return c
}

// HueOffset returns color with a different hue angle
func HueOffset(c color.Color, degrees int) color.Color {
	col, _ := colorful.MakeColor(c)

	h, s, v := col.Hsv()
	h += float64(degrees)
	if h < 0 {
		h += 360
	} else if h > 360 {
		h -= 360
	}

	return colorful.Hsv(h, s, v).Clamped()
}

// Tetradic returns the tetradic values for any given color
func Tetradic(c1 color.Color, c2 color.Color) []color.Color {
	var cc []color.Color
	cc = append(cc, Complementary(c1))
	cc = append(cc, Complementary(c2))
	return cc
}

// Triadic returns the triadic values for any given color
func Triadic(c color.Color) []color.Color {
	var cc []color.Color
	cc = append(cc, HueOffset(c, -120))
	cc = append(cc, HueOffset(c, 120))
	return cc
}

// Quadratic returns the quadratic values for any given color
func Quadratic(c color.Color) []color.Color {
	var cc []color.Color
	cc = append(cc, HueOffset(c, -90))
	cc = append(cc, Complementary(c))
	cc = append(cc, HueOffset(c, 90))
	return cc
}

// Analogous returns the analogous values for any given color
func Analogous(c color.Color) []color.Color {
	var cc []color.Color
	cc = append(cc, HueOffset(c, -30))
	cc = append(cc, HueOffset(c, 30))
	return cc
}

// SplitComplementary returns the split complementary values for any given color
func SplitComplementary(c color.Color) []color.Color {
	return Analogous(Complementary(c))
}

// Complementary returns the complementary value for any given color
func Complementary(c color.Color) color.Color {
	return HueOffset(c, 180)
}

// Contrast returns the color with the most contrast (hence either black or white)
func Contrast(c color.Color) color.Color {
	col, _ := colorful.MakeColor(c)
	wf := colorful.Color{R: 1, G: 1, B: 1}
	bf := colorful.Color{R: 0, G: 0, B: 0}

	_, _, l := col.Hcl()
	if l < 0.5 {
		return wf
	}

	return bf
}

// Monochromatic returns the specified amount of monochromatic colors based on
// a given color's hues
func Monochromatic(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)

	h, s, _ := col.Hsl()
	dl := 1.0 / float64(count+1)

	var cc []color.Color
	for i := 1; i <= count; i++ {
		cc = append(cc, colorful.Hsl(h, s, dl*float64(i)).Clamped())
	}

	return cc
}

// Blends returns a slice of interpolated colors, blended between two colors
func Blends(c1, c2 color.Color, count int) []color.Color {
	col1, _ := colorful.MakeColor(c1)
	col2, _ := colorful.MakeColor(c2)

	dl := 1.0 / float64(count+1)

	var cc []color.Color
	for i := 1; i <= count; i++ {
		cc = append(cc, col1.BlendLab(col2, dl*float64(i)).Clamped())
	}

	return cc
}

// Shades returns the specified amount of a color's shades
func Shades(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	return Blends(col, colorful.Color{R: 0.0, G: 0.0, B: 0.0}, count)
}

// Tints returns the specified amount of a color's tints
func Tints(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	return Blends(col, colorful.Color{R: 1.0, G: 1.0, B: 1.0}, count)
}

// Tones returns the specified amount of a color's tone
func Tones(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	return Blends(col, colorful.Color{R: 0.5, G: 0.5, B: 0.5}, count)
}

// Cool returns whether a color is considered to have a cool temperature
func Cool(c color.Color) bool {
	col, _ := colorful.MakeColor(c)
	h, _, _ := col.Hsv()

	return 90 <= h && h < 270
}

// Warm returns whether a color is considered to have a warm temperature
func Warm(c color.Color) bool {
	return !Cool(c)
}

// Lighter returns a lighter version of the specified color
func Lighter(c color.Color, percent float64) color.Color {
	col, _ := colorful.MakeColor(c)
	h, cv, l := col.Hcl()

	return colorful.Hcl(h, cv, l+(l*percent))
}

// Darker returns a darker version of the specified color
func Darker(c color.Color, percent float64) color.Color {
	col, _ := colorful.MakeColor(c)
	h, cv, l := col.Hcl()

	return colorful.Hcl(h, cv, l-(l*percent))
}
