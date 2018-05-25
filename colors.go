package gamut

import (
	colorful "github.com/lucasb-eyer/go-colorful"
)

type Color struct {
	Name         string
	Color        colorful.Color
	WikipediaURL string
}
type Colors []Color

var (
	Monokai   Palette
	Wikipedia Palette
	Crayola   Palette
)

func Hex(s string) colorful.Color {
	c, _ := colorful.Hex(s)
	return c
}

func Complementary(c colorful.Color) colorful.Color {
	h, s, l := c.Hsl()
	h -= 180
	if h < 0 {
		h += 360
	}

	return colorful.Hsl(h, s, l)
}
