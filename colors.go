package gamut

import (
	"image/color"

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

// Complementary returns the complementary value for any given color
func Complementary(c color.Color) color.Color {
	col, _ := colorful.MakeColor(c)

	h, s, l := col.Hsl()
	h -= 180
	if h < 0 {
		h += 360
	}

	return colorful.Hsl(h, s, l)
}
