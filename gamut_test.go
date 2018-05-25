package gamut

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestComplementary(t *testing.T) {
	c, _ := colorful.Hex("#2F1B82")
	cc, _ := colorful.MakeColor(Complementary(c))
	exp, _ := colorful.Hex("#6E821B")

	if cc.Hex() != exp.Hex() {
		t.Errorf("Expected complementary color %v, got %v", exp.Hex(), cc.Hex())
	}
}

func TestDistance(t *testing.T) {
	tt := []struct {
		Hex string
		Exp string
	}{
		{"#FFBF00", "Amber"},
		{"#FFBE00", "Amber"},
		{"#FA6E79", "Begonia"},
		{"#FB6E79", "Begonia"},
		{"#0095B6", "Bondi blue"},
		{"#0095B7", "Bondi blue"},
		{"#F92672", "Neon pink"},
		{"#AE81FF", "Medium purple"},
		{"#66D9EF", "Sky blue (Crayola)"},
		{"#E6DB74", "Straw"},
	}

	for _, test := range tt {
		c, _ := colorful.Hex(test.Hex)
		m, d := Wikipedia.Name(c)

		if m[0].Name != test.Exp {
			t.Errorf("Expected %v, got %v, distance %f", test.Exp, m, d)
		}
	}
}

func TestColors(t *testing.T) {
	c := len(Crayola.Colors())
	exp := 180
	if c != exp {
		t.Errorf("Expected %d colors in Crayola palette, got %d", exp, c)
	}
}

func TestMixing(t *testing.T) {
	p := Crayola.MixedWith(Resene)
	c := len(p.Colors())
	exp := len(Crayola.Colors()) + len(Resene.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}

	p = Wikipedia.MixedWith(Crayola)
	c = len(p.Colors())
	exp = 1766 // some Crayola colors are duped in the Wikipedia palette
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}

	p = Crayola.MixedWith(Crayola)
	c = len(p.Colors())
	exp = len(Crayola.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}
}

func TestFilter(t *testing.T) {
	cc := Wikipedia.Filter("Pastel")
	exp := 16
	if len(cc) != exp {
		t.Errorf("Expected %d results, got %d", exp, len(cc))
	}
}

func TestThemeColors(t *testing.T) {
	cc := MonokaiTheme.Colors()
	exp := 6
	if len(cc) != exp {
		t.Errorf("Expected %d results, got %d", exp, len(cc))
	}
}
