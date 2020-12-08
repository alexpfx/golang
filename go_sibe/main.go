package main

import (
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_sibe/internal/sibe/script"
	"os"
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
	var outFmt string

	deploysCmd.StringVar(&outFmt, "-fmt", "%3v %v\n", "formato de cada linha de saida. ex: 'Id: %v Script: %v \\n'")
	clientsCmd.StringVar(&outFmt, "-fmt", "%3v %v\n", "formato de cada linha de saida. ex: 'Id: %v Script: %v \\n'")

	args := os.Args
	if len(args) < 2 {
		flag.Usage()
		os.Exit(1)
	}

	switch args[1] {
	case "deploys":
		_ = deploysCmd.Parse(args[2:])
		for _, s := range script.DeployScripts {
			fmt.Printf(outFmt, s.Id, s.Name)
		}
	case "clients":
		_ = clientsCmd.Parse(args[2:])

		for _, s := range script.ClientScripts {
			fmt.Printf(outFmt, s.Id, s.Name)
		}
	}
}
