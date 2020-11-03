package main

import (
	"bmark/cmd"

	"github.com/spf13/cobra"
)

func main() {

	cmd.Init([]*cobra.Command{
		cmd.NewSearchCmd(), cmd.NewAddCmd(),
	})

	cmd.Execute()
}
