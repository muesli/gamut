package palette

import "github.com/muesli/gamut"

func init() {
	EverforestDarkMedium.AddColors(
		gamut.Colors{
			{"bg_dim", gamut.Hex("#232A2E"), "Background Dim"},
			{"bg0", gamut.Hex("#2D353B"), "Background 0"},
			{"bg1", gamut.Hex("#343F44"), "Background 1"},
			{"bg2", gamut.Hex("#3D484D"), "Background 2"},
			{"bg3", gamut.Hex("#475258"), "Background 3"},
			{"bg4", gamut.Hex("#4F585E"), "Background 4"},
			{"bg5", gamut.Hex("#56635f"), "Background 5"},
			{"bg_visual", gamut.Hex("#543A48"), "Background Visual"},
			{"bg_red", gamut.Hex("#514045"), "Background Red"},
			{"bg_green", gamut.Hex("#425047"), "Background Green"},
			{"bg_blue", gamut.Hex("#3A515D"), "Background Blue"},
			{"bg_yellow", gamut.Hex("#4D4C43"), "Background Yellow"},
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
