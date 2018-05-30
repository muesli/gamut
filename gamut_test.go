package gamut

import (
	"image/color"
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestWarmCool(t *testing.T) {
	cols := []struct {
		hex  string
		cool bool
	}{
		{"#2f1b82", true},
		{"#ff1b82", false},
	}

	for _, col := range cols {
		c, _ := colorful.Hex(col.hex)
		if Warm(c) == col.cool {
			t.Errorf("Expected warm for %s to be %t, got %t", col.hex, !col.cool, col.cool)
		}
		if Cool(c) != col.cool {
			t.Errorf("Expected cool for %s to be %t, got %t", col.hex, col.cool, !col.cool)
		}
	}
}

func TestComplementary(t *testing.T) {
	c, _ := colorful.Hex("#2f1b82")
	cc, _ := colorful.MakeColor(Complementary(c))
	exp, _ := colorful.Hex("#6e821b")

	if cc.Hex() != exp.Hex() {
		t.Errorf("Expected complementary color %v, got %v", exp.Hex(), cc.Hex())
	}
}

func TestContrast(t *testing.T) {
	cols := []struct {
		hex      string
		contrast string
	}{
		{"#2f1b82", "#ffffff"},
		{"#ff1b82", "#000000"},
	}

	for _, col := range cols {
		c, _ := colorful.Hex(col.hex)
		cc, _ := colorful.MakeColor(Contrast(c))
		exp, _ := colorful.Hex(col.contrast)

		if cc.Hex() != exp.Hex() {
			t.Errorf("Expected contrast color %v, got %v", exp.Hex(), cc.Hex())
		}
	}
}

func TestHueOffsets(t *testing.T) {
	cols := []struct {
		fn  func(color.Color) []color.Color
		hex string
		exp []string
	}{
		{Triadic, "#2f1b82", []string{"#1b822f", "#822f1b"}},
		{Quadratic, "#2f1b82", []string{"#1b8263", "#6e821b", "#821b3a"}},
		{Analogous, "#2f1b82", []string{"#1b3a82", "#621b82"}},
		{SplitComplementary, "#2f1b82", []string{"#82621b", "#3a821b"}},
	}

	for coli, col := range cols {
		cc := col.fn(Hex(col.hex))

		for i := 0; i < len(cc); i++ {
			colc, _ := colorful.MakeColor(cc[i])
			expc, _ := colorful.Hex(col.exp[i])
			if expc.Hex() != colc.Hex() {
				t.Errorf("Expected offset color %v, got %v (iteration %d)", expc.Hex(), colc.Hex(), coli)
			}
		}
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
