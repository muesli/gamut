package main

import (
	"bytes"
	"io/ioutil"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
	"github.com/muesli/gamut/examples/htmlgen"
)

func main() {
	buffer := bytes.NewBuffer([]byte{})
	htmlgen.Header(buffer)

	c, _ := colorful.Hex("#2F1B82")

	htmlgen.Table(buffer, "Shades", gamut.Shades(c, 8))
	htmlgen.Table(buffer, "Tints", gamut.Tints(c, 8))
	htmlgen.Table(buffer, "Tones", gamut.Tones(c, 8))
	htmlgen.Table(buffer, "Analogous", gamut.Analogous(c))
	htmlgen.Table(buffer, "Split Complementary", gamut.SplitComplementary(c))
	htmlgen.Table(buffer, "Triadic", gamut.Triadic(c))
	htmlgen.Table(buffer, "Quadratic", gamut.Quadratic(c))
	htmlgen.Table(buffer, "Tetradic", gamut.Tetradic(c, gamut.HueOffset(c, 60)))

	htmlgen.Footer(buffer)
	ioutil.WriteFile("palette.html", buffer.Bytes(), 0644)
}
