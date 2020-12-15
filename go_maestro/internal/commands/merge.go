package commands

func MergeFetch() Cmd {
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
	}

}


