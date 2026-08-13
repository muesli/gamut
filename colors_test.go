package gamut

import (
	"image/color"
	"testing"

	colorful "github.com/lucasb-eyer/go-colorful"
)

func TestHueOffsetNormalization(t *testing.T) {
	c := Hex("#ff0000") // hue 0

	// Large offsets must still produce a hue within 0 and 360. A single +/-360
	// wrap only normalizes offsets in (-360, 360).
	for _, deg := range []int{180, 360, 400, -400, 720, -720, 900} {
		got, _ := colorful.MakeColor(HueOffset(c, deg))
		h, _, _ := got.Hsv()
		if h < 0 || h >= 360 {
			t.Errorf("HueOffset(%d) yielded out-of-range hue %f (want 0-360)", deg, h)
		}
	}
}

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

func TestLightness(t *testing.T) {
	lightness := func(c color.Color) float64 {
		cc, _ := colorful.MakeColor(c)
		_, _, l := cc.Hcl()
		return l
	}

	// Lightening increases lightness, darkening decreases it. We assert the
	// property change rather than brittle exact hex values, since HCL->RGB
	// rounding makes hard-coded strings fragile.
	cols := []struct {
		op      string
		percent float64
		hex     string
	}{
		{"lighter", 0.1, "#2f1b82"},
		{"lighter", 0.8, "#ea621f"},
		{"lighter", 0.5, "#000000"},
		{"lighter", 1.0, "#000000"},
		{"darker", 0.3, "#2f1b82"},
		{"darker", 0.8, "#ea621f"},
		{"darker", 0.5, "#ffffff"},
		{"darker", 1.0, "#ffffff"},
	}

	for _, col := range cols {
		c, _ := colorful.Hex(col.hex)
		in := lightness(c)

		var got color.Color
		if col.op == "lighter" {
			got = Lighter(c, col.percent)
		} else {
			got = Darker(c, col.percent)
		}
		out := lightness(got)

		if col.op == "lighter" && out <= in {
			t.Errorf("Lighter(%s, %v): expected lightness to increase, got %f <= %f", col.hex, col.percent, out, in)
		}
		if col.op == "darker" && out >= in {
			t.Errorf("Darker(%s, %v): expected lightness to decrease, got %f >= %f", col.hex, col.percent, out, in)
		}
	}

	// Regression: pure black must lighten (not stay black) and pure white must
	// darken (not stay white) under interpolation toward the extreme.
	black, _ := colorful.Hex("#000000")
	if lightness(Lighter(black, 0.5)) == 0 {
		t.Error("Lighter(#000000, 0.5) remained black (no-op)")
	}
	white, _ := colorful.Hex("#ffffff")
	if lightness(Darker(white, 0.5)) == 1 {
		t.Error("Darker(#ffffff, 0.5) remained white (no-op)")
	}

	// Identity: a zero fraction must return an equal color.
	if ToHex(Lighter(black, 0)) != "#000000" {
		t.Error("Lighter(c, 0) is not identity")
	}
	if ToHex(Darker(white, 0)) != "#ffffff" {
		t.Error("Darker(c, 0) is not identity")
	}
}

func TestComplementary(t *testing.T) {
	cols := []struct {
		hex           string
		complementary string
	}{
		{"#2f1b82", "#6e821b"},
		{"#00ffff", "#ff0000"},
	}

	for _, col := range cols {
		c, _ := colorful.Hex(col.hex)
		cc, _ := colorful.MakeColor(Complementary(c))
		exp, _ := colorful.Hex(col.complementary)

		if cc.Hex() != exp.Hex() {
			t.Errorf("Expected complementary color %v, got %v", exp.Hex(), cc.Hex())
		}
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
		{Quadratic, "#2f1b82", []string{"#1b8262", "#6e821b", "#821b3b"}},
		{Analogous, "#2f1b82", []string{"#1b3b82", "#621b82"}},
		{SplitComplementary, "#2f1b82", []string{"#82631b", "#3a821b"}},
	}

	for coli, col := range cols {
		cc := col.fn(Hex(col.hex))

		for i := 0; i < len(col.exp); i++ {
			colc, _ := colorful.MakeColor(cc[i])
			expc, _ := colorful.Hex(col.exp[i])
			if expc.Hex() != colc.Hex() {
				t.Errorf("Expected offset color %v, got %v (iteration %d)", expc.Hex(), colc.Hex(), coli)
			}
		}
	}
}

func TestLightnessOffsets(t *testing.T) {
	cols := []struct {
		fn    func(color.Color, int) []color.Color
		count int
		hex   string
		exp   []string
	}{
		{Monochromatic, 8, "#2f1b82", []string{"#110a2f", "#22135e", "#331d8d", "#4427bc", "#6043d8", "#8872e2", "#b0a1ec", "#d7d0f5"}},
		{Shades, 8, "#2f1b82", []string{"#2c1973", "#291864", "#251656", "#211447", "#1d123a", "#190f2d", "#150b20", "#0c0514"}},
		{Tints, 8, "#2f1b82", []string{"#4b3290", "#634a9e", "#7a62ac", "#917ab9", "#a794c7", "#bdaed5", "#d3c8e3", "#e9e3f1"}},
		{Tones, 8, "#2f1b82", []string{"#3d2782", "#483282", "#523d82", "#5b4882", "#635382", "#6b5e82", "#726981", "#797480"}},
	}

	for coli, col := range cols {
		cc := col.fn(Hex(col.hex), col.count)

		for i := 0; i < len(col.exp); i++ {
			colc, _ := colorful.MakeColor(cc[i])
			expc, _ := colorful.Hex(col.exp[i])
			if expc.Hex() != colc.Hex() {
				t.Errorf("Expected offset color %v, got %v (iteration %d)", expc.Hex(), colc.Hex(), coli)
			}
		}
	}
}

func TestTetradic(t *testing.T) {
	c1, _ := colorful.Hex("#2f1b82")
	c2 := HueOffset(c1, 60)
	exp1, _ := colorful.Hex("#6e821b")
	exp2, _ := colorful.Hex("#1b822f")

	tc := Tetradic(c1, c2)
	t1, _ := colorful.MakeColor(tc[0])
	t2, _ := colorful.MakeColor(tc[1])

	if t1.Hex() != exp1.Hex() {
		t.Errorf("Expected first tetradic color %v, got %v", exp1.Hex(), t1.Hex())
	}
	if t2.Hex() != exp2.Hex() {
		t.Errorf("Expected second tetradic color %v, got %v", exp2.Hex(), t2.Hex())
	}
}
