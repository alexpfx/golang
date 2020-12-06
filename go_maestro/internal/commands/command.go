package commands

type Cmd struct {
	Binary     string
	Name       string
	Desc       string
	Args       []string
	UserInput  map[string]string
	CopyOutput bool
	CopySelection bool
	FilterOutput []string
	FormatOutput []string
	Next *Cmd
}


