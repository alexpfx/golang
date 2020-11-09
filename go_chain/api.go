package go_chain

import (
	js "encoding/json"
	"github.com/pelletier/go-toml"
)

func Parse(filePath string) (*Request, error) {
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	var out Request

	err = tree.Unmarshal(&out)
	if err != nil {
		return nil, err
	}

	return &Request{
		Method:   out.Method,
		Output:   out.Output,
		Input:    out.Input,
		Endpoint: out.Endpoint,
		Json:     out.Json,
	}, nil
}


func ReplaceInput(json string, input map[string]string) (string, error) {
	var msgMap map[string]interface{}

	err := js.Unmarshal([]byte(json), &msgMap)

	if err != nil {
		return "", err
	}

	for key, value := range input {
		msgMap[key] = value
	}

	marshal, err := js.Marshal(msgMap)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
