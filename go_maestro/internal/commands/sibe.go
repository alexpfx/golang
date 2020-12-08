package commands


func SibeListaDeploys() Cmd {

	massa := Cmd{
		Binary:     "go_sibe",
		Name:       "deploys",
		Desc:       "Lista de Opções de Deploy",
		Args:       []string{"list", "-deploys"},
	}

	return massa
}
