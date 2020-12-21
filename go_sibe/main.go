package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/golang/go_sibe/internal/sibe/script"
	"github.com/urfave/cli"
)

var deploysCmd *flag.FlagSet
var clientsCmd *flag.FlagSet

func usage() {
	deploysCmd.PrintDefaults()
}

func main() {
	app := &cli.App{
		Commands: []cli.Command{
			{
				Name:  "call",
				Usage: "executa um ou mais scripts",
				Flags: []cli.Flag{
					&cli.BoolFlag{
						Name:  "deploy",
						Usage: "executa o script sibeDeploy.sh",
					},
					&cli.BoolFlag{

						Name:  "client",
						Usage: "executa o script sibeClient.sh",
					},
				},
				Action: func(c *cli.Context) error {
					cl := c.Bool("client")

					if cl {
						rcl := script.NewRunner(script.SibeClient())
						ch := make(chan string, 2)

						go rcl.Run(c.Args(), ch)

						for m := range ch {
							fmt.Println(m)
						}

					}

					return nil
				},
			},
			{
				Name:  "clients",
				Usage: "lista as opções de sibeClient",
				Action: func(c *cli.Context) error {
					jsonStr, _ := str.FormatJson(script.ClientScripts)
					fmt.Println(jsonStr)
					return nil
				},
			},
			{
				Name:  "deploys",
				Usage: "lista as opções de sibeDeploy",
				Action: func(c *cli.Context) error {
					jsonStr, _ := str.FormatJson(script.DeployScripts)
					fmt.Println(jsonStr)
					return nil
				},
			},
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
