package main

import (
	"fmt"
	"github.com/alexpfx/golang/go_maestro/internal/executor"
	clip "github.com/atotto/clipboard"
	"io"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {

	cmds := []executor.Command{
		executor.NewCmd("merge", []string{
			"fetch",
			"-auto",
		}, map[string]string{
			"mergeId": "",
		}),
	}

	output := callRofi(buildRofiFromCmds(cmds))

	chosenCmd := parseChosenCmd(output, cmds)

	if chosenCmd == nil {
		return
	}

	var ua []string
	if len(chosenCmd.UserInput()) != 0 {
		ua = appendUserArgs(chosenCmd)
		fmt.Println(ua)
	}
	callCmd(chosenCmd, ua)

}

func callCmd(cmd executor.Command, ua []string) {
	args := append(cmd.Args(), ua...)

	command := exec.Command(cmd.Name(), args...)

	command.Env = os.Environ()

	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Println(string(output))
	}
	fmt.Println(string(output))
}

func appendUserArgs(chosenCmd executor.Command) []string {
	var moreArgs []string
	for _, argName := range chosenCmd.UserInput() {
		if argName != "" {
			moreArgs = append(moreArgs, argName)
		}
		input, err := clip.ReadAll()
		if err != nil {
			log.Fatal(err.Error())
		}
		moreArgs = append(moreArgs, input)
	}
	return moreArgs
}

func callRofi(rofiMenu string) []byte {
	rofi := exec.Command("rofi", "-dmenu", "-format", "s", "-p", "tools")

	stdin, err := rofi.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}
	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, rofiMenu)
	}()

	output, err := rofi.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
	}
	return output

}

func parseChosenCmd(output []byte, list []executor.Command) executor.Command {
	strOutput := string(output)
	for _, c := range list {
		fmt.Println(c.Name())
		fmt.Println(strOutput)
		if strings.EqualFold(strings.TrimSuffix(strOutput, "\n"), c.Name()) {
			return c
		}
	}
	return nil
}

func buildRofiFromCmds(cmdList []executor.Command) string {
	rofiMenu := strings.Builder{}

	for _, c := range cmdList {
		rofiMenu.WriteString(c.Name())
		rofiMenu.WriteString("\n")
	}
	return rofiMenu.String()
}
