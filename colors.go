package gamut

import (
	"image/color"
	"math"

	colorful "github.com/lucasb-eyer/go-colorful"
)

// A Color is a color including its name and reference URL
type Color struct {
	Name         string
	Color        color.Color
	WikipediaURL string
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

	h, s, l := col.Hsl()
	h += float64(degrees)
	if h < 0 {
		h += 360
	} else if h > 360 {
		h -= 360
	}

	return colorful.Hsl(h, s, l).Clamped()
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
	wf := colorful.Color{1, 1, 1}
	bf := colorful.Color{0, 0, 0}

	_, _, l := col.Hsl()
	if l < 0.5 {
		return wf
	}

	return bf
}

// Shades returns the specified amount of a color's shades
func Shades(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	h, s, l := col.Hsl()

	dl := l / float64(count+1)

	var cc []color.Color
	for i := 1; i <= count; i++ {
		nl := math.Max(l-dl*float64(i), 0)
		cc = append(cc, colorful.Hsl(h, s, nl).Clamped())
	}

	return cc
}

// Tints returns the specified amount of a color's tints
func Tints(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	h, s, l := col.Hsl()

	dl := (1.0 - l) / float64(count+1)

	var cc []color.Color
	for i := 1; i <= count; i++ {
		nl := math.Max(l+dl*float64(i), 0)
		cc = append(cc, colorful.Hsl(h, s, nl).Clamped())
	}

	return cc
}

// Tones returns the specified amount of a color's tone
func Tones(c color.Color, count int) []color.Color {
	col, _ := colorful.MakeColor(c)
	grey := colorful.Color{0.5, 0.5, 0.5}

	dl := 1.0 / float64(count+1)

	var cc []color.Color
	for i := 1; i <= count; i++ {
		cc = append(cc, col.BlendLab(grey, dl*float64(i)))
	}

	return cc
}

// Cool returns whether a color is considered to have a cool temperature
func Cool(c color.Color) bool {
	col, _ := colorful.MakeColor(c)
	h, _, _ := col.Hsl()

	return 90 <= h && h < 270
}

// Warm returns whether a color is considered to have a warm temperature
func Warm(c color.Color) bool {
	return !Cool(c)
}

// Lighter returns a lighter version of the specified color
func Lighter(c color.Color, percent float64) color.Color {
	col, _ := colorful.MakeColor(c)
	h, s, l := col.Hsl()

	return colorful.Hsl(h, s, l+(l*percent))
}

// Darker returns a darker version of the specified color
func Darker(c color.Color, percent float64) color.Color {
	col, _ := colorful.MakeColor(c)
	h, s, l := col.Hsl()

	return colorful.Hsl(h, s, l-(l*percent))
}
