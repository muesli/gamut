package gamut

import (
	"testing"
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
	if c != Monokai.Filter("Extra White")[0] {
		t.Errorf("Expected role color %s, got %s", Monokai.Filter("Extra White")[0], c)
	}
}
