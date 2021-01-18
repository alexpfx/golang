package commands

import "github.com/alexpfx/go_common/cmd"

func NewXSel() *cmd.Cmd{
	return &cmd.Cmd{
		Binary:    cmd.Binary{
			CmdPath: "xsel",
			FixArgs: []string{"-b"},
		},
		Pipe:      true,
		Converter: func(bytes []byte) ([]byte, error) {
			return bytes, nil
		},
	}
}

func NewJQ() *cmd.Cmd{

	return &cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "",
			Name:    "",
			Desc:    "",
			FixArgs: nil,
		},
		Pipe: true,
	}

}
