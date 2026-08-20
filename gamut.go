package gamut

import (
	"image/color"
	"maps"
	"sort"
	"strings"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/xrash/smetrics"
)

// A Palette is a collection of colors.
type Palette struct {
	colors map[color.Color]Colors
	names  map[string]color.Color
}

// MixedWith mixes two palettes.
func (g Palette) MixedWith(p Palette) Palette {
	np := Palette{}
	np.AddColors(g.Colors())
	np.AddColors(p.Colors())
	return np
}

// Clone returns an independent copy of the palette. A plain value copy of a
// Palette shares its interior maps, so mutating the copy would affect the
// original. Clone detaches those maps (and the per-key Colors slices) so the
// result can be mutated freely.
func (g Palette) Clone() Palette {
	colors := make(map[color.Color]Colors, len(g.colors))
	for k, v := range g.colors {
		colors[k] = append(Colors{}, v...)
	}
	names := maps.Clone(g.names)
	return Palette{colors: colors, names: names}
}

// AddColors adds colors to the palette.
func (g *Palette) AddColors(cc Colors) {
	if g.colors == nil {
		g.colors = make(map[color.Color]Colors)
	}
	if g.names == nil {
		g.names = make(map[string]color.Color)
	}

	for _, c := range cc {
		found := false
		for _, v := range g.colors[c.Color] {
			if v.Name == c.Name {
				found = true
				break
			}
		}

		if !found {
			g.colors[c.Color] = append(g.colors[c.Color], c)
		}

		g.names[strings.ToUpper(c.Name)] = c.Color
	}
}

// Colors returns the Palette's colors.
func (g Palette) Colors() Colors {
	var r Colors
	for _, c := range g.colors {
		r = append(r, c...)
	}
	return r
}

// Clamped expects a slice of colors and returns a slice of the nearest matching
// colors from the palette. Colors for which no match can be found (e.g. an
// empty palette) are skipped.
func (g Palette) Clamped(cc []color.Color) Colors {
	var r Colors
	for _, c := range cc {
		nm, _ := g.Name(c)
		if len(nm) == 0 {
			continue
		}
		r = append(r, nm[0])
	}
	return r
}

// Color returns the color with a specific name.
func (g Palette) Color(name string) (color.Color, bool) {
	c, ok := g.names[strings.ToUpper(name)]
	return c, ok
}

// Name returns the name of the closest matching color.
func (g Palette) Name(color color.Color) (Colors, float64) {
	var d float64 = -1
	var m Colors

	c, _ := colorful.MakeColor(color)

	for _, v := range g.colors {
		col, _ := colorful.MakeColor(v[0].Color)
		if nd := col.DistanceLab(c); nd < d || d == -1 {
			d = nd
			m = Colors{}
			m = append(m, v...)
		}
	}

	return m, d
}

// Filter returns colors matching name.
func (g Palette) Filter(name string) Colors {
	s := strings.ToLower(name)
	var c Colors

	for _, v := range g.colors {
		for _, vv := range v {
			if strings.Contains(strings.ToLower(vv.Name), s) {
				c = append(c, vv)
			}
		}
	}

	sort.SliceStable(c, func(i, j int) bool {
		if c[i].Name == c[j].Name {
			return ToHex(c[i].Color) < ToHex(c[j].Color)
		}
		di := smetrics.WagnerFischer(strings.ToLower(c[i].Name), s, 1, 1, 2)
		dj := smetrics.WagnerFischer(strings.ToLower(c[j].Name), s, 1, 1, 2)
		return di < dj
	})

	return c
}
