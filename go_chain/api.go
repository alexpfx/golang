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
	var x Request

	err = tree.Unmarshal(&x)
	check(err)

	//unquoted, err := strconv.Unquote(x.Json)

	message := js.RawMessage(x.Json)

	unquoted, _ := message.MarshalJSON()

	request, err2 := check(err)
	if err2 != nil {
		return request, err2
	}
	return &Request{
		Method:   x.Method,
		Output:   x.Output,
		Input:    x.Input,
		Endpoint: x.Endpoint,
		Json:     string(unquoted),
	}, nil
}

func check(err error) (*Request, error) {
	if err != nil {
		return nil, err
	}
	return nil, nil
}

func ReplaceInput(json string, input map[string]string) (string, error) {
	var value interface{}
	json.replace(/\\n/g, '')
	err := js.Unmarshal([]byte(json), value)

	if err != nil {
		return "", err
	}

	var msgMap = value.(map[string]string)

	for key, value := range input {
		msgMap[key] = value
	}

	marshal, err := js.Marshal(value)
	if err != nil {
		return "", err
	}
	return string(marshal), nil
}
