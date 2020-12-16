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
	"github.com/BurntSushi/xgbutil"
	"github.com/BurntSushi/xgbutil/keybind"
	"github.com/BurntSushi/xgbutil/xevent"
	"github.com/spf13/cobra"
	"log"
)

// serviceCmd represents the service command
var serviceCmd = &cobra.Command{
	Use:   "service",
	Short: "comando usado com systemd",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {


		connectToX()

	},
}

func connectToX() {
	xConn, err := xgbutil.NewConn()
	checkError(err)

	keybind.Initialize(xConn)

	fu := func(xu *xgbutil.XUtil, event xevent.KeyPressEvent) {
		log.Printf("keycode: %v", keybind.LookupString(xConn, event.State, event.Detail))

		checkError(err)
	}

	//err = keybind.KeyReleaseFun(fu).Connect(xConn, xConn.RootWin(), "Menu-g", true)
	err = keybind.KeyPressFun(fu).Connect(xConn, xConn.RootWin(), "Menu-g", true)
	checkError(err)

	xevent.Main(xConn)

}

func checkError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func init() {
	rootCmd.AddCommand(serviceCmd)
	rootCmd.Flags()

}
