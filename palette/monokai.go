package palette

import "github.com/muesli/gamut"

func init() {
	Monokai.AddColors(
		gamut.Colors{
			{"Spray", gamut.Hex("#66D9EF"), ""},           // Type
			{"Tree Poppy", gamut.Hex("#FD971F"), ""},      // Function Argument
			{"Armadillo", gamut.Hex("#49483E"), ""},       // Selection
			{"El Paso", gamut.Hex("#3E3D32"), ""},         // Highlight
			{"Center Stage", gamut.Hex("#A6E22E"), ""},    // Function, Class
			{"Feverish Pink", gamut.Hex("#F92672"), ""},   // Keyword
			{"Lavish Lavender", gamut.Hex("#AE81FF"), ""}, // Number, constant
			{"Funky Yellow", gamut.Hex("#E6DB74"), ""},    // String
			{"Cocoon", gamut.Hex("#75715E"), ""},          // Comment
			{"Extra White", gamut.Hex("#F8F8F2"), ""},     // Text
			{"Caviar", gamut.Hex("#272822"), ""},          // Background
			{"Caviar Dark", gamut.Hex("#141411"), ""},
			{"Blue Beyond", gamut.Hex("#89BDFF"), ""},
			{"Urbane Bronze", gamut.Hex("#595959"), ""},
			{"Tricorn Black", gamut.Hex("#383830"), ""},
			{"Soothing White", gamut.Hex("#E6E6E6"), ""},
			{"Ice Plant", gamut.Hex("#FD5FF1"), ""},
		})
}
