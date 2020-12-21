package commands

func QuickFixActions() Cmd{
	return Cmd{
		Binary:              "go_quickfix",
		Name:                "quick fix",
		Desc:                "quick fix",
		Args: []string{"-x", "--query"},
		UserInput: map[string]string{"query":""},
		CopyOutput:          false,
		FilterOutput:        nil,
		FormatOutput:        nil,
		DynamicFormatOutput: nil,
		OutputConverter:     nil,
		CallNext:            nil,
	}
}
