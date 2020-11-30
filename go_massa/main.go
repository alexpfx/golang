package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_massa/internal/massa"
	"log"
	"os"
)


const path_ambiente = "/ambiente"
const host_local = "http://localhost:7001"

func main() {
	args := os.Args
	listCmd := flag.NewFlagSet("list", flag.ExitOnError)
	getCmd := flag.NewFlagSet("get", flag.ExitOnError)

	var amb, cat int
	listCmd.IntVar(&amb, "a", 2, "Número do ambiente {1 - desenvolvimento | 2 - homologação}")

	getCmd.IntVar(&cat, "c", 8, "Número do catálogo massa-sibe")
	getCmd.IntVar(&amb, "a", 2, "Número do ambiente {1 - desenvolvimento | 2 - homologação}")

	switch args[1] {
	case "list":
		err := listCmd.Parse(args[2:])
		if err != nil {
			listCmd.PrintDefaults()
			os.Exit(1)
		}

		list := massa.NewCatalogoList()
		catalogos, err := list.Catalogos(amb)
		if err != nil {
			log.Fatal(err.Error())
		}
		fmt.Println(ToJsonStr(catalogos))

	case "get":
		err := getCmd.Parse(args[2:])

		getter := massa.NewMassaGetter()
		massa, err := getter.GetRecent(cat, amb)
		if err != nil {
			listCmd.PrintDefaults()
			os.Exit(0)
		}
		fmt.Println(ToJsonStr(massa))


	default:
		flag.PrintDefaults()
		os.Exit(1)
		//log.Fatalf("argumento(s) inválido(s): %v", args[1:])

	}

}

func ToJsonStr(results interface{}) string {
	bytes, err := json.MarshalIndent(results, "", "   ")
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
