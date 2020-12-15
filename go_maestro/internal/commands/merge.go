package commands

import (
	"encoding/json"
)

func MergeFetch() Cmd {
	defaultFmt := []string{
		"#.merge.web_url",
		"#.merge.author.username",
		"#.merge.commit.created_at",
	}

	formatMap := map[string][]string{
		"desenvolvimento": {
			"#.merge.web_url",
			"#.merge.author.username",
			"#.merge.commit.username",
			"#.merge.commit.created_at",
		},
		"homologacao": defaultFmt,
	}

	return Cmd{
		Binary:     "go_merge",
		Name:       "Merge Fetch",
		Desc:       "Obtém Informações de um Merge Request",
		Args:       []string{"info"},
		UserInput:  map[string]string{"mergeId": ""},
		CopyOutput: true,
		FormatOutput: []string{
			"#.merge.web_url",
			"#.merge.author.username",
			"#.merge.commit.username",
			"#.merge.commit.created_at",
		},
		DynamicFormatOutput: func(output string) []string {
			var objmap map[string]string

			_ = json.Unmarshal([]byte(output), &objmap)
			branch := objmap["target_branch"]
			v, ok := formatMap[branch]
			if ok {
				return v
			}
			return defaultFmt
		},
	}

}
