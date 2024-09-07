package palette

import "github.com/muesli/gamut"

func init() {
	EverforestLightSoft.AddColors(
		gamut.Colors{
			{"bg_dim", gamut.Hex("#FAF4E6"), "Background Dim"},
			{"bg0", gamut.Hex("#FFFAE8"), "Background 0"},
			{"bg1", gamut.Hex("#F9F5E5"), "Background 1"},
			{"bg2", gamut.Hex("#F4F1E0"), "Background 2"},
			{"bg3", gamut.Hex("#F0EDDB"), "Background 3"},
			{"bg4", gamut.Hex("#ECE8D6"), "Background 4"},
			{"bg5", gamut.Hex("#D8DEBF"), "Background 5"},
			{"bg_visual", gamut.Hex("#F0F2D4"), "Background Visual"},
			{"bg_red", gamut.Hex("#FFE3DD"), "Background Red"},
			{"bg_green", gamut.Hex("#F0F3D6"), "Background Green"},
			{"bg_blue", gamut.Hex("#E9F3EC"), "Background Blue"},
			{"bg_yellow", gamut.Hex("#FDF0D4"), "Background Yellow"},
			{"fg", gamut.Hex("#6C707A"), "Foreground"},
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
			{"statusline2", gamut.Hex("#6C707A"), "Statusline 2"},
			{"statusline3", gamut.Hex("#F85552"), "Statusline 3"},
		})
}
