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
colors, err := gamut.Generate(8, gamut.SimilarHueGenerator{Color: color})
// returns a slice of 8 colors with a hue similar to the given color
```

![Pastel Palette](https://github.com/muesli/gamut/blob/master/docs/palette_pastel.png)

Instead of `gamut.PastelGenerator` you can also use `gamut.WarmGenerator` or
`gamut.HappyGenerator`. Using the `ColorGenerator` interface, you can write your
own color generators.

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
colors = gamut.Monokai.Clamped(colors)
// returns a slice of the nearest matching colors in the Monokai palette
```

#### Mixing Palettes

```go
p = gamut.Crayola.MixedWith(gamut.Monokai)
// returns a palette with all colors from both Crayola and Monokai
```

## Themes

| Name    | Colors |
| ------- | ------:|
| Monokai |      7 |

### Examples

#### Roles

```go
color = gamut.MonokaiTheme.Role(gamut.Foreground)
```

Available roles are `Foreground`, `Background`, `Base`, `AlternateBase`, `Text`,
`Selection`, `Highlight`.

## Colors

### Examples

#### Lighter/Darker Colors

The `Darker` and `Lighter` functions darken or lighten a given color value by
a specified percentage, without changing the color's hue.

```go
c = gamut.Darker(color, 0.1)
// returns a 10% darker version of color
c = gamut.Lighter(color, 0.3)
// returns a 30% lighter version of color
```

#### Complementary Colors

The `Complementary` function returns the complementary color for a given color.

```go
c = gamut.Complementary(color)
```

#### Contrast Colors

The `Contrast` function returns the color with the highest contrast to a given
color, either black or white.

```go
c = gamut.Contrast(color)
```

#### Warm/Cool Colors

```go
b = gamut.Warm(color)
b = gamut.Cool(color)
// either true or false
```

#### Around the Color Wheel

All the following functions return colors of a different hue, but with the same
lightness and saturation as the given colors.

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

#### Shades, Tints & Tones

```go
colors = gamut.Monochromatic(color, 8)
// returns colors of the same hue, but with a different saturation/lightness
colors = gamut.Shades(color, 8)
// returns a slice of 8 shades, from the given color to black
colors = gamut.Tints(color, 8)
// returns a slice of 8 tints, from the given color to white
colors = gamut.Tones(color, 8)
// returns a slice of 8 tones, from the given color to gray
```

![Shades, Tints and Tones](https://github.com/muesli/gamut/blob/master/docs/shades_tints_tones.png)

#### Blending Colors

```go
colors = gamut.Blends(color1, color2, 8)
// returns a slice of interpolated colors by blending two colors
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
