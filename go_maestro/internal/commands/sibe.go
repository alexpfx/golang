package commands

func SibeSibeDeploy() Cmd {

	massa := Cmd{
		Binary:       "go_sibe",
		Name:         "sibeDeploy",
		Desc:         "Lista de Opções de Deploy",
		Args:         []string{"deploys"},
		FilterOutput: []string{"id", "name"},
		FormatOutput: []string{"#.id", "#.name"},
	}

	return massa
}
func SibeSibeClient() Cmd {

	massa := Cmd{
		Binary:       "go_sibe",
		Name:         "sibeClient",
		Desc:         "Clients",
		Args:         []string{"clients"},
		FilterOutput: []string{"id", "name"},
		FormatOutput: []string{"#.id", "#.name"},
	}

	return massa
}
