package commands

func MassaListaCatalogos() Cmd {
	massa := Cmd{
		Binary:    "go_massa",
		Name:      "Massa-Sibe: Lista catálogos",
		Desc:      "Lista de Catálogos (CNIS HOM)",
		Args: []string{"list", "-a", "2"},
		UserInput: map[string]string{},
		Clipboard: false,
		FilterOutput: []string{
			"#.id", "#.nome",
		},

	}

	return massa
}

type NameCpf struct {


}
//FilterOutput: []string{".cpfMassa as cpf", ".nomePfMassa as nomeTitular"},
func NewMassaCnisHomCat8() Cmd{
	return Cmd{
		Binary:       "go_massa",
		Name:         "Nova Massa Cnis #8",
		Desc:         "Obtém uma nova massa do CNIS Homologação Catálogo 8",
		Args: []string{"get", "-c", "8", "-a", "2"},
		UserInput:    nil,
		Clipboard:    true,
		FilterOutput: []string{".cpfMassa @cpf", ".nomePfMassa @nomeTitular"},
			Next:         nil,
	}

}