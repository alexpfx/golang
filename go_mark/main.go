package main

import (
	"log"
	"os"

	"github.com/alexpfx/golang/go_mark/internal/commands"
	"github.com/urfave/cli/v2"
)
/* */
func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			commands.NewAddCommand(),
			commands.NewSearchCommand(),
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
