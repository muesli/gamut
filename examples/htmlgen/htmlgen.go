package htmlgen

import (
	"bytes"
	"fmt"
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/gamut"
)

var (
	header = `
	<html>
	<body>
	`
	tablecaption = `
	<h2>%s</h2>
	`
	tableheader = `
	<table width="100%" height="64">
	<tr>
	`
	cell        = `<td bgcolor="%s" align="center"><font face="Courier" color="%s">%s</font></td>`
	tablefooter = `
	</tr>
	</table>
	`
	footer = `
	</body>
	</html>
	`
)

func Header(buffer *bytes.Buffer) {
	buffer.Write([]byte(header))
}

func Footer(buffer *bytes.Buffer) {
	buffer.Write([]byte(footer))
}

func Cell(buffer *bytes.Buffer, c color.Color) {
	col, _ := colorful.MakeColor(c)
	comp, _ := colorful.MakeColor(gamut.Contrast(c))
	buffer.Write([]byte(fmt.Sprintf(cell, col.Hex(), comp.Hex(), col.Hex())))
}

func Table(buffer *bytes.Buffer, name string, cc []color.Color) {
	if name != "" {
		buffer.Write([]byte(fmt.Sprintf(tablecaption, name)))
	}
	buffer.Write([]byte(tableheader))
	for _, c := range cc {
		Cell(buffer, c)
	}
	buffer.Write([]byte(tablefooter))
}
