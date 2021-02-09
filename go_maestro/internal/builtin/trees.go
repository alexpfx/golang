package builtin

import (
	"github.com/alexpfx/go_action/action"
	"github.com/alexpfx/go_action/action/builtin"
	"github.com/alexpfx/go_action/input"
)

var (
	dtpTree = action.NewFzfTree("comando", []action.Action{
		{
			Name:   "Merge Info",
			Binary: &goMergeInfo,
			InputConfig: &input.Config{
				Resolver: input.ClipResolver{},
			},
			Next: &action.Action{
				Binary:        builtin.Jq,
				InputFromPipe: true,
			},
		},
	})
)
