package commands

import (
	"github.com/alexpfx/golang/go_maestro/internal/output"
	"strings"
)

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
