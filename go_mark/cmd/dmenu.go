/*
Copyright © 2020 NAME HERE <EMAIL ADDRESS>

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/
package cmd

import (
	"bmark/bookmark"
	"fmt"
	"github.com/spf13/cobra"
	"os/exec"
)

// dmenuCmd represents the dmenu command
var dmenuCmd = &cobra.Command{
	Use:   "dmenu",
	Short: "A brief description of your command",
	Long: `A longer description that spans multiple lines and likely contains examples
and usage of using your command. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	Run: func(cmd *cobra.Command, args []string) {
		bookmarker := bookmark.ReadBookmarks()
		items := bookmarker.All()

		if len(items) == 0 {
			return
		}

		openRofi(items)

	},
}

func openRofi(items []bookmark.Item) {

	rofiCmd := exec.Command("/usr/bin/rofi", "-dmenu", "-p", "Selecione para abrir no browser", "-multi-select")
	stdin, err := rofiCmd.StdinPipe()
	if stdin == nil {
		return
	}

	for _, item := range items {
		_, err = fmt.Fprintf(stdin, " %s[%s]\n", item.Url, item.Desc)
	}

	err = stdin.Close()
	checkPanic(err)
	output, err := rofiCmd.Output()
	checkExitOk(err)

	fmt.Println(string(output))

}

func init() {
	rootCmd.AddCommand(dmenuCmd)

}
