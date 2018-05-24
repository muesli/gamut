package gamut

type Theme struct {
	Name   string
	Colors map[Role]Color
}

var (
	MonokaiTheme = Theme{
		Name:   "monokai",
		Colors: make(map[Role]Color),
	}
)

func init() {
	MonokaiTheme.Colors[Foreground] = Monokai.Filter("Extra White")[0]
	MonokaiTheme.Colors[Background] = Monokai.Filter("Caviar")[0]
	MonokaiTheme.Colors[Base] = Monokai.Filter("Caviar")[0]
	MonokaiTheme.Colors[AlternateBase] = Monokai.Filter("Caviar Dark")[0]
	MonokaiTheme.Colors[Text] = Monokai.Filter("Cocoon")[0]
	MonokaiTheme.Colors[Selection] = Monokai.Filter("Armadillo")[0]
	MonokaiTheme.Colors[Highlight] = Monokai.Filter("El Paso")[0]
}
