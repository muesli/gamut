package gamut

import (
	"testing"
)

func TestGenerator(t *testing.T) {
	cc, err := Generate(8, PastelGenerator{})
	if err != nil {
		t.Error(err)
		return
	}

	if len(cc) != 8 {
		t.Errorf("Expected 8 colors from generator, got %d", len(cc))
	}
}
