package main

import (
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_todo/internal/api"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

var finalistArg string

func main() {

	var rounds = flag.Int("rounds", 20, "rounds")
	flag.StringVar(&finalistArg, "f", "", "-f '1 2 3'")
	flag.Parse()

	listener := make(chan api.Info, 0)

	x := *rounds / 3

	var finalistsIndex []string
	if finalistArg != "" {
		finalistsIndex = strings.Split(finalistArg, " ")
	}
	game := api.NewGame(createTodos(finalistsIndex), *rounds, x, 1)

	go game.Start(listener)

	for i := 0; i < (*rounds); i++ {

		clear()
		fmt.Println(<-listener)
		time.Sleep(time.Millisecond * 1000)

	}

}

func clear() {
	cmd := exec.Command("clear") //Linux example, its tested
	cmd.Stdout = os.Stdout
	cmd.Run()
}

func createTodos(index []string) []api.Todo {
	all := createAll()
	todos := make([]api.Todo, 0)

	if index != nil {
		for _, s := range index {
			index, _ := strconv.Atoi(s)
			todos = append(todos, all[index])
		}
	} else {
		todos = all
	}

	for i, s := range todos {
		todos[i].Name = fmt.Sprintf("%v_%d: %4s", string(rune(i+65)), i, s)
	}
	return todos
}



func createAll() []api.Todo {
	todos := []api.Todo{
		{"Lavar louça"}, //A
		{"Varrer areia e limpar caixa rosa banheiro quarto"},
		{"Varrer areia e limpar caixa banheiro sala"},
		{"Varrer sala"},
		{"recolher e Estender roupa"},
		{"Colocar lençois"},
		{"Levar Lixo"},
		{"Varrer quarto computador"},
		{"Varrer quarto dormir"},
//		{"Recolher e ensacar lixo"},
		{"Limpar pias banheiro"},
		{"Guardar compras"},
		{"Varrer cozinha"},
		{"Recolher Roupa"},
		{"Recolher coisas fora do lugar"},
		{"Recolher louças fora do lugar"},
	}
	return todos
}
