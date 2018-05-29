# gamut

Go package to generate and manage color palettes & schemes

## Examples

### Generating Color Palettes

```go
import "github.com/muesli/gamut"
)

cc, err := gamut.Generate(8, gamut.Pastel{})
for i, c := range cc {
	fmt.Printf("Color #%d: %s\n", i, c)
}
```

Instead of `gamut.Pastel` you can also use `gamut.Warm` or `gamut.Happy`.

### Complementary Colors

```go
c := gamut.Complementary(color)
```

### Name A Color

```go
name, distance := gamut.Wikipedia.Name(color)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

### Retrieving Colors

```go
cc := gamut.Crayola.Filter("Red")
// cc is a slice of all "Red" colors in the Crayola palette

cc := gamut.Crayola.Colors()
// cc is a slice of all colors in the Crayola palette
```

### Mixing Palettes

```go
p := gamut.Crayola.MixedWith(gamut.Monokai)
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
