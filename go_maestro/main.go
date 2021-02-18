package main

import (
	"fmt"
	"github.com/alexpfx/go_action/action"

	"github.com/alexpfx/golang/go_maestro/internal/trees"
	"github.com/urfave/cli/v2"
	"os"
)

func main() {
	app := &cli.App{
		Name: "go maestro",
		Commands: []*cli.Command{
			{
				Name:      "tree",
				Aliases:   []string{},
				Usage:     "monta uma árvore de menu de nome passado como parâmetro. ",
				ArgsUsage: "<tree_name>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "menu",
						Aliases: []string{"m"},
						Usage:   "especifica o comando usado para montar o menu [fzf|rofi]",
						Value:   "fzf",
					},
					&cli.BoolFlag{
						Name:    "list",
						Usage:   "lista as árvores disponíveis e encerra a execução do comando",
						Aliases: []string{"l"},
					},
				},
				Action: func(ctx *cli.Context) error {
					if ctx.Bool("list") {
						printTrees()
						return nil
					}
					if ctx.NArg() < 1 {
						fmt.Println(ctx.Command.Usage)
						return nil
					}

					arvoreBuscada := ctx.Args().First()

					vtree, foundTree := searchTree(arvoreBuscada)

					if !foundTree {
						fmt.Println("Não encontrada: " + arvoreBuscada)
						return nil
					}


					selectedAction, found, err := vtree.Show()

					if err != nil {
						return err
					}
					if !found {
						return fmt.Errorf("selecione uma action")
					}

					//err = log.Output(1, fmt.Sprintf("executando action: %s", selectedAction.Name))
					if err != nil {
						return err
					}
					actRes, err := selectedAction.Run()
					if err != nil {
						return err
					}

					fmt.Println(string(actRes))

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s!\n", err.Error())
		os.Exit(1)
	}

}

func printTrees() {

	allTrees()

}

func allTrees() {

	for _, tree := range trees.BuiltInTrees {
		fmt.Println(tree.Name())

	}
}

func searchTree(treeName string) (action.Tree, bool) {
	builtInTrees := trees.BuiltInTrees

	for _, tree := range builtInTrees {
		if tree.Name() == treeName {
			return tree, true
		}
	}
	return nil, false

}
