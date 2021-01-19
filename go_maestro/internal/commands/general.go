package commands

import "github.com/alexpfx/go_common/cmd"

func NewXSel() *cmd.Cmd {
	return &cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "xsel",
			FixArgs: []string{"-b"},
		},
		PipeIntoNext: true,
		OutConverter: func(bytes []byte) ([]byte, error) {
			return bytes, nil
		},
	}
}

func NewRofiMenu() *cmd.Cmd {
	return &cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "rofi",
			Name:    "",
			Desc:    "",
			FixArgs: []string{"-dmenu", "-sep", "\n", "-format", "s"},
		},
	}
}

func NewSelectMonitor(identifier, monitor string) *cmd.Cmd {
	return &cmd.Cmd{
		Identifier: identifier,
		Binary: cmd.Binary{
			CmdPath: "bspc",
			Name:    "Foco no Monitor " + monitor,
			FixArgs: []string{"query", "-m", monitor, "--desktops", "--names"},
		},
		Next:   NewRofiMenu(),
		PipeIntoNext:      true,
		OutConverter: nil,
	}
}



func NewJQ() *cmd.Cmd {

	return &cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "",
			Name:    "",
			Desc:    "",
			FixArgs: nil,
		},
	}

}
