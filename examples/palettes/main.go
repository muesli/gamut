package main

import (
	"bytes"
	"fmt"
	"image/color"
	"io/ioutil"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

var (
	c     = gamut.Hex("#2F1B82")
	cells = 8
)

var (
	header = `
	<html>
	<body bgcolor="white" style="margin: 0; padding: 0;">
		<table width="100%" height="48">
			<tr>
	`
	cell   = `<td bgcolor="%s" />`
	footer = `
			</tr>
		</table>
	</body>
	</html>
	`
)

func palette(name string, cc []color.Color) {
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(header))

	for _, c := range cc {
		col, _ := colorful.MakeColor(c)
		buffer.Write([]byte(fmt.Sprintf(cell, col.Hex())))
	}

	buffer.Write([]byte(footer))
	ioutil.WriteFile(fmt.Sprintf("palette_%s.html", name), buffer.Bytes(), 0644)
}

func main() {
	// generators
	cc, _ := gamut.Generate(cells, gamut.SimilarHueGenerator{Color: c})
	palette("similarhue", cc)
	cc, _ = gamut.Generate(cells, gamut.PastelGenerator{})
	palette("pastel", cc)
	cc, _ = gamut.Generate(cells, gamut.HappyGenerator{})
	palette("happy", cc)
	cc, _ = gamut.Generate(cells, gamut.WarmGenerator{})
	palette("warm", cc)

	// angular
	palette("triadic", gamut.Triadic(c))
	palette("quadratic", gamut.Quadratic(c))
	palette("tetradic", gamut.Tetradic(c, gamut.HueOffset(c, 60)))
	palette("analogous", gamut.Analogous(c))
	palette("splitcomplementary", gamut.SplitComplementary(c))

	// color wheel
	palette("monochromatic", gamut.Monochromatic(c, cells))
	palette("shades", gamut.Shades(c, cells))
	palette("tints", gamut.Tints(c, cells))
	palette("tones", gamut.Tones(c, cells))

	// blends
	palette("blends", gamut.Blends(c, gamut.HueOffset(c, 90), cells))
}
