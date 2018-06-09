package gamut

import (
	"image/color"
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

var (
	p1, p2, p3 Palette
)

func init() {
	p1.AddColors(
		Colors{
			{"Spray", Hex("#66D9EF"), ""},
			{"Tree Poppy", Hex("#FD971F"), ""},
			{"Armadillo", Hex("#49483E"), ""},
			{"El Paso", Hex("#3E3D32"), ""},
			{"Center Stage", Hex("#A6E22E"), ""},
			{"Feverish Pink", Hex("#F92672"), ""},
			{"Lavish Lavender", Hex("#AE81FF"), ""},
			{"Funky Yellow", Hex("#E6DB74"), ""},
			{"Cocoon", Hex("#75715E"), ""},
		})
	p2.AddColors(
		Colors{
			{"Extra White", Hex("#F8F8F2"), ""},
			{"Caviar", Hex("#272822"), ""},
			{"Caviar Dark", Hex("#141411"), ""},
			{"Blue Beyond", Hex("#89BDFF"), ""},
			{"Urbane Bronze", Hex("#595959"), ""},
			{"Tricorn Black", Hex("#383830"), ""},
			{"Soothing White", Hex("#E6E6E6"), ""},
			{"Ice Plant", Hex("#FD5FF1"), ""},
		})
	p3.AddColors(
		Colors{
			{"Spray", Hex("#66D9EF"), ""},
			{"Extra White", Hex("#F8F8F2"), ""},
		})
}

func TestDistance(t *testing.T) {
	tt := []struct {
		Hex string
		Exp string
	}{
		{"#66D9EF", "Spray"},
		{"#66D9EE", "Spray"},
		{"#3E3D32", "El Paso"},
		{"#3F3D32", "El Paso"},
	}

	for _, test := range tt {
		c, _ := colorful.Hex(test.Hex)
		m, d := p1.Name(c)

		if m[0].Name != test.Exp {
			t.Errorf("Expected %v, got %v, distance %f", test.Exp, m, d)
		}
	}
}

func TestClamped(t *testing.T) {
	cc := []color.Color{Hex("#66D9EE"), Hex("#3F3D32")}
	exp := []color.Color{Hex("#66D9EF"), Hex("#3E3D32")}

	c := p1.Clamped(cc)
	for i := range c {
		c1, _ := colorful.MakeColor(c[i].Color)
		c2, _ := colorful.MakeColor(exp[i])

		if c1.Hex() != c2.Hex() {
			t.Errorf("Expected %s, got %s", c1.Hex(), c2.Hex())
		}
	}
}

func TestColors(t *testing.T) {
	c := len(p1.Colors())
	exp := 9
	if c != exp {
		t.Errorf("Expected %d colors in Crayola palette, got %d", exp, c)
	}
}

func TestMixing(t *testing.T) {
	p := p1.MixedWith(p2)
	c := len(p.Colors())
	exp := len(p1.Colors()) + len(p2.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}

	p = p1.MixedWith(p3)
	c = len(p.Colors())
	exp = 10 // some p3 colors are duped in the p1 palette
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}

	p = p1.MixedWith(p1)
	c = len(p1.Colors())
	exp = len(p1.Colors())
	if c != exp {
		t.Errorf("Expected %d colors in palette mix, got %d", exp, c)
	}
}

func TestFilter(t *testing.T) {
	cc := p1.Filter("el")
	exp := 2
	if len(cc) != exp {
		t.Errorf("Expected %d results, got %d", exp, len(cc))
	}
}
