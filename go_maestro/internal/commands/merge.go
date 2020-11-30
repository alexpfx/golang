package commands

func MergeFetch() Cmd {
	return Cmd{
		Binary:    "merge",
		Name:      "Merge Fetch",
		Desc:      "Obtém Informações de um Merge Request",
		Args:      []string{"fetch", "-auto"},
		UserInput: map[string]string{"mergeId": ""},
		Clipboard: true,
	}
}