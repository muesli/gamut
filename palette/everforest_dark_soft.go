package palette

import "github.com/muesli/gamut"

func init() {
	EverforestDarkSoft.AddColors(
		gamut.Colors{
			{"bg_dim", gamut.Hex("#293136"), "Background Dim"},
			{"bg0", gamut.Hex("#333C43"), "Background 0"},
			{"bg1", gamut.Hex("#3A464C"), "Background 1"},
			{"bg2", gamut.Hex("#434F55"), "Background 2"},
			{"bg3", gamut.Hex("#4D5960"), "Background 3"},
			{"bg4", gamut.Hex("#555F66"), "Background 4"},
			{"bg5", gamut.Hex("#5D6B66"), "Background 5"},
			{"bg_visual", gamut.Hex("#5C3F4F"), "Background Visual"},
			{"bg_red", gamut.Hex("#59464C"), "Background Red"},
			{"bg_green", gamut.Hex("#48584E"), "Background Green"},
			{"bg_blue", gamut.Hex("#3F5865"), "Background Blue"},
			{"bg_yellow", gamut.Hex("#55544A"), "Background Yellow"},
			{"fg", gamut.Hex("#D3C6AA"), "Foreground"},
			{"red", gamut.Hex("#E67E80"), "Red"},
			{"orange", gamut.Hex("#E69875"), "Orange"},
			{"yellow", gamut.Hex("#DBBC7F"), "Yellow"},
			{"green", gamut.Hex("#A7C080"), "Green"},
			{"aqua", gamut.Hex("#83C092"), "Aqua"},
			{"blue", gamut.Hex("#7FBBB3"), "Blue"},
			{"purple", gamut.Hex("#D699B6"), "Purple"},
			{"grey0", gamut.Hex("#7A8478"), "Grey 0"},
			{"grey1", gamut.Hex("#859289"), "Grey 1"},
			{"grey2", gamut.Hex("#9DA9A0"), "Grey 2"},
			{"statusline1", gamut.Hex("#A7C080"), "Statusline 1"},
			{"statusline2", gamut.Hex("#D3C6AA"), "Statusline 2"},
			{"statusline3", gamut.Hex("#E67E80"), "Statusline 3"},
		})
}
