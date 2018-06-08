package palette

import (
	"image/color"
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

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

func TestClamped(t *testing.T) {
	cc := []color.Color{gamut.Hex("#FFBE00"), gamut.Hex("#FB6E79"), gamut.Hex("#0095B7")}
	exp := []color.Color{gamut.Hex("#FFBF00"), gamut.Hex("#FA6E79"), gamut.Hex("#0095B6")}

	c := Wikipedia.Clamped(cc)
	for i := range c {
		c1, _ := colorful.MakeColor(c[i].Color)
		c2, _ := colorful.MakeColor(exp[i])

		if c1.Hex() != c2.Hex() {
			t.Errorf("Expected %s, got %s", c1.Hex(), c2.Hex())
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
