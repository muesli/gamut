package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

var (
	header = `
	<html>
	<body>
		<table width="100%" height="100%">
			<tr>
	`

	cell = `
				<td bgcolor="%s" />
	`

	footer = `
			</tr>
		</table>
	</body>
	</html>
	`
)

func main() {
	// Write HTML header
	buffer := bytes.NewBuffer([]byte{})
	buffer.Write([]byte(header))

	cc, err := gamut.Generate(8, gamut.PastelGenerator{})
	if err != nil {
		panic(err)
	}

	for _, c := range cc {
		col, ok := colorful.MakeColor(c)
		if !ok {
			panic(fmt.Sprintf("invalid RGB color: %s", c))
		}

		fmt.Println("Color as Hex:", col.Hex())
		buffer.Write([]byte(fmt.Sprintf(cell, col.Hex())))
	}

	// Write HTML footer and generate palette.html
	buffer.Write([]byte(footer))
	ioutil.WriteFile("palette.html", buffer.Bytes(), 0644)
}
