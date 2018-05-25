package gamut

import (
	colorful "github.com/lucasb-eyer/go-colorful"
)

// A Color is a color including its name and reference URL
type Color struct {
	Name         string
	Color        colorful.Color
	WikipediaURL string
}

// Colors is a slice of colors
type Colors []Color

// Hex returns the color encoded by a hex-string, e.g. "#ABCDEF"
func Hex(s string) colorful.Color {
	c, _ := colorful.Hex(s)
	return c
}

// Complementary returns the complementary value for any given color
func Complementary(c colorful.Color) colorful.Color {
	h, s, l := c.Hsl()
	h -= 180
	if h < 0 {
		h += 360
	}

	return colorful.Hsl(h, s, l)
}
