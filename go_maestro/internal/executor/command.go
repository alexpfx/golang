package executor

type Command interface {
	Name() string
	Args() []string
	UserInput() map[string]string
	Clipboard() bool
}

func NewCmd(name string, args []string, userInput map[string]string, clipboard bool) Command {
	return Cmd{
		name:      name,
		args:      args,
		userInput: userInput,
		clipboard: clipboard,
	}
}


type Cmd struct {
	name      string
	args      []string
	userInput map[string]string
	clipboard bool
}

func (c Cmd) Clipboard() bool {
	return c.clipboard
}

func (c Cmd) UserInput() map[string]string {
	return c.userInput

}

func (c Cmd) Name() string {
	return c.name
}

func (c Cmd) Args() []string {
	return c.args
}


