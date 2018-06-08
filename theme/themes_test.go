package theme

import (
	"testing"

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
	if c != palette.Monokai.Filter("Extra White")[0] {
		t.Errorf("Expected role color %s, got %s", palette.Monokai.Filter("Extra White")[0], c)
	}
}
