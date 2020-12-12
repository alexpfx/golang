package main

import (
	"encoding/json"
	
	"fmt"
	"github.com/alexpfx/golang/go_sibe/internal/sibe/script"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

func main() {

	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name: "deploys",
				Action: func(c *cli.Context) error {

					printJson(script.DeployScripts)

					return nil

				},
			},
			{
				Name: "clients",
				Action: func(c *cli.Context) error {
					printJson(script.ClientScripts)
					return nil
				},
			},
		},
	}

	app.Run(os.Args)
}

func printJson(all []script.Script) {
	bytes, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bytes))
}
