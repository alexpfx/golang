package go_chain

import (
	js "encoding/json"
	"github.com/pelletier/go-toml"
	"net/http"
	"regexp"
	"strings"
)

var tagPattern *regexp.Regexp = regexp.MustCompile("{{(.*?)}}")

///CreateConfig cria um novo arquivo de configuração básico
func CreateConfig(configDir, configFile string) *Chain {
	return nil
}

///ExecuteRequest é responsável por realizar a chamada HTTP
func ExecuteRequest(request Request) {

	if strings.EqualFold(request.Method, "get") {
		http.Get(request.Endpoint)
	}

}

func parseTomlFile(filePath string) (*Request, error) {
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
		Method:           out.Method,
		Output:           out.Output,
		Input:            out.Input,
		Endpoint:         out.Endpoint,
		EndpointReplaces: out.EndpointReplaces,
		Json:             out.Json,
	}, nil
}

func replaceAll(sourceStr string, replaceStr []string) string {
	strResult := sourceStr
	for i, s := range tagPattern.FindAllString(sourceStr, len(replaceStr)) {
		strResult = strings.Replace(strResult, s, replaceStr[i], -1)
	}

	return strResult
}

func replaceInput(json string, input map[string]string) (string, error) {
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
