package go_chain_test

import (
	"fmt"
	"github.com/alexpfx/golang/go_chain"
	"github.com/pelletier/go-toml"
	"testing"
)

func TestParse(t *testing.T) {
	filePath := "data/judicial.toml"
	tree, err := toml.LoadFile(filePath)
	var x go_chain.Request

	tree.Unmarshal(&x)
	if err != nil {
		fmt.Println(err)
	}

	go_chain.Parse(filePath)

}
