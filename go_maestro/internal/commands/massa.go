package commands

import (
	"github.com/alexpfx/go_common/cmd"
	"github.com/alexpfx/go_common/user"
	"strings"
)

func NewGoMassaCustomInput() cmd.Cmd {
	input := &cmd.Input{
		InputList: user.MultiInput{
			Keys:   []string{"-c"},
			ArgSep: " ",
		},
		Reader: user.ClipInputReader{},
	}

	c := cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "go_massa",
			Name:    "Massa X",
			Desc:    "Obtém massa Cnis com catálogo customizado",
			FixArgs: []string{"get", "-a", "2"},
		},
		UserInput: input,
	}

	return c
}

func MassaListaCatalogos() Cmd {

	massa := Cmd{
		Binary:     "go_massa",
		Name:       "Massa-Sibe: Lista catálogos",
		Desc:       "Lista de Catálogos (CNIS HOM)",
		Args:       []string{"list", "-a", "2"},
		UserInput:  nil,
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

func NewMassaCnisFromCustomCat(args ...string) *Cmd {

	return &Cmd{
		Binary:       "go_massa",
		Name:         "Nova Massa Cnis #8",
		Desc:         "Obtém uma nova massa do CNIS Homologação Catálogo 8",
		Args:         append([]string{"get", "-a", "2"}, args...),
		UserInput:    nil,
		CopyOutput:   true,
		FilterOutput: []string{"cpfMassa cpf", "nomePfMassa nomeTitular"},
		CallNext:     nil,
	}

}

func NewMassaCnisHomCat8() Cmd {
	return Cmd{
		Binary:       "go_massa",
		Name:         "Nova Massa Cnis #8",
		Desc:         "Obtém uma nova massa do CNIS Homologação Catálogo 8",
		Args:         []string{"get", "-c", "8", "-a", "2"},
		UserInput:    nil,
		CopyOutput:   true,
		FilterOutput: []string{"cpfMassa cpf", "nomePfMassa nomeTitular"},
		CallNext:     nil,
	}

}
