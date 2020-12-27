package main

import (
	"fmt"
	"github.com/alexpfx/go_common/slices"
	"github.com/alexpfx/golang/go_merge/internal/merge"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

const baseUrl = "https://www-scm.prevnet/api/v3/projects"
const sibeProject = "754"

func main() {
	log.Println("iniciando go_merge")
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "info",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:  "author",
						Usage: "filtro por autor",
					},
					&cli.StringFlag{
						Name:  "targetBranch",
						Usage: "filtro por target branch",
					},
					&cli.StringFlag{
						Name:    "token",
						Usage:   "gitlab token",
						EnvVars: []string{"PRIVATE_TOKEN"},
					},
				},

				Action: func(c *cli.Context) error {
					if c.NArg() < 1 {
						return fmt.Errorf("esperava o id ou a URL do merge request")
					}
					ids, err := merge.ParseIds(c.Args().Slice())
					if err != nil {
						return err
					}

					ids = slices.IntUniqueSorted(ids)
					checkIdsCount(ids)

					author := c.String("author")
					targetBranch := c.String("targetBranch")
					token := c.String("token")

					filter := mapFilter(author, targetBranch)
					mrInfo, er, err := merge.Fetch(token, baseUrl, sibeProject, ids, filter)
					if err != nil {
						fmt.Println(err)
					}
					if len(er) > 0 && len(mrInfo) == 0 {
						res := make([]string, 0)
						for _, result := range er {
							res = append(res, fmt.Sprintf("%d : %s", result.MergeId, result.Err))
						}
						return fmt.Errorf("nenhum MR foi retornado. \ninválidos: %v", res)
					}

					formatedOutput := merge.FormatOutput(mrInfo, merge.FormatJson)
					fmt.Println(formatedOutput)

					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatalf("erro na chamada a %v: \n%v", os.Args[0], err.Error())

	}
}

func checkIdsCount(ids []int) {
	if len(ids) > 50 {
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
