package gamut

import (
	"image/color"

	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/muesli/kmeans"
)

type ColorChecker interface {
	Valid(col colorful.Color) bool
}

type Warm struct {
}
type Happy struct {
}
type Pastel struct {
}

func (cc Warm) Valid(col colorful.Color) bool {
	_, c, l := col.Hcl()
	return 0.1 <= c && c <= 0.4 && 0.2 <= l && l <= 0.5
}
func (cc Happy) Valid(col colorful.Color) bool {
	_, c, l := col.Hcl()
	return 0.3 <= c && 0.4 <= l && l <= 0.8
}
func (cc Pastel) Valid(col colorful.Color) bool {
	_, s, v := col.Hsv()
	return 0.2 <= s && s <= 0.4 && 0.7 <= v && v <= 1.0
}

func Generate(count int, checker ColorChecker) ([]color.Color, error) {
	// Create data points in the CIE L*a*b color space
	// l for lightness channel
	// a, b for color channels
	var cc []color.Color
	dl := 0.05
	dab := 0.1

	var d kmeans.Points
	for l := 0.0; l <= 1.0; l += dl {
		for a := 0.0; a < 1.0; a += dab {
			for b := 0.0; b < 1.0; b += dab {
				col := colorful.Lab(l, -1.0+(a*2.0), -1.0+(b*2.0))
				// col = colorful.Hcl(a*360.0, b, c)

				if !col.IsValid() || !checker.Valid(col) {
					continue
				}

				d = append(d, kmeans.Point{l, a, b})
			}
		}
	}

	// Enable graph generation (.png files) for each iteration
	// km, _ := kmeans.NewWithOptions(0.02, Plotter{})
	km, err := kmeans.NewWithOptions(0.02, nil)
	if err != nil {
		return cc, err
	}

	// Partition the color space into multiple clusters
	clusters, err := km.Partition(d, count)
	if err != nil {
		return cc, err
	}

	for _, c := range clusters {
		col := colorful.Lab(c.Center[0], -1.0+(c.Center[1]*2.0), -1.0+(c.Center[2]*2.0)).Clamped()
		cc = append(cc, col)
	}

	return cc, nil
}
