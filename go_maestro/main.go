package main

import (
	"bytes"
	"fmt"
	"github.com/alexpfx/golang/go_maestro/internal/commands"
	"github.com/alexpfx/golang/go_maestro/internal/output"
	clip "github.com/atotto/clipboard"
	"io"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	cmds := []commands.Cmd{
		commands.NewMassaCnisHomCat8(),
		commands.MergeFetch(),
		commands.MassaListaCatalogos(),
	}

	output := callRofi(buildRofiFromCmds(cmds))

	chosenCmd := parseChosenCmd(output, cmds)

	if chosenCmd == nil {
		return
	}

	var ua []string
	if len(chosenCmd.UserInput) != 0 {
		ua = appendUserArgs(chosenCmd)
		fmt.Println(ua)
	}
	callCmd(chosenCmd, ua)

}

func callCmd(cmd *commands.Cmd, ua []string) {
	args := append(cmd.Args, ua...)

	command := exec.Command(cmd.Binary, args...)

	command.Env = os.Environ()

	var stdOut, stdErr bytes.Buffer

	command.Stdout = &stdOut
	command.Stderr = &stdErr

	err := command.Run()
	outStr, errStr := string(stdOut.Bytes()), string(stdErr.Bytes())
	if err != nil {
		callRofiMessage(cmd.Binary, errStr)
		log.Fatal(err.Error())
	}

	if outStr != "" {
		var out string
		if len(cmd.FilterOutput) != 0 {
			out = output.Filter(outStr, cmd.FilterOutput)
		} else {
			out = outStr
		}

		if cmd.Clipboard {
			callRofiMessage(cmd.Binary, out)
			clip.WriteAll(out)
		} else {
			callRofi(out)
		}
	}

}

func appendUserArgs(chosenCmd *commands.Cmd) []string {
	var moreArgs []string
	for _, argName := range chosenCmd.UserInput {
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

func callRofiMessage(title, msg string) {
	rofi := exec.Command("rofi", "-e", fmt.Sprintf("%s:\n\n%s", title, msg))
	_ = rofi.Run()
}

func callRofi(rofiMenu string) int {
	rofi := exec.Command("rofi", "-i", "-cache-dir", "./tmp","-dmenu", "-format", "i", "-p", "tools")

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
	println("output", string(output))
	atoi, _ := strconv.Atoi(strings.TrimRight(string(output), "\n"))
	return atoi
}

func parseChosenCmd(index int, list []commands.Cmd) *commands.Cmd {
	return &list[index]
}

func buildRofiFromCmds(cmdList []commands.Cmd) string {
	rofiMenu := strings.Builder{}

	max := 0
	for _, cmd := range cmdList {
		nLen := len(cmd.Name)
		if nLen > max {
			max = nLen
		}
	}

	for _, c := range cmdList {
		rofiMenu.WriteString(fmt.Sprintf("%-*s", max, c.Name))
		rofiMenu.WriteString("\t")
		rofiMenu.WriteString(c.Desc)

		rofiMenu.WriteString("\n")
	}
	return rofiMenu.String()
}
