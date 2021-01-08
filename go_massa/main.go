package main

import (
	"fmt"
	"log"
	"os"
	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/golang/go_massa/internal/massa"
	"github.com/urfave/cli/v2"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:  "list",
				Usage: "lista todos os catálogos com massa disponível",
				Flags: []cli.Flag{
					&cli.IntFlag{
						Name:    "ambiente",
						Aliases: []string{"a"},
						Usage:   "listar catálogos do ambiente",
						Value:   2,
					},
				},
				Action: func(context *cli.Context) error {
					list := massa.NewCatalogoList()
					catalogos, err := list.Catalogos(context.Int("ambiente"))
					if err != nil {
						return err
					}
					jsonStr, err := str.FormatJson(catalogos)
					if (err != nil){
						log.Fatal(err.Error())
					}
					fmt.Println(jsonStr)
					return nil
				},
			},
			{
				Name:  "get",
				Usage: "obtém uma massa do catálogo e ambiente",
				Action: func(context *cli.Context) error {
					getter := massa.NewRetriever()

					var novaMassa massa.Massa
					var err error
					if context.Bool("oldest"){
						novaMassa, err = getter.Oldest(context.Int("catalogo"), context.Int("ambiente"))
					}else{
						novaMassa, err = getter.Newest(context.Int("catalogo"), context.Int("ambiente"))
					}

					if err != nil {
						return err
					}
					fmt.Println(str.FormatJson(novaMassa))
					return nil

				},
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:    "oldest",
						Aliases: []string{"o"},
						Usage:   "obtém a massa mais antigo, em vez da mais recente",
						Value:   false,
					},
					&cli.IntFlag{
						Name:    "ambiente",
						Aliases: []string{"a"},
						Usage:   "ambiente",
						Value:   2,
					},
					&cli.IntFlag{
						Name:    "catalogo",
						Aliases: []string{"c"},
						Usage:   "catalogo",
						Value:   8,
					},
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
