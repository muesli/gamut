# gamut

Go package to generate and manage color palettes & schemes

## Examples

### Generating Color Palettes

```go
import "github.com/muesli/gamut"
)

cc, err := gamut.Generate(8, gamut.PastelGenerator{})
for i, c := range cc {
	fmt.Printf("Color #%d: %s\n", i, c)
}
```

<table width="100%" height="64">
<tr>
<td bgcolor="#deaba4" />
<td bgcolor="#b5c9a7" />
<td bgcolor="#beeea3" />
<td bgcolor="#a090e3" />
<td bgcolor="#b58fb1" />
<td bgcolor="#b7e1e9" />
<td bgcolor="#e7a1d7" />
<td bgcolor="#8fa6c8" />
</tr>
</table>

Instead of `gamut.PastelGenerator` you can also use `gamut.WarmGenerator` or
`gamut.HappyGenerator`.

### Name A Color

```go
name, distance := gamut.Wikipedia.Name(color)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

### Retrieving Colors

```go
cc = gamut.Crayola.Filter("Red")
// cc is a slice of all "Red" colors in the Crayola palette
cc = gamut.Crayola.Colors()
// cc is a slice of all colors in the Crayola palette
```

### Complementary Colors

```go
c = gamut.Complementary(color)
```

### Contrast Colors

```go
c = gamut.Contrast(color)
```

### Warm/Cool Colors

```go
b = gamut.Warm(color)
b = gamut.Cool(color)
// either true or false
```

### Hue Offsets

```go
cc = gamut.Triadic(color)
// slice of triadic colors
cc = gamut.Quadratic(color)
// slice of quadratic colors
cc = gamut.Analogous(color)
// slice of analogous colors
cc = gamut.SplitComplementary(color)
// slice of split-complementary colors
```

### Shades & Tints

```go
cc = gamut.Shades(color, 8)
// slice of 8 shades, from color to black
cc = gamut.Tints(color, 8)
// slice of 8 tints, from color to white
```

### Mixing Palettes

```go
p = gamut.Crayola.MixedWith(gamut.Monokai)
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
