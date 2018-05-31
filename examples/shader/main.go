package main

import (
	"bytes"
	"image/color"
	"io/ioutil"

	"github.com/muesli/gamut"
	"github.com/muesli/gamut/examples/htmlgen"
)

func main() {
	buffer := bytes.NewBuffer([]byte{})
	htmlgen.Header(buffer)

	colors, err := gamut.Generate(8, gamut.PastelGenerator{})
	if err != nil {
		panic(err)
	}

	// print tints
	var tints [][]color.Color
	for x := 0; x < len(colors); x++ {
		tints = append(tints, gamut.Tints(colors[x], 8))
	}
	for y := len(tints[0]) - 1; y >= 0; y-- {
		var cc []color.Color
		for x := 0; x < len(colors); x++ {
			cc = append(cc, tints[x][y])
		}
		htmlgen.Table(buffer, "", cc)
	}

	// print original colors
	htmlgen.Table(buffer, "", colors)

	// print shades
	var shades [][]color.Color
	for x := 0; x < len(colors); x++ {
		shades = append(shades, gamut.Shades(colors[x], 8))
	}
	for y := 0; y < len(colors); y++ {
		var cc []color.Color
		for x := 0; x < len(shades[0]); x++ {
			cc = append(cc, shades[x][y])
		}
		htmlgen.Table(buffer, "", cc)
	}

	htmlgen.Footer(buffer)
	ioutil.WriteFile("palette.html", buffer.Bytes(), 0644)
}
