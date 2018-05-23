package gamut

import (
	"image/color"
	"sort"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/xrash/smetrics"
)

type Palette struct {
	colors []Color
}

// MixedWith mixes two palettes
func (g *Palette) MixedWith(p Palette) *Palette {
	g.colors = append(g.colors, p.colors...)
	return g
}

// Colors returns the Gamut's palette
func (g *Palette) Colors() []Color {
	return g.colors
}

// Name returns the name of the closest matching color
func (g *Palette) Name(color color.Color) (Color, float64) {
	var d float64 = -1
	var m Color

	c, _ := colorful.MakeColor(color)

	for _, v := range g.colors {
		if nd := v.Color.DistanceLab(c); nd < d || d == -1 {
			d = nd
			m = v
		}
	}

	return m, d
}

// Filter returns colors matching name
func (g *Palette) Filter(name string) []Color {
	s := strings.ToLower(name)
	var c []Color

	for _, v := range g.colors {
		if strings.Contains(strings.ToLower(v.Name), s) {
			c = append(c, v)
		}
	}

	sort.Slice(c, func(i, j int) bool {
		di := smetrics.WagnerFischer(strings.ToLower(c[i].Name), s, 1, 1, 2)
		dj := smetrics.WagnerFischer(strings.ToLower(c[j].Name), s, 1, 1, 2)
		return di < dj
	})

	return c
}
