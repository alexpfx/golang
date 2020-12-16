package commands

import "github.com/urfave/cli/v2"

func action (context *cli.Context) error{



	return nil
}

func NewAddCommand() *cli.Command {

	return &cli.Command{
		Action: action,
		Name: "add",
		ArgsUsage: `add http://www.duckduckgo.com` +
			`[--desc 'site de pesquisa baseado em privacidade']` +
			`[--category buscadores] [--tag buscador][--tag pesquisa] [--auto]`,

		Flags: []cli.Flag{
			&cli.StringFlag{

				Name:    "category",
				Usage:   "categoria da url",
				Value:   "",
				Aliases: []string{"c"},
			},
			&cli.StringFlag{
				Name:    "desc",
				Usage:   "descrição do bookmark",
				Value:   "",
				Aliases: []string{"d"},
			},
			&cli.StringSliceFlag{
				Name:    "tag",
				Usage:   "permite atribuir uma ou mais tags ao bookmark",
				Aliases: []string{"t"},
			},
			&cli.BoolFlag{
				Name: "auto",
				Usage: "se esta flag for informada, o programa irá acessar a url do bookmark para " +
					" tentar obter as informações sobre o bookimark que não foram passadas via " +
					"linha de comando",
			},
		},
		Usage: "adiciona um bookmark",
	}

}
