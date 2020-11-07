package go_chain_test

import (
	"fmt"
	"github.com/pelletier/go-toml"
	"testing"
)

func TestParse(t *testing.T) {
	tree, err := toml.LoadFile("data/judicial.toml")
	if err != nil {
		fmt.Println(err)
	}
	t.Log("Tree ", tree)
}
