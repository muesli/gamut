package htmlgen

import (
	"fmt"
	"image/color"
	"io"

	"github.com/muesli/gamut"

	colorful "github.com/lucasb-eyer/go-colorful"
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

// Header writes the HTML header to buffer
func Header(buffer io.Writer) {
	buffer.Write([]byte(header))
}

// Footer writes the HTML footer to buffer
func Footer(buffer io.Writer) {
	buffer.Write([]byte(footer))
}

// Cell writes a colored HTML table cell to buffer
func Cell(buffer io.Writer, c color.Color) {
	col, _ := colorful.MakeColor(c)
	comp, _ := colorful.MakeColor(gamut.Contrast(c))
	buffer.Write([]byte(fmt.Sprintf(cell, col.Hex(), comp.Hex(), col.Hex())))
}

// Table writes a palette of colors as an HTML table to buffer
func Table(buffer io.Writer, name string, cc []color.Color) {
	if name != "" {
		buffer.Write([]byte(fmt.Sprintf(tablecaption, name)))
	}
	buffer.Write([]byte(tableheader))
	for _, c := range cc {
		Cell(buffer, c)
	}
	buffer.Write([]byte(tablefooter))
}
