package main

import (
	"fmt"
	"github.com/alexpfx/golang/go_merge/internal/util"

	"github.com/alexpfx/golang/go_merge/internal/merge"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)
const version = "0.0.1+2"
const baseUrl = "https://www-scm.prevnet/api/v3/projects"
const sibeProject = "754"

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "version",
				Action: func(c *cli.Context) error {
					print(version)
				return nil
				},
			},
			{
				Name: "info",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "author",
						Usage: "filtro por autor",
						Aliases: []string{"a"},
					},
					&cli.StringFlag{
						Name:  "targetBranch",
						Usage: "filtro por target branch",
						Aliases: []string{"b"},

					},
					&cli.StringFlag{
						Name:    "token",
						Usage:   "gitlab token",
						EnvVars: []string{"PRIVATE_TOKEN"},
					},
				},
				ArgsUsage: "mergeId ou mergeIdInicial",

				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("esperava o id ou a URL do merge request")
					}
					ids, err := merge.ParseIds(c.Args().Slice())
					if err != nil {
						return err
					}

					ids = util.IntUniqueSorted(ids)
					checkIdsCount(ids)

					author := c.String("author")
					targetBranch := c.String("targetBranch")
					token := c.String("token")
					if token == "" {
						return fmt.Errorf("faltando variável de ambiente: PRIVATE_TOKEN")
					}

					filter := mapFilter(author, targetBranch)
					mrInfo, er, err := merge.Fetch(token, baseUrl, sibeProject, ids, filter)
					if err != nil {
						fmt.Println(err)
					}
					if len(er) > 0 {
						fmt.Println("Invalidos: ")
						for _, result := range er {
							fmt.Printf("%d: %v", result.MergeId, result.Err)
						}
					}

					formatedOutput := merge.FormatOutput(mrInfo, merge.FormatDev)
					fmt.Println(formatedOutput)

					return nil
				},
			},
		},
	}



	err := app.Run(os.Args)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func checkIdsCount(ids []int) {
	if len(ids) > 120 {
		err := fmt.Errorf("máximo de %v ids por chamada", 50)
		log.Fatal(err)
	}
}

func mapFilter(author string, branch string) map[string]string {
	result := make(map[string]string)

	if author != "" {
		result["author"] = author
	}

	if branch != "" {
		result["target_branch"] = branch
	}

	return result
}
