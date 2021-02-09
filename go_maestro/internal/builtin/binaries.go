package builtin

import (
	"github.com/alexpfx/go_action/action/binary"
)

var (
	goMergeInfo = binary.Binary{
		CmdPath: "go_merge",
		Name:    "go_merge info",
		FixArgs: []string{"info"},
	}
)
