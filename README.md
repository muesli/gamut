# gamut

Go package to generate and manage color palettes & schemes

## Palettes

| Name      | Colors | Source                                                      |
| --------- | ------:| ----------------------------------------------------------- |
| Wikipedia |   1609 | https://en.wikipedia.org/wiki/List_of_colors_(compact)      |
| Crayola   |    180 | https://en.wikipedia.org/wiki/List_of_Crayola_crayon_colors |
| Resene    |    759 | http://www.resene.co.nz                                     |
| Monokai   |     17 |                                                             |

### Examples

#### Generating Color Palettes

```go
import "github.com/muesli/gamut"

colors, err := gamut.Generate(8, gamut.PastelGenerator{})
// returns a slice of 8 pastel colors
```

![Pastel Palette](https://github.com/muesli/gamut/blob/master/docs/palette_pastel.png)

Instead of `gamut.PastelGenerator` you can also use `gamut.WarmGenerator` or
`gamut.HappyGenerator`.

#### Name A Color

```go
name, distance := gamut.Wikipedia.Name(color)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

#### Retrieving Colors

```go
colors = gamut.Crayola.Filter("Red")
// returns a slice of all "Red" colors in the Crayola palette
colors = gamut.Crayola.Colors()
// returns a slice of all colors in the Crayola palette
```

#### Mixing Palettes

```go
p = gamut.Crayola.MixedWith(gamut.Monokai)
// returns a palette with all colors from both Crayola and Monokai
```

## Themes

| Name    | Colors |
| ------- | ------:|
| Monokai |      9 |

### Examples

#### Roles

```go
color = gamut.MonokaiTheme.Role(gamut.Foreground)
```

Available roles are `Foreground`, `Background`, `Base`, `AlternateBase`, `Text`,
`Selection`, `Highlight`.

## Colors

### Examples

#### Complementary Colors

```go
c = gamut.Complementary(color)
```

#### Contrast Colors

```go
c = gamut.Contrast(color)
```

#### Warm/Cool Colors

```go
b = gamut.Warm(color)
b = gamut.Cool(color)
// either true or false
```

#### Hue Offsets

```go
colors = gamut.Triadic(color)
// slice of triadic colors
colors = gamut.Quadratic(color)
// slice of quadratic colors
colors = gamut.Tetradic(color1, color2)
// slice of tetradic colors
colors = gamut.Analogous(color)
// slice of analogous colors
colors = gamut.SplitComplementary(color)
// slice of split-complementary colors
```

#### Shades & Tints

```go
colors = gamut.Shades(color, 8)
// returns a slice of 8 shades, from color to black
colors = gamut.Tints(color, 8)
// returns a slice of 8 tints, from color to white
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
