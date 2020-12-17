package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/golang/go_mark/bookmark"
	"github.com/urfave/cli/v2"
)

func searchAction(context *cli.Context) error {

	file := storageFilePath(context)

	bmHolder := bookmark.LoadBookmarks(file)

	if context.Bool("all") {
		js, err := str.FormatJson(bmHolder.All())
		if err != nil {
			return err
		}

		fmt.Println(js)

		return nil
	}

	return nil
}

func storageFilePath(context *cli.Context) string {
	storageFile := context.String("storage")
	return storageFile
}

func NewSearchCommand() *cli.Command {
	configDir, err := os.UserConfigDir()
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		log.Fatal("variavel não inicializada")
	}

	configFile := configDir + "/bmark/conf"
	bookmarkFile := cacheDir + "/bmark/bookmarks"

	return &cli.Command{
		Name:  "search",
		Usage: "busca um bookmark pela palavra chave",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:  "all",
				Usage: "traz todos os bookmarks",
				Aliases: []string{
					"a",
				},
			},
			&cli.StringFlag{
				Name:    "config",
				Usage:   "roda o comando apontando para um arquivo de configuração",
				Aliases: []string{"c"},
				Value:   configFile,
			},
			&cli.StringFlag{
				Name:  "storage",
				Usage: "arquivo de onde os bookmarks devem ser lidos",
				Value: bookmarkFile,
			},
		},

		Action: searchAction,
	}

}
