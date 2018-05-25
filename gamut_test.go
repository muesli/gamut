package gamut

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestComplementary(t *testing.T) {
	c, _ := colorful.Hex("#2F1B82")
	exp, _ := colorful.Hex("#6E821B")

	if cc := Complementary(c); cc.Hex() != exp.Hex() {
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
	c := len(Monokai.Colors())
	if c != 17 {
		t.Errorf("Expected 17 colors in monokai palette, got %d", c)
	}
}

func TestMixing(t *testing.T) {
	p := Monokai.MixedWith(Crayola)
	c := len(p.Colors())
	exp := len(Monokai.Colors()) + len(Crayola.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}

	p = Monokai.MixedWith(Monokai)
	c = len(p.Colors())
	exp = len(Monokai.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}
}

func TestFilter(t *testing.T) {
	cc := Wikipedia.Filter("Pastel")
	if len(cc) != 16 {
		t.Errorf("Expected 16 results, got %d", len(cc))
	}
}
