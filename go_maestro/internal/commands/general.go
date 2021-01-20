package commands

import (
	"github.com/alexpfx/go_common/cmd"
	"strings"
)

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

var rofi = cmd.Binary{
	CmdPath: "rofi",
	Name:    "",
	Desc:    "",
	FixArgs: []string{"-dmenu", "-sep", "\n", "-format", "s"},
}

var bspcNode = cmd.Binary{
	CmdPath: "bspc",
	FixArgs: []string{"node"},
}

var bspcQuery = cmd.Binary{
	CmdPath: "bspc",
	Name:    "BSPC Query ",
	Desc:    "",
	FixArgs: []string{"query"},
}

func FocusMonitor() []cmd.Cmd {
	displays := getDisplays()

	var bspcFocusDesktop = cmd.Binary{
		CmdPath: "bspc",
		FixArgs: []string{"desktop", "-f"},
	}
	next := &cmd.Cmd{
		Binary: rofi,
		Next: &cmd.Cmd{
			Binary: bspcFocusDesktop,
		},
	}

	res := make([]cmd.Cmd, 0)
	for _, display := range displays {
		selectMonitor := SelectMonitor(join("focus", display), display, next)
		res = append(res, selectMonitor)
	}
	return res
}

func join(prefix, display string) string {
	return strings.Join([]string{
		prefix, strings.ToLower(display),
	}, "_")
}

func getDisplays() []string {
	return []string{
		"HDMI1", "HDMI2", "DP1",
	}
}

func SelectMonitor(identifier, monitor string, next *cmd.Cmd) cmd.Cmd {

	return cmd.Cmd{
		Identifier:   identifier,
		Binary:       bspcQuery,
		Args:         []string{"-m", monitor, "--desktops", "--names"},
		PipeIntoNext: true,
		Next:         next,
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
