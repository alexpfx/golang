package main

import (
	"fmt"
	"github.com/alexpfx/golang/go_merge/internal/merge"
	"os"
)

func main() {
	mrs, err := merge.ParseIds(os.Args)
	if err != nil {
		fmt.Errorf("erro ao extrair merges")
	}

}