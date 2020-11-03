package data

type CommandType int

const (
	MergeInfo CommandType = iota
)

type Result struct {
	CommandType CommandType
	Results interface{}
}
