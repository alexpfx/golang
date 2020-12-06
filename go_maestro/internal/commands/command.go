package commands

type Cmd struct {
	Binary        string
	Name          string
	Desc          string
	Args          []string
	UserInput     map[string]string
	CopyOutput    bool
	CopySelection bool
	FilterOutput  []string
	FormatOutput  []string
	OutputConverter func(out string) (string, []string)
	OutputFormat string
	CallNext      func(...string) *Cmd
}
