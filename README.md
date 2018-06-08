# gamut

Go package to generate and manage color palettes & schemes

```go
import "github.com/muesli/gamut"
import "github.com/muesli/gamut/palette"
import "github.com/muesli/gamut/theme"
```

## Colors

#### Around the Color Wheel

The `Darker` and `Lighter` functions darken and lighten respectively a given
color value by a specified percentage, without changing the color's hue:

```go
gamut.Darker(color, 0.1) // => color.Color
// returns a 10% darker version of color
gamut.Lighter(color, 0.3) // => color.Color
// returns a 30% lighter version of color
```

`Complementary` returns the complementary color for a given color:

```go
gamut.Complementary(color) // => color.Color
```

`Contrast` returns the color with the highest contrast to a given color, either
black or white:

```go
gamut.Contrast(color) // => color.Color
```

All the following functions return colors of a different hue, but with the same
lightness and saturation as the given colors:

```go
gamut.Triadic(color)            // => []color.Color{...}
gamut.Quadratic(color)          // => []color.Color{...}
gamut.Tetradic(color1, color2)  // => []color.Color{...}
gamut.Analogous(color)          // => []color.Color{...}
gamut.SplitComplementary(color) // => []color.Color{...}
```

#### Warm/Cool Colors

```go
gamut.Warm(color) // => bool
gamut.Cool(color) // => bool
```

#### Shades, Tints & Tones

`Monochromatic` returns colors of the same hue, but with a different
saturation/lightness:

```go
gamut.Monochromatic(color, 8) // => []color.Color{...}
```

![Monochromatic Palette](https://github.com/muesli/gamut/blob/master/docs/palette_monochromatic.png)

`Shades` returns colors blended from the given color to black:

```go
gamut.Shades(color, 8) // => []color.Color{...}
```

![Shades Palette](https://github.com/muesli/gamut/blob/master/docs/palette_shades.png)

`Tints` returns colors blended from the given color to white:

```go
gamut.Tints(color, 8) // => []color.Color{...}
```

![Tints Palette](https://github.com/muesli/gamut/blob/master/docs/palette_tints.png)

`Tones` returns colors blended from the given color to gray:

```go
gamut.Tones(color, 8) // => []color.Color{...}
```

![Tones Palette](https://github.com/muesli/gamut/blob/master/docs/palette_tones.png)

#### Blending Colors

`Blends` returns interpolated colors by blending two colors:

```go
gamut.Blends(color1, color2, 8) // => []color.Color{...}
```

![Blends Palette](https://github.com/muesli/gamut/blob/master/docs/palette_blends.png)

## Palettes

| Name      | Colors | Source                                                      |
| --------- | ------:| ----------------------------------------------------------- |
| Wikipedia |   1609 | https://en.wikipedia.org/wiki/List_of_colors_(compact)      |
| Crayola   |    180 | https://en.wikipedia.org/wiki/List_of_Crayola_crayon_colors |
| Resene    |    759 | http://www.resene.co.nz                                     |
| Monokai   |     17 |                                                             |

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
palette.Wikipedia.Name(color) // => (name string, distance float64)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

#### Retrieving Colors

```go
palette.Crayola.Filter("Red") // => []color.Color{...}
// returns a slice of all "Red" colors in the Crayola palette
palette.Resene.Colors() // => []color.Color{...}
// returns a slice of all colors in the Resene palette
palette.Monokai.Clamped(colors) // => []color.Color{...}
// returns a slice of the nearest matching colors in the Monokai palette
```

#### Mixing Palettes

You can combine all colors of two palettes by mixing them:

```go
palette.Crayola.MixedWith(palette.Monokai) // => gamut.Palette
```

## Themes

| Name    | Colors |
| ------- | ------:|
| Monokai |      7 |

#### Roles

```go
theme.MonokaiTheme.Role(theme.Foreground) // => color.Color
```

Available roles are `Foreground`, `Background`, `Base`, `AlternateBase`, `Text`,
`Selection`, `Highlight`.

## Development

[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://godoc.org/github.com/muesli/gamut)
[![Build Status](https://travis-ci.org/muesli/gamut.svg?branch=master)](https://travis-ci.org/muesli/gamut)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](http://goreportcard.com/badge/muesli/gamut)](http://goreportcard.com/report/muesli/gamut)
