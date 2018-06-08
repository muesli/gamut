package theme

import (
	"image/color"

	"github.com/muesli/gamut"
	"github.com/muesli/gamut/palette"
)

// A Theme is a collection of colors mapped to a role (or function)
type Theme struct {
	Name string

	colors map[Role]gamut.Color
}

// MonokaiTheme is a popular theme used for syntax highlighting
var (
	MonokaiTheme = Theme{
		Name:   "monokai",
		colors: make(map[Role]gamut.Color),
	}
)

// Role returns the theme's color for a specific role
func (t Theme) Role(r Role) gamut.Color {
	return t.colors[r]
}

// Colors returns all (unique) colors used in this theme
func (t Theme) Colors() gamut.Colors {
	cm := make(map[color.Color]gamut.Color)
	for _, c := range t.colors {
		if _, ok := cm[c.Color]; !ok {
			cm[c.Color] = c
		}
	}

	var cc gamut.Colors
	for _, c := range cm {
		cc = append(cc, c)
	}

	return cc
}

func init() {
	MonokaiTheme.colors[Foreground] = palette.Monokai.Filter("Extra White")[0]
	MonokaiTheme.colors[Background] = palette.Monokai.Filter("Caviar")[0]
	MonokaiTheme.colors[Base] = palette.Monokai.Filter("Caviar")[0]
	MonokaiTheme.colors[AlternateBase] = palette.Monokai.Filter("Caviar Dark")[0]
	MonokaiTheme.colors[Text] = palette.Monokai.Filter("Cocoon")[0]
	MonokaiTheme.colors[Selection] = palette.Monokai.Filter("Armadillo")[0]
	MonokaiTheme.colors[Highlight] = palette.Monokai.Filter("El Paso")[0]
}
