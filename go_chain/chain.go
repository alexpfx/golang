package go_chain

type Chain struct {
	Request *Request
	Next    *Request
}

type Request struct {
	Method           string              `toml:"method"`
	Input            []string            `toml:"input"`
	Json             string              `toml:"json"`
	EndpointReplaces map[string][]string `toml:"endpointReplaces"`
	Endpoint         string              `toml:"endpoint"`
	Output           []string            `toml:"output"`
}
