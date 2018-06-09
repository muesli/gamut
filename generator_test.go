package gamut

import (
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

type TestColorGenerator struct {
	BroadGranularity
}

func (cc TestColorGenerator) Valid(col colorful.Color) bool {
	l, _, _ := col.Lab()
	return 0.4 <= l && l <= 0.6
}

func TestGenerator(t *testing.T) {
	tt := []struct {
		generator ColorGenerator
		amount    int
	}{
		{PastelGenerator{}, 8},
		{WarmGenerator{}, 8},
		{HappyGenerator{}, 8},
		{SimilarHueGenerator{Color: Hex("#ABCDEF")}, 8},
		{TestColorGenerator{}, 8},
	}

	for _, g := range tt {
		cc, err := Generate(g.amount, g.generator)
		if err != nil {
			t.Error(err)
			return
		}

		if len(cc) != g.amount {
			t.Errorf("Expected %d colors from generator, got %d", g.amount, len(cc))
		}
	}
}
