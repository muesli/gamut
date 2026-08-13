package theme

import (
	"testing"

	"github.com/muesli/gamut"
	"github.com/muesli/gamut/palette"
)

func TestThemeColors(t *testing.T) {
	cc := MonokaiTheme.Colors()
	exp := 6
	if len(cc) != exp {
		t.Errorf("Expected %d results, got %d", exp, len(cc))
	}
}

func TestThemeRoles(t *testing.T) {
	c := MonokaiTheme.Role(Foreground)
	exp, ok := palette.Monokai.Color("Extra White")
	if !ok {
		t.Fatal("expected 'Extra White' to exist in the Monokai palette")
	}
	if c.Color != exp {
		t.Errorf("Expected role color %s, got %s", exp, c.Color)
	}
}

func TestNewMonokaiTheme(t *testing.T) {
	tm, err := NewMonokaiTheme()
	if err != nil {
		t.Fatalf("NewMonokaiTheme returned unexpected error: %v", err)
	}
	if tm.Name != "monokai" {
		t.Errorf("expected theme name 'monokai', got %q", tm.Name)
	}
	if len(tm.Colors()) != 6 {
		t.Errorf("expected 6 unique colors, got %d", len(tm.Colors()))
	}
}

// A missing role color must surface as an error, not an index-out-of-range panic.
func TestNewThemeMissingColor(t *testing.T) {
	_, err := NewTheme("missing", palette.Monokai, map[Role]string{
		Foreground: "This Color Does Not Exist",
	})
	if err == nil {
		t.Error("expected an error for a missing role color, got nil")
	}
}

// NewTheme should accept an injected palette rather than reaching into a global.
func TestNewThemeInjectsPalette(t *testing.T) {
	p := gamut.Palette{}
	p.AddColors(gamut.Colors{
		{Name: "fg", Color: gamut.Hex("#ffffff")},
		{Name: "bg", Color: gamut.Hex("#000000")},
	})

	tm, err := NewTheme("injected", p, map[Role]string{Foreground: "fg", Background: "bg"})
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if tm.Role(Foreground).Color != gamut.Hex("#ffffff") {
		t.Error("injected palette foreground not resolved")
	}
	if tm.Role(Background).Color != gamut.Hex("#000000") {
		t.Error("injected palette background not resolved")
	}
}
