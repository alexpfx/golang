package tests

type NodeType int

const (
	Input NodeType = 1 << iota
	Output
)

type Node struct {
	Type NodeType `json:"type"`
	Vars []string `json:"variable"`
	Json string   `json:"json"`
}
