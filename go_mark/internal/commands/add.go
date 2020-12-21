package commands

import (
	"fmt"
	"github.com/urfave/cli/v2"
)

func addAction(context *cli.Context) error {

	if context.NArg() < 1 {
		return fmt.Errorf("a URL do bookmark é obrigatória")
	}
	url := context.Args().First()
	err := checkUrl(&url)
	if err != nil {
		return fmt.Errorf("a URL passada é inválida: %s", url)
	}

	context.String("desc")
	context.String("category")
	context.StringSlice("tag")
	context.Bool("auto")

	return nil
}

func checkUrl(url *string) error{
	return nil
}

func NewAddCommand() *cli.Command {
	
	return &cli.Command{
		Action: addAction,
		Usage: "adiciona um bookmark",
		Name:   "add",
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
	}

}
