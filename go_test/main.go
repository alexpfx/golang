package main

import (
	"encoding/json"
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/alexpfx/golang/go_test/internal/tests"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"log"
	"os"
	"reflect"
	"strings"
)

func main() {
	app := &cli.App{
		Commands: []*cli.Command{
			{
				Name:    "node",
				Usage:   "gera um nodo de teste de forma interativa",

				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:     "input",
						Aliases:  []string{"i"},
						Usage:    "arquivo json de entrada do teste",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "output",
						Aliases: []string{"o"},
						Usage:   "arquivo json de saída gerado pelo teste",
					},
					&cli.StringFlag{
						Name:     "name",
						Aliases:  []string{"n"},
						Usage:    "nome para o teste",
						Required: true,
					},
					&cli.StringFlag{
						Name:    "dir",
						Aliases: []string{"d"},
						Usage:   "diretório raiz onde os testes serão gerados",
					},
				},
				Action: func(c *cli.Context) error {
					inputFilePath := c.String("input")
					jsonInput, err := readJsonFile(inputFilePath)
					if err != nil {
						return err
					}
					fmt.Println(jsonInput)

					outputFilePath := c.String("output")
					var jsonOutput map[string]interface{}
					if outputFilePath != "" {
						jsonOutput, err = readJsonFile(outputFilePath)
						if err != nil {
							return err
						}
					}

					fmt.Println(jsonOutput)

					testName := c.String("name")
					fmt.Println(testName)
					rootDirectory := c.String("dir")
					fmt.Println(rootDirectory)

					var inputVariableFields []string
					inputVariableFields = promptVars(jsonInput)

					var outputVarFields []string
					outputVarFields = promptVars(jsonOutput)

					nodeInput := createTestNode(inputVariableFields, jsonInput, 0)
					nodeOutput := createTestNode(outputVarFields, jsonOutput, 1)

					fmt.Printf("input %v\n", nodeInput)
					fmt.Printf("output %v\n", nodeOutput)

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

func createTestNode(varFields []string, jsMap map[string]interface{}, t int) tests.Node {
	jsBytes, _ := json.Marshal(jsMap)

	return tests.Node{
		Type: t,
		Vars: varFields,
		Json: string(jsBytes),
	}

}

func promptVars(input map[string]interface{}) []string {
	items := make([]string, 0)

	items = traverseMap("", input)

	var res []string
	prompt := &survey.MultiSelect{
		Message: "Selecione os campos de entrada que serão variáveis:",
		Options: items,
	}
	_ = survey.AskOne(prompt, &res)

	return res
}

func traverseMap(key string, m map[string]interface{}) []string {
	res := make([]string, 0)

	for k, v := range m {
		kind := reflect.ValueOf(v).Kind()
		var tk string

		if key != "" {
			tk = strings.Join([]string{key, k}, "/")
		} else {
			tk = k
		}
		if kind == reflect.Map {
			res = append(res, traverseMap(tk, v.(map[string]interface{}))...)
		} else {
			res = append(res, tk)
		}
	}

	return res
}

func readJsonFile(inputFilePath string) (map[string]interface{}, error) {
	var result map[string]interface{}
	f, err := os.Open(inputFilePath)
	if err != nil {
		return nil, fmt.Errorf("não pode encontrar arquivo de entrada: %s", inputFilePath)
	}

	bytes, _ := ioutil.ReadAll(f)

	json.Unmarshal(bytes, &result)

	defer f.Close()

	return result, err

}
