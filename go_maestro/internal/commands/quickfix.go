package commands

const quickFixBinary = "go_quickfix"

func QuickFixExecute() Cmd {
	return Cmd{
		Binary:              quickFixBinary,
		Name:                "Executa",
		Desc:                "",
		Args:                []string{},
		UserInput:           nil,
		CopyOutput:          false,
		FilterOutput:        nil,
		FormatOutput:        nil,
		DynamicFormatOutput: nil,
		OutputConverter:     nil,
		CallNext:            nil,
	}

}

func QuickFixQuery() Cmd {
	return Cmd{
		Binary:              quickFixBinary,
		Name:                "Quick fix",
		Desc:                "Quick fix query",
		Args:                []string{"--query"},
		UserInput:           map[string]string{"query": ""},
		FilterOutput:        nil,
		FormatOutput:        []string{"#.name"},
		DynamicFormatOutput: nil,
		OutputConverter:     nil,
		CallNext:            nil,
	}
}
