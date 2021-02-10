package trees

import (
	"github.com/alexpfx/go_action/action"
)

var (
	goMergeInfo = action.Binary{
		CmdPath: "go_merge",
		Name:    "go_merge info",
		FixArgs: []string{"info"},
	}
)
