package trees

import (
	"github.com/alexpfx/go_action/action"
	
	"github.com/alexpfx/go_action/input"
)

var (
	DtpTree = action.NewFzfTree("comando", []action.Action{
		{
			Name:   "Merge Info",
			Binary: &goMergeInfo,
			InputConfig: &input.ResolverConfig{
				Resolver: input.ScannerResolver{},
				Keys: []string{"mergeId"},
				Prompt: "informe o Id do merge",
			},
		},
	})
)
