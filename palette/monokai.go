package palette

import "github.com/muesli/gamut"

func init() {
	Monokai.AddColors(
		gamut.Colors{
			{"spray", gamut.Hex("#66d9ef"), ""},           // type
			{"tree poppy", gamut.Hex("#fd971f"), ""},      // function argument
			{"armadillo", gamut.Hex("#49483e"), ""},       // selection
			{"el paso", gamut.Hex("#3e3d32"), ""},         // highlight
			{"center stage", gamut.Hex("#a6e22e"), ""},    // function, class
			{"feverish pink", gamut.Hex("#f92672"), ""},   // keyword
			{"lavish lavender", gamut.Hex("#ae81ff"), ""}, // number, constant
			{"funky yellow", gamut.Hex("#e6db74"), ""},    // string
			{"cocoon", gamut.Hex("#75715e"), ""},          // comment
			{"extra white", gamut.Hex("#f8f8f2"), ""},     // text
			{"caviar", gamut.Hex("#272822"), ""},          // background
			{"caviar dark", gamut.Hex("#141411"), ""},
			{"blue beyond", gamut.Hex("#89bdff"), ""},
			{"urbane bronze", gamut.Hex("#595959"), ""},
			{"tricorn black", gamut.Hex("#383830"), ""},
			{"soothing white", gamut.Hex("#e6e6e6"), ""},
			{"ice plant", gamut.Hex("#fd5ff1"), ""},
		})
}
