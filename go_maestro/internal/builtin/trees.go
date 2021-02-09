package builtin

import (
	"github.com/alexpfx/go_action/action"
	"github.com/alexpfx/go_action/builtin"
	
	"github.com/alexpfx/go_action/input"
)

var (
	dtpTree = action.NewFzfTree("comando", []action.Action{
		{
			Name:   "Merge Info",
			Binary: &goMergeInfo,
			InputConfig: &input.ResolverConfig{
				Resolver: builtin.ClipResolver{},
			},
			Next: &action.Action{
				Binary:        builtin.Jq,
				InputFromPipe: true,
			},
		},
	})
)
