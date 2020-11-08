package go_chain_test

import (
	"fmt"
	"github.com/alexpfx/golang/go_chain"
	"testing"
)

const fragment = "%s: esperava %v, porém obteve %v"
const errFragment = "espera que erro fosse %v, porém obteve %v"

func TestReplaceInput(t *testing.T) {
	tests := []struct {
		name  string
		input map[string]string
	}{
		{"", map[string]string{"cpf": "1122"}},
	}
	filePath := fmt.Sprintf("data/%s", "inclusao1.toml")

	parse, err := go_chain.Parse(filePath)
	if err != nil {
		t.Error(err)
	}

	json := parse.Json

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			input, err := go_chain.ReplaceInput(json, map[string]string{"cpf": "1122"})
			checkErr(t, err)

			t.Log(input)

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

			t.Log(test.name)

			if test.e != e {

				t.Errorf(errFragment, test.e, e)
				return
			}

			expected := test.expected

			if r.Endpoint != expected.Endpoint {
				t.Errorf(fragment, "endpoint", expected.Endpoint, r.Endpoint)
			}

			if r.Method != expected.Method {
				t.Errorf(fragment, "method", expected.Method, r.Method)
			}

			for i, input := range r.Input {
				expectedInput := expected.Input[i]
				if input != expectedInput {
					t.Errorf(fragment, "input", expectedInput, input)
				}
			}

			for i, output := range r.Output {
				expectedOutput := expected.Output[i]
				if output != expectedOutput {
					t.Errorf(fragment, "output", expectedOutput, output)
				}
			}

		})
	}
}
