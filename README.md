# gamut

Go package to generate and manage color palettes & schemes

```go
import "github.com/muesli/gamut"
```

## Palettes

| Name      | Colors | Source                                                      |
| --------- | ------:| ----------------------------------------------------------- |
| Wikipedia |   1609 | https://en.wikipedia.org/wiki/List_of_colors_(compact)      |
| Crayola   |    180 | https://en.wikipedia.org/wiki/List_of_Crayola_crayon_colors |
| Resene    |    759 | http://www.resene.co.nz                                     |
| Monokai   |     17 |                                                             |

### Examples

#### Generating Color Palettes

Color Generators, like the provided `PastelGenerator`, `WarmGenerator` or
`HappyGenerator` can produce random (within the color space constraits of the
generator) color palettes:

```go
gamut.Generate(8, gamut.PastelGenerator{})
// => ([]color.Color{...}, error)
```

![Pastel Palette](https://github.com/muesli/gamut/blob/master/docs/palette_pastel.png)

The `SimilarHueGenerator` produces colors with a hue similar to a given color:

```go
gamut.Generate(8, gamut.SimilarHueGenerator{Color: gamut.Hex("#2F1B82")})
// => ([]color.Color{...}, error)
```

![Similar Hue Palette](https://github.com/muesli/gamut/blob/master/docs/palette_similarhue.png)

Using the `ColorGenerator` interface, you can also write your own color generators.

#### Name A Color

```go
gamut.Wikipedia.Name(color) // => (name string, distance float64)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

#### Retrieving Colors

```go
gamut.Crayola.Filter("Red") // => []color.Color{...}
// returns a slice of all "Red" colors in the Crayola palette
gamut.Crayola.Colors() // => []color.Color{...}
// returns a slice of all colors in the Crayola palette
gamut.Monokai.Clamped(colors) // => []color.Color{...}
// returns a slice of the nearest matching colors in the Monokai palette
```

#### Mixing Palettes

```go
gamut.Crayola.MixedWith(gamut.Monokai) // => []gamut.Palette
// returns a palette with all colors from both Crayola and Monokai
```

## Themes

| Name    | Colors |
| ------- | ------:|
| Monokai |      7 |

### Examples

#### Roles

```go
gamut.MonokaiTheme.Role(gamut.Foreground) // => color.Color
```

Available roles are `Foreground`, `Background`, `Base`, `AlternateBase`, `Text`,
`Selection`, `Highlight`.

## Colors

### Examples

#### Lighter/Darker Colors

The `Darker` and `Lighter` functions darken and lighten respectively a given
color value by a specified percentage, without changing the color's hue.

```go
gamut.Darker(color, 0.1) // => color.Color
// returns a 10% darker version of color
gamut.Lighter(color, 0.3) // => color.Color
// returns a 30% lighter version of color
```

#### Complementary Colors

The `Complementary` function returns the complementary color for a given color.

```go
gamut.Complementary(color) // => color.Color
```

#### Contrast Colors

The `Contrast` function returns the color with the highest contrast to a given
color, either black or white.

```go
gamut.Contrast(color) // => color.Color
```

#### Warm/Cool Colors

```go
gamut.Warm(color) // => bool
gamut.Cool(color) // => bool
```

#### Around the Color Wheel

All the following functions return colors of a different hue, but with the same
lightness and saturation as the given colors.

```go
gamut.Triadic(color)            // => []color.Color{...}
gamut.Quadratic(color)          // => []color.Color{...}
gamut.Tetradic(color1, color2)  // => []color.Color{...}
gamut.Analogous(color)          // => []color.Color{...}
gamut.SplitComplementary(color) // => []color.Color{...}
```

#### Shades, Tints & Tones

```go
gamut.Monochromatic(color, 8) // => []color.Color{...}
// returns colors of the same hue, but with a different saturation/lightness
gamut.Shades(color, 8) // => []color.Color{...}
// returns a slice of 8 shades, from the given color to black
gamut.Tints(color, 8) // => []color.Color{...}
// returns a slice of 8 tints, from the given color to white
gamut.Tones(color, 8) // => []color.Color{...}
// returns a slice of 8 tones, from the given color to gray
```

![Shades, Tints and Tones](https://github.com/muesli/gamut/blob/master/docs/shades_tints_tones.png)

#### Blending Colors

```go
gamut.Blends(color1, color2, 8) // => []color.Color{...}
// returns a slice of interpolated colors by blending two colors
```

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
