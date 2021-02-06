package actions

import (
	"github.com/alexpfx/go_action/action"
	"strings"
)

func NewXSel() *action.Action {
	return &action.Action{
		Binary: action.Binary{
			CmdPath: "xsel",
			FixArgs: []string{"-b"},
		},
		InputFromPipe: true,
		Converter: func(bytes []byte) ([]byte, error) {
			return bytes, nil
		},
	}
}

var rofi = action.Binary{
	CmdPath: "rofi",
	Name:    "",
	Desc:    "",
	FixArgs: []string{"-dmenu", "-sep", "\n", "-format", "s"},
}

var bspcNode = action.Binary{
	CmdPath: "bspc",
	FixArgs: []string{"node"},
}

var bspcQuery = action.Binary{
	CmdPath: "bspc",
	Name:    "BSPC Query ",
	Desc:    "",
	FixArgs: []string{"query"},
}

func FocusMonitor() []action.Action {
	displays := getDisplays()

	var bspcFocusDesktop = action.Binary{
		CmdPath: "bspc",
		FixArgs: []string{"desktop", "-f"},
	}
	next := &action.Action{
		Binary: rofi,
		Next: &action.Action{
			Binary: bspcFocusDesktop,
		},
	}

	res := make([]action.Action, 0)
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

func SelectMonitor(identifier, monitor string, next *action.Action) action.Action {

	return action.Action{
		Identifier:    identifier,
		Binary:        bspcQuery,
		Args:          []string{"-m", monitor, "--desktops", "--names"},
		InputFromPipe: true,
		Next:          next,
		Converter:     nil,
	}
}

func NewJQ() *action.Action {

	return &action.Action{
		Binary: action.Binary{
			CmdPath: "",
			Name:    "",
			Desc:    "",
			FixArgs: nil,
		},
	}

}
