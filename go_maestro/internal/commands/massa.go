package commands

func MassaListaCatalogos() Cmd {
	massa := Cmd{
		Binary:    "go_massa",
		Name:      "Massa-Sibe: Lista catálogos",
		Desc:      "Lista de Catálogos (CNIS HOM)",
		Args: []string{"list", "-a", "2"},
		UserInput: map[string]string{},
		Clipboard: true,
	}

	return massa
}