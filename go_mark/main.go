package main

import (
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

	app.Run(os.Args)
}
