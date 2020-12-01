package commands

import "reflect"

type Cmd struct {
	Binary string
	Name      string
	Desc string
	Args      []string
	UserInput map[string]string
	Clipboard bool
	FilterOutput []string
	Filter reflect.Type
	Next *Cmd
}


