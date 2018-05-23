package gamut

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
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

		if m.Name != test.Exp {
			t.Errorf("Expected %v, got %v, distance %f", test.Exp, m, d)
		}
	}
}

func TestFindByName(t *testing.T) {
	cc := Wikipedia.Filter("Pastel")
	if len(cc) != 16 {
		t.Errorf("Expected 16 results, got %d", len(cc))
	}
}
