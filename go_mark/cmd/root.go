package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "bmark",
	Short: "Gerenciador de bookmarks para o console",
	Long:  "",
}

//Init serve para facilitar testes
func Init(cmds []*cobra.Command) {
	for _, cmd := range cmds {
		rootCmd.AddCommand(cmd)
	}
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
