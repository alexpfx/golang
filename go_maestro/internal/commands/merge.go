package commands

import (
	"github.com/alexpfx/go_common/cmd"
	"github.com/alexpfx/go_common/user"
	"github.com/alexpfx/golang/go_maestro/internal/output"
	"strings"
)


func NewMergeFetch() *cmd.Cmd {
	input := cmd.Input{
		InputList: user.MultiInput{
			ArgSep: " ",
		},
		Reader: user.ClipInputReader{},
	}

	c := cmd.Cmd{
		Binary: cmd.Binary{
			CmdPath: "go_merge",
			Name:    "Merge Info",
			Desc:    "Obtém informações sobre um merge request",
			FixArgs: []string{"info"},
		},
		UserInput: &input,

	}
	return &c
}

func MergeFetch() Cmd {
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

	return Cmd{
		Binary:     "go_merge",
		Name:       "Merge Fetch",
		Desc:       "Obtém Informações de um Merge Request",
		Args:       []string{"info"},
		UserInput:  map[string]string{"mergeId": ""},
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
