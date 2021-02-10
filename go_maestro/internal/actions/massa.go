package actions

/*func NewGoMassaCustomInput() action.Action {
	input := &input.ResolverConfig{
		Config: input.Config{
			Keys:   []string{"-c"},
			ArgSep: " ",
		},
		Resolver: input.ClipResolver{},
	}

	c := action.Action{
		Binary: action.Binary{
			CmdPath: "go_massa",
			Name:    "Massa X",
			Desc:    "Obtém massa Cnis com catálogo customizado",
			FixArgs: []string{"get", "-a", "2"},
		},
		InputConfig: input,
	}

	return c
}*/

/*func MassaListaCatalogos() action.Action {

	massa := action.Action{
		Binary:     "go_massa",
		Name:       "Massa-Sibe: Lista catálogos",
		Desc:       "Lista de Catálogos (CNIS HOM)",
		Args:       []string{"list", "-a", "2"},
		Config:  nil,
		CopyOutput: false,
		FormatOutput: []string{
			"#.id", "#.nome",
		},
		OutputConverter: OutputConverterListaCatalogos,
		CallNext:        NewMassaCnisFromCustomCat,
	}

	return massa
}

func OutputConverterListaCatalogos(choosen string) (string, []string) {
	catId := strings.Split(choosen, "\t")[0]
	return "", []string{"-c", catId}
}

func NewMassaCnisFromCustomCat(args ...string) *action.Action {

	return &action.Action{
		Binary:       "go_massa",
		Name:         "Nova Massa Cnis #8",
		Desc:         "Obtém uma nova massa do CNIS Homologação Catálogo 8",
		Args:         append([]string{"get", "-a", "2"}, args...),
		Config:    nil,
		CopyOutput:   true,
		FilterOutput: []string{"cpfMassa cpf", "nomePfMassa nomeTitular"},
		CallNext:     nil,
	}

}

func NewMassaCnisHomCat8() action.Action {
	return action.Action{
		Binary:       "go_massa",
		Name:         "Nova Massa Cnis #8",
		Desc:         "Obtém uma nova massa do CNIS Homologação Catálogo 8",
		Args:         []string{"get", "-c", "8", "-a", "2"},
		Config:    nil,
		CopyOutput:   true,
		FilterOutput: []string{"cpfMassa cpf", "nomePfMassa nomeTitular"},
		CallNext:     nil,
	}

}*/
