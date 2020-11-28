package executor

type Command interface {
	Name() string
	Args() []string
	UserInput() map[string]string
}

func NewCmd(name string, args []string, userInput map[string]string) Command {
	return cmd{
		name:      name,
		args:      args,
		userInput: userInput,
	}

}

type cmd struct {
	name      string
	args      []string
	userInput map[string]string
}

func (c cmd) UserInput() map[string]string {
	return c.userInput

}

func (c cmd) Name() string {
	return c.name
}

func (c cmd) Args() []string {
	return c.args
}
