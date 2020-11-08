package go_chain

import (
	js "encoding/json"
	"github.com/pelletier/go-toml"
	"strings"
)

func Parse(filePath string) (*Request, error) {
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return nil, err
	}

	var x Request

	err = tree.Unmarshal(&x)
	if err != nil {
		return nil, err
	}

	jsn := strings.Replace(x.Json, "\n", "", -1)

	return &Request{
		Method:   x.Method,
		Output:   x.Output,
		Input:    x.Input,
		Endpoint: x.Endpoint,
		Json:     jsn,
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
