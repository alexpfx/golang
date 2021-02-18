package trees

import (
	"github.com/alexpfx/go_action/action"
	
	"github.com/alexpfx/go_action/input"
)


var (

	DtpTree = action.NewFzfTree("dtp","comando", []action.Action{
		{
			Name:   "Merge Info",
			Binary: &goMergeInfo,
			InputConfig: &input.ResolverConfig{
				Resolver: input.FzfResolver{},
				Keys: []string{"mergeId"},
				Prompt: "informe o Id do merge",
			},
		},
		{
			Name:   "Duplicar código",
			Binary: &goMergeInfo,
			InputConfig: &input.ResolverConfig{
				Resolver: input.FzfResolver{},
				Keys: []string{"mergeId"},
				Prompt: "informe o Id do merge",
			},
		},
	})
	BuiltInTrees = []action.Tree{
		DtpTree,
	}
)
