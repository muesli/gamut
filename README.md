# gamut

[![Latest Release](https://img.shields.io/github/release/muesli/gamut.svg)](https://github.com/muesli/gamut/releases)
[![Build Status](https://github.com/muesli/gamut/workflows/build/badge.svg)](https://github.com/muesli/gamut/actions)
[![Coverage Status](https://coveralls.io/repos/github/muesli/gamut/badge.svg?branch=master)](https://coveralls.io/github/muesli/gamut?branch=master)
[![Go ReportCard](https://goreportcard.com/badge/muesli/gamut)](https://goreportcard.com/report/muesli/gamut)
[![GoDoc](https://godoc.org/github.com/golang/gddo?status.svg)](https://pkg.go.dev/github.com/muesli/gamut)

Go package to generate and manage color palettes & schemes

```go
import "github.com/muesli/gamut"
import "github.com/muesli/gamut/palette"
import "github.com/muesli/gamut/theme"
```

## Colors

gamut operates on various color spaces internally, but all color values you pass
in as parameters and all return values will match Go’s color.Color interface.

Let’s start with the basics. Just for convenience there’s a hex-value parser:

```go
color = gamut.Hex("#333")
color = gamut.Hex("#ABCDEF")
```

Both the short and standard formats are supported.

Conversely you can retrieve the hex encoding of any `color.Color` value:

```go
hex = gamut.ToHex(color)
```

### Around the Color Wheel

The `Darker` and `Lighter` functions darken and lighten respectively a given
color value by a specified percentage, without changing the color's hue:

```go
// returns a 10% darker version of color
color = gamut.Darker(color, 0.1)
// returns a 30% lighter version of color
color = gamut.Lighter(color, 0.3)
```

`Complementary` returns the complementary color for a given color:

```go
color = gamut.Complementary(color)
```

`Contrast` returns the color with the highest contrast to a given color, either
black or white:

```go
color = gamut.Contrast(color)
```

To retrieve a color with the same lightness and saturation, but a different
angle on the color wheel, you can use the HueOffset function:

```go
color = gamut.HueOffset(color, 90)
```

You can also go in the opposite direction by using negative values.

### Schemes

All the following functions return colors of a different hue, but with the same
lightness and saturation as the given colors:

Triadic schemes are made up of three hues equally spaced around the color wheel:

```go
colors = gamut.Triadic(color)
```

Quadratic schemes are made up of four hues equally spaced around the color wheel:

```go
colors = gamut.Quadratic(color)
```

Tetradic schemes are made up by two colors and their complementary values:

```go
colors = gamut.Tetradic(color1, color2)
```

Analogous schemes are created by using colors that are next to each other on the
color wheel:

```go
colors = gamut.Analogous(color)
```

SplitComplementary schemes are created by using colors next to the complementary
value of a given color:

```go
colors = gamut.SplitComplementary(color)
```

### Warm/Cool Colors

```go
ok = gamut.Warm(color)
ok = gamut.Cool(color)
```

### Shades, Tints & Tones

`Monochromatic` returns colors of the same hue, but with a different
saturation/lightness:

```go
colors = gamut.Monochromatic(color, 8)
```

![Monochromatic Palette](https://github.com/muesli/gamut/blob/master/docs/palette_monochromatic.png)

`Shades` returns colors blended from the given color to black:

```go
colors = gamut.Shades(color, 8)
```

![Shades Palette](https://github.com/muesli/gamut/blob/master/docs/palette_shades.png)

`Tints` returns colors blended from the given color to white:

```go
colors = gamut.Tints(color, 8)
```

![Tints Palette](https://github.com/muesli/gamut/blob/master/docs/palette_tints.png)

`Tones` returns colors blended from the given color to gray:

```go
colors = gamut.Tones(color, 8)
```

![Tones Palette](https://github.com/muesli/gamut/blob/master/docs/palette_tones.png)

### Blending Colors

`Blends` returns interpolated colors by blending two colors:

```go
colors = gamut.Blends(color1, color2, 8)
```

![Blends Palette](https://github.com/muesli/gamut/blob/master/docs/palette_blends.png)

## Palettes

Gamut comes with seven curated color palettes: Wikipedia, Crayola, CSS, RAL,
Resene, Monokai and Everforest. The Wikipedia palette is an import of common colors from
Wikipedia’s List of Colors. New curated palettes and importers are welcome. Send me
a pull request!

| Name       | Colors | Source                                                       |
| ---------- | -----: | ------------------------------------------------------------ |
| Wikipedia  |   1609 | https://en.wikipedia.org/wiki/List_of_colors_(compact)       |
| Crayola    |    180 | https://en.wikipedia.org/wiki/List_of_Crayola_crayon_colors  |
| CSS        |    147 | https://developer.mozilla.org/en-US/docs/Web/CSS/color_value |
| RAL        |    213 | https://en.wikipedia.org/wiki/List_of_RAL_colors             |
| Resene     |    759 | http://www.resene.co.nz                                      |
| Monokai    |     17 |                                                              |
| Everforest |    156 | https://github.com/sainnhe/everforest                        |

The function Colors lets you retrieve all colors in a palette:

```go
for _, c := range palette.Wikipedia.Colors() {
    fmt.Println(c.Name, c.Color)
}
```

This will print out a list of 1609 color names, as defined by Wikipedia.

### Creating Your Own Palettes

```go
var p gamut.Palette
p.AddColors(
    gamut.Colors{
        {"Name", gamut.Hex("#123456"), "Reference"},
        ...
    }
)
```

Name and Reference are optional when creating your own palettes.

### Names

Each color in the curated palettes comes with an “official” name. You can filter
palettes by colors with specific names. This code snippet will return a list of
all “blue” colors in the Wikipedia palette:

```go
colors = palette.Wikipedia.Filter("blue")
```

You can access a color with a specific name using the `Color` function:

```go
color, ok = palette.Wikipedia.Color("Pastel blue")
```

Calling a palette’s `Name` function with a given color returns the name & distance
of the closest (perceptually) matching color in it:

```go
name, distance = palette.Wikipedia.Name(color)
// name = "Baby blue"
// distance between 0.0 and 1.0
```

### Mixing Palettes

You can combine all colors of two palettes by mixing them:

```go
p = palette.Crayola.MixedWith(palette.Monokai)
```

### Perception

Sometimes you got a slice of colors, but you have a limited color palette to
work with. The Clamped function returns a slice of the closest perceptually
matching colors in a palette, maintaining the same order as the original slice
you provided. Finally you can remix your favorite wallpapers in Crayola-style!

```go
colors = palette.Crayola.Clamped(colors)
```

### Generating Color Palettes

Color Generators, like the provided `PastelGenerator`, `WarmGenerator` or
`HappyGenerator` can produce random (within the color space constraints of the
generator) color palettes:

```go
colors, err = gamut.Generate(8, gamut.PastelGenerator{})
```

![Pastel Palette](https://github.com/muesli/gamut/blob/master/docs/palette_pastel.png)

The `SimilarHueGenerator` produces colors with a hue similar to a given color:

```go
colors, err = gamut.Generate(8, gamut.SimilarHueGenerator{Color: gamut.Hex("#2F1B82")})
```

![Similar Hue Palette](https://github.com/muesli/gamut/blob/master/docs/palette_similarhue.png)

Using the `ColorGenerator` interface, you can also write your own color generators:

```go
type BrightGenerator struct {
	BroadGranularity
}

func (cc BrightGenerator) Valid(col colorful.Color) bool {
	_, _, l := col.Lab()
	return 0.7 <= l && l <= 1.0
}

...
colors, err := gamut.Generate(8, BrightGenerator{})
```

Only colors with a lightness between 0.7 and 1.0 will be accepted by this generator.

## Themes

| Name    | Colors |
| ------- | -----: |
| Monokai |      7 |

### Roles

```go
color = theme.MonokaiTheme.Role(theme.Foreground)
```

Available roles are `Foreground`, `Background`, `Base`, `AlternateBase`, `Text`,
`Selection`, `Highlight`.

## Feedback

Got some feedback or suggestions? Please open an issue or drop me a note!

* [Twitter](https://twitter.com/mueslix)
* [The Fediverse](https://mastodon.social/@fribbledom)
