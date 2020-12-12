package main

import (
	"bufio"
	"flag"
	"fmt"
	"github.com/alexpfx/go_common/str"
	"github.com/alexpfx/golang/go_sibe/internal/sibe/script"
	"github.com/urfave/cli"
	"log"
	"os"
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
						Usage: "executa o script sibeDeploy.sh",
					},
				},
				Action: func(c *cli.Context) error {
					fmt.Println("executado :", c.Args())
					cl := c.Bool("client")
					fmt.Println(cl)

					if cl {
						rcl := script.NewRunner(script.SibeClient())
						pout, _ := rcl.Run(c.Args())

						scan := bufio.NewScanner(pout)
						go func() {
							for scan.Scan() {
								fmt.Println(scan.Text())
							}
						}()
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
