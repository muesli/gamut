package gamut

import "image/color"

// A Theme is a collection of colors mapped to a role (or function)
type Theme struct {
	Name string

	colors map[Role]Color
}

// MonokaiTheme is a popular theme used for syntax highlighting
var (
	MonokaiTheme = Theme{
		Name:   "monokai",
		colors: make(map[Role]Color),
	}
)

// Role returns the theme's color for a specific role
func (t Theme) Role(r Role) Color {
	return t.colors[r]
}

// Colors returns all (unique) colors used in this theme
func (t Theme) Colors() Colors {
	cm := make(map[color.Color]Color)
	for _, c := range t.colors {
		if _, ok := cm[c.Color]; !ok {
			cm[c.Color] = c
		}
	}

	var cc Colors
	for _, c := range cm {
		cc = append(cc, c)
	}

	return cc
}

func init() {
	MonokaiTheme.colors[Foreground] = Monokai.Filter("Extra White")[0]
	MonokaiTheme.colors[Background] = Monokai.Filter("Caviar")[0]
	MonokaiTheme.colors[Base] = Monokai.Filter("Caviar")[0]
	MonokaiTheme.colors[AlternateBase] = Monokai.Filter("Caviar Dark")[0]
	MonokaiTheme.colors[Text] = Monokai.Filter("Cocoon")[0]
	MonokaiTheme.colors[Selection] = Monokai.Filter("Armadillo")[0]
	MonokaiTheme.colors[Highlight] = Monokai.Filter("El Paso")[0]
}
