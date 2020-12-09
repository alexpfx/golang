package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/alexpfx/golang/go_sibe/internal/sibe/script"
)

var deploysCmd *flag.FlagSet
var clientsCmd *flag.FlagSet

func usage() {
	deploysCmd.PrintDefaults()
}
func main() {
	flag.Usage = usage

	deploysCmd = flag.NewFlagSet("deploys", flag.ExitOnError)
	clientsCmd = flag.NewFlagSet("clients", flag.ExitOnError)
	runScriptCmd = flag.NewFlagSet("run", flag.ExitOnError)

	args := os.Args
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch args[1] {
	case "deploys":
		_ = deploysCmd.Parse(args[2:])
		printJson(script.DeployItems)
	case "clients":
		_ = clientsCmd.Parse(args[2:])
		printJson(script.ClientsItems)
	case "run":

	}
}

func printJson(all []script.Item) {
	bytes, err := json.MarshalIndent(all, "", "  ")
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Println(string(bytes))
}
