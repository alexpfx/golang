package commands


type Cmd struct {
	Binary string
	Name      string
	Desc string
	Args      []string
	UserInput map[string]string
	Clipboard bool
	FilterOutput []string
}


