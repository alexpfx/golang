package main

import (
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_merge/internal/merge"
	"github.com/alexpfx/golang/go_merge/internal/slices"
	"log"
	"os"
)

const baseUrl = "https://www-scm.prevnet/api/v3/projects"
const sibeProject = "754"

func main() {
	fetchCmd := flag.NewFlagSet("fetch", flag.ExitOnError)

	var author string
	fetchCmd.StringVar(&author, "a", "", "filtro por author. ex: -a nome.sobrenome")

	var targetBranch string
	fetchCmd.StringVar(&targetBranch, "b", "", "-b desenvolvimento")

	var output string
	fetchCmd.StringVar(&output, "o", "", `-o ".author.username "`)

	outputFlags := map[string]*bool{
		"dev":  new(bool),
		"hom":  new(bool),
		"auto": new(bool),
		"json": new(bool),
	}

	fetchCmd.BoolVar(outputFlags["dev"], "dev", false, "--dev")
	fetchCmd.BoolVar(outputFlags["hom"], "hom", false, "--hom")
	fetchCmd.BoolVar(outputFlags["auto"], "auto", false, "--auto")
	fetchCmd.BoolVar(outputFlags["json"], "json", false, "--json")

	args := os.Args
	checkArgLen(args, 2)
	switch args[1] {

	case "fetch":
		checkArgLen(args, 3)
		err := fetchCmd.Parse(args[2:])

		if err != nil {
			log.Fatal(err.Error())
		}

		ids, err := merge.ParseIds(fetchCmd.Args())
		if err != nil {
			log.Fatal(err)
		}

		ids = slices.IntUniqueSorted(ids)

		checkIdsCount(ids)
		token := checkToken()
		filter := mapFilter(author, targetBranch)
		mrInfo, _, err := merge.Fetch(token, baseUrl, sibeProject, ids, filter)
		if err != nil {
			fmt.Println(err)
		}
		formatter := chooseFormatter(output, outputFlags)

		formatedOutput := merge.FormatOutput(mrInfo, formatter)
		fmt.Println(formatedOutput)

	default:
		log.Fatalf("argumento(s) inválido(s): %v", args[1:])
	}

}

func chooseFormatter(output string, flags map[string]*bool) merge.Formatter {
	if output != "" {
		return merge.NewFormatter(output)
	}
	if *flags["auto"] {
		return merge.FormatAuto
	}
	if *flags["dev"] {
		return merge.FormatDev
	}
	if *flags["hom"] {
		return merge.FormatHom
	}
	if *flags["json"] {
		return merge.FormatJson
	}


	return merge.FormatDev
}

func checkArgLen(args []string, length int) {
	if len(args) < length {
		log.Fatal("número de argumentos inválido")
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

func checkToken() string {
	token := os.Getenv("PRIVATE_TOKEN")
	if token == "" {
		log.Fatal(fmt.Errorf("token inválido. o token deve ser passado a aplicação através da variável de ambiente PRIVATE_TOKEN"))
	}
	return token
}

func checkIdsCount(ids []int) {
	if len(ids) > 50 {
		err := fmt.Errorf("máximo de %v ids por chamada", 50)
		log.Fatal(err)
	}
}
