package palette

import "github.com/muesli/gamut"

func init() {
	Monokai.AddColors(
		gamut.Colors{
			{Name: "Spray", Color: gamut.Hex("#66D9EF"), Reference: ""},           // Type
			{Name: "Tree Poppy", Color: gamut.Hex("#FD971F"), Reference: ""},      // Function Argument
			{Name: "Armadillo", Color: gamut.Hex("#49483E"), Reference: ""},       // Selection
			{Name: "El Paso", Color: gamut.Hex("#3E3D32"), Reference: ""},         // Highlight
			{Name: "Center Stage", Color: gamut.Hex("#A6E22E"), Reference: ""},    // Function, Class
			{Name: "Feverish Pink", Color: gamut.Hex("#F92672"), Reference: ""},   // Keyword
			{Name: "Lavish Lavender", Color: gamut.Hex("#AE81FF"), Reference: ""}, // Number, constant
			{Name: "Funky Yellow", Color: gamut.Hex("#E6DB74"), Reference: ""},    // String
			{Name: "Cocoon", Color: gamut.Hex("#75715E"), Reference: ""},          // Comment
			{Name: "Extra White", Color: gamut.Hex("#F8F8F2"), Reference: ""},     // Text
			{Name: "Caviar", Color: gamut.Hex("#272822"), Reference: ""},          // Background
			{Name: "Caviar Dark", Color: gamut.Hex("#141411"), Reference: ""},
			{Name: "Blue Beyond", Color: gamut.Hex("#89BDFF"), Reference: ""},
			{Name: "Urbane Bronze", Color: gamut.Hex("#595959"), Reference: ""},
			{Name: "Tricorn Black", Color: gamut.Hex("#383830"), Reference: ""},
			{Name: "Soothing White", Color: gamut.Hex("#E6E6E6"), Reference: ""},
			{Name: "Ice Plant", Color: gamut.Hex("#FD5FF1"), Reference: ""},
		})
}
