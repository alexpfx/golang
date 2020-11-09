package go_chain_test

import (
	"fmt"
	"github.com/alexpfx/golang/go_chain"
	"github.com/stretchr/testify/assert"
	"testing"
)

const errFragment = "espera que erro fosse %v, porém obteve %v"

func TestReplaceInput(t *testing.T) {
	tests := []struct {
		name     string
		input    map[string]string
		contains string
	}{
		{"replace cpf", map[string]string{"cpf": "1122"}, `"cpf":"1122"`},
		{"replace nit", map[string]string{"nit": "1234", "nome": "alexandre"}, `"nit":"1234"`},
		{"replace nome", map[string]string{"nit": "1234", "nome": "alexandre"}, `"nome":"alexandre"`},
	}
	filePath := fmt.Sprintf("data/%s", "inclusao1.toml")

	parse, err := go_chain.Parse(filePath)
	if err != nil {
		t.Error(err)
	}

	json := parse.Json

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input, err := go_chain.ReplaceInput(json, test.input)
			checkErr(t, err)

			assert.Contains(t, input, test.contains)

		})
	}
}

func checkErr(t *testing.T, err error) {
	if err != nil {
		t.Errorf("erro inexperado: %v", err)
	}
}
func TestParsing(t *testing.T) {
	tests := []struct {
		name     string
		file     string
		expected go_chain.Request
		e        error
	}{
		{"judicial", "judicial.toml", go_chain.Request{
			Method:   "post",
			Endpoint: "{{ baseUrl }}/PortalSibe",
			Input:    []string{"cpf", "nome"},
			Output:   []string{"nb"},
		}, nil},
		{"massa", "massa.toml", go_chain.Request{
			Method: "get",
		}, nil},
		{"inclusao1", "inclusao1.toml", go_chain.Request{
			Method:   "post",
			Endpoint: "{{ baseUrl }}/sibews/rest/requerimento/incluir/bi/inicial",
			Input:    []string{"cpf", "nomeTitular"},
			Output:   []string{"nb"},
		}, nil},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			filePath := fmt.Sprintf("data/%s", test.file)

			r, e := go_chain.Parse(filePath)

			if r == nil {

				t.Error("nao esperava valor nulo")
				return
			}

			t.Log(test.name)

			if test.e != e {

				t.Errorf(errFragment, test.e, e)
				return
			}

			expected := test.expected

			assert.Equal(t, r.Endpoint, expected.Endpoint)
			assert.Equal(t, r.Method, expected.Method)
			assert.Equal(t, r.Json, expected.Json)

			assert.ElementsMatch(t, r.Input, expected.Input)

			assert.ElementsMatch(t, r.Output, expected.Output)

		})
	}
}
