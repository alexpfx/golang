package actions

import (
	"github.com/alexpfx/go_action/action"
	"github.com/alexpfx/go_action/input"
)

func NewMergeFetch() action.Action {
	input := action.InputConfig{
		Config: input.Config{
			ArgSep: " ",
		},
		Resolver: input.ClipResolver{},
	}

	c := action.Action{
		Binary: action.Binary{
			CmdPath: "go_merge",
			Name:    "Merge Info",
			Desc:    "Obtém informações sobre um merge request",
			FixArgs: []string{"info"},
		},
		InputConfig: &input,
	}
	return c
}

/*func MergeFetch() cmd.Action {
	defaultFmt := []string{
		"#.merge.web_url",
		"#.merge.author.username",
		"#.merge.commit.username",
		"#.merge.commit.created_at",
	}

	formatMap := map[string][]string{
		"desenvolvimento": defaultFmt,
		"homologacao": {
			"#.merge.web_url",
			"#.merge.author.username",
			"#.merge.commit.created_at",
		},
	}

	return cmd.Action{
		Binary:     "go_merge",
		Name:       "Merge Fetch",
		Desc:       "Obtém Informações de um Merge Request",
		Args:       []string{"info"},
		Config:  map[string]string{"mergeId": ""},
		CopyOutput: true,
		DynamicFormatOutput: func(out string) []string {
			tbranch := output.Format(out, []string{"#.merge.target_branch"})
			branchFormat, ok := formatMap[strings.TrimSpace(tbranch)]
			if ok {
				return branchFormat
			}
			return defaultFmt
		},
	}

}
*/
