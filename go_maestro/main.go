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
		commands.SibeSibeDeploy(),
		commands.SibeSibeClient(),
	}

	rofiOutput := callRofi(buildRofiFromCmds(cmds), "i")
	index, _ := strconv.Atoi(strings.TrimRight(rofiOutput, "\n"))
	chosenCmd := parseChosenCmd(index, cmds)

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
		afterFilterStr := tryFilter(cmd.FilterOutput, outStr)
		afterFormatStr, hasFormat := tryFormat(cmd.FormatOutput, afterFilterStr)
		if hasFormat {
			if cmd.CallNext != nil {
				if cmd.OutputConverter != nil {
					callRofiWithCmd(afterFormatStr, cmd.OutputConverter, cmd)
				}
			}else{
				callRofi(afterFormatStr, "i")
			}
		} else {
			callRofiMessage(cmd.Binary, afterFormatStr)
		}

		if cmd.CopyOutput {
			_ = clip.WriteAll(afterFormatStr)
		}
	}

}

func tryFormat(formatOutput []string, str string) (string, bool) {
	if len(formatOutput) == 0 {
		return str, false
	}
	return output.Format(str, formatOutput), true
}

func tryFilter(filter []string, str string) string {
	if len(filter) == 0 {
		return str
	}
	return output.Filter(str, filter)

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

func callRofiWithCmd(rofiMenu string, converter func(string) (string, []string), cmd *commands.Cmd) {
	out := callRofi(rofiMenu, "s")

	_, args := converter(out)
	nextCmd := cmd.CallNext(args...)
	callCmd(nextCmd, []string{})
}

func callRofi(rofiMenu string, format string) string {
	rofi := exec.Command("rofi", "-i", "-dmenu", "-p", "selecione", "-format", format)


	stdin, err := rofi.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		_, _ = io.WriteString(stdin, rofiMenu)
	}()

	out, err := rofi.CombinedOutput()
	if err != nil {
		log.Fatal(err.Error())
	}
	return string(out)

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
