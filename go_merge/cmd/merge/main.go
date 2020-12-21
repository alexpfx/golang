package main

import (
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_merge/internal/merge"
	"log"
	"os"
)

const baseUrl = "https://www-scm.prevnet/api/v3/projects"
const sibeProject = "754"

func main() {
	fetchCmd := flag.NewFlagSet("fetch", flag.ExitOnError)

	var author string
	fetchCmd.StringVar(&author, "a", "", "Filtro por autor")

	var targetBranch string
	fetchCmd.StringVar(&targetBranch, "b", "", "Filtro por target branch ex: desenvolvimento, homologacao")

	var output string
	fetchCmd.StringVar(&output, "o", "", "Mostra apenas os campos passados como parâmetro ex: <.author.username>. Para ver os campos possíveis use o argumento -fields ")


	outputFlags := map[string]*bool{
		"dev":  new(bool),
		"hom":  new(bool),
		"auto": new(bool),
		"json": new(bool),
	}

	fetchCmd.BoolVar(outputFlags["dev"], "dev", true, "Imprime saída de Dev: <Url> <Solicitante> <Aprovador> <Data>")
	fetchCmd.BoolVar(outputFlags["hom"], "hom", false, "Imprime saída de Hom: <Url> <Solicitante> <Data>")
	fetchCmd.BoolVar(outputFlags["auto"], "auto", false, "Imprime saída de acordo com o Target Branch desenvolvimento ou homologacao")
	fetchCmd.BoolVar(outputFlags["json"], "json", false, "Imprime saída completa em formato Json")

	args := os.Args
	checkArgLen(args, 2)

	var help bool
	fetchCmd.BoolVar(&help, "h", false, "Mostra a ajuda")

	switch args[1] {

	case "fetch":

		err := fetchCmd.Parse(args[2:])

		if help {
			fetchCmd.PrintDefaults()
			return
		}

		if err != nil {
			log.Fatal(err.Error())
			
		}

		ids, err := merge.ParseIds(fetchCmd.Args())
		if err != nil {
			log.Fatal(err)
		}

		//ids = slices.IntUniqueSorted(ids)

		checkIdsCount(ids)
		token := checkToken()
		filter := mapFilter(author, targetBranch)
		mrInfo, er, err := merge.Fetch(token, baseUrl, sibeProject, ids, filter)
		if err != nil {
			fmt.Println(err)
		}
		if len(er) > 0 {
			fmt.Println("Invalidos: ")
			for _, result := range er {
				fmt.Printf("%d: %v", result.MergeId, result.Err)
			}
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

//133
