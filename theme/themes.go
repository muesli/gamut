package theme

import (
	"fmt"
	"image/color"

	"github.com/muesli/gamut"
	"github.com/muesli/gamut/palette"
)

// A Theme is a collection of colors mapped to a role (or function)
type Theme struct {
	Name string

	colors map[Role]gamut.Color
}

// monokaiRoles maps each Role to the exact color name in the Monokai palette.
var monokaiRoles = map[Role]string{
	Foreground:    "Extra White",
	Background:    "Caviar",
	Base:          "Caviar",
	AlternateBase: "Caviar Dark",
	Text:          "Cocoon",
	Selection:     "Armadillo",
	Highlight:     "El Paso",
}

// MonokaiTheme is a popular theme used for syntax highlighting. It is populated
// eagerly and will panic if the underlying palette data is missing or renamed.
// Prefer NewMonokaiTheme for a safe, error-returning variant.
var MonokaiTheme = MustMonokaiTheme()

// MustMonokaiTheme builds the Monokai theme, panicking on any missing color.
// Retained for backward compatibility with the previous init()-based global.
func MustMonokaiTheme() Theme {
	t, err := NewMonokaiTheme()
	if err != nil {
		panic(err)
	}
	return t
}

// NewMonokaiTheme builds the Monokai theme from the Monokai palette, returning
// an error if any role's color is missing or ambiguous.
func NewMonokaiTheme() (Theme, error) {
	return NewTheme("monokai", palette.Monokai, monokaiRoles)
}

// NewTheme builds a Theme by resolving each Role to a color looked up by exact
// name in the provided palette. This injects the palette dependency explicitly
// instead of reaching into a package-level global at init time. It returns an
// error (rather than panicking) when a role name cannot be resolved.
func NewTheme(name string, p gamut.Palette, roles map[Role]string) (Theme, error) {
	colors := make(map[Role]gamut.Color, len(roles))
	for r, nam := range roles {
		c, ok := p.Color(nam)
		if !ok {
			return Theme{}, fmt.Errorf("theme: color %q not found in palette", nam)
		}
		colors[r] = gamut.Color{Name: nam, Color: c}
	}

	return Theme{Name: name, colors: colors}, nil
}

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
