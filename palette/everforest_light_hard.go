package palette

import "github.com/muesli/gamut"

func init() {
	EverforestLightHard.AddColors(
		gamut.Colors{
			{"bg_dim", gamut.Hex("#F2EFDF"), "Background Dim"},
			{"bg0", gamut.Hex("#FFFBEF"), "Background 0"},
			{"bg1", gamut.Hex("#F8F5E4"), "Background 1"},
			{"bg2", gamut.Hex("#F2EFDF"), "Background 2"},
			{"bg3", gamut.Hex("#EDEADA"), "Background 3"},
			{"bg4", gamut.Hex("#E8E5D5"), "Background 4"},
			{"bg5", gamut.Hex("#BEC5B2"), "Background 5"},
			{"bg_visual", gamut.Hex("#F0F2D4"), "Background Visual"},
			{"bg_red", gamut.Hex("#FFE7DE"), "Background Red"},
			{"bg_green", gamut.Hex("#F3F5D9"), "Background Green"},
			{"bg_blue", gamut.Hex("#ECF5ED"), "Background Blue"},
			{"bg_yellow", gamut.Hex("#FEF2D5"), "Background Yellow"},
			{"fg", gamut.Hex("#5C6A72"), "Foreground"},
			{"red", gamut.Hex("#F85552"), "Red"},
			{"orange", gamut.Hex("#F57D26"), "Orange"},
			{"yellow", gamut.Hex("#DFA000"), "Yellow"},
			{"green", gamut.Hex("#8DA101"), "Green"},
			{"aqua", gamut.Hex("#35A77C"), "Aqua"},
			{"blue", gamut.Hex("#3A94C5"), "Blue"},
			{"purple", gamut.Hex("#DF69BA"), "Purple"},
			{"grey0", gamut.Hex("#A6B0A0"), "Grey 0"},
			{"grey1", gamut.Hex("#919B84"), "Grey 1"},
			{"grey2", gamut.Hex("#7D8778"), "Grey 2"},
			{"statusline1", gamut.Hex("#8DA101"), "Statusline 1"},
			{"statusline2", gamut.Hex("#5C6A72"), "Statusline 2"},
			{"statusline3", gamut.Hex("#F85552"), "Statusline 3"},
		})
}
