package main

import (
	"bytes"
	"fmt"
	"github.com/alexpfx/go_common/cmd"
	rofi2 "github.com/alexpfx/go_common/rofi"
	"github.com/alexpfx/golang/go_maestro/internal/commands"
	"github.com/alexpfx/golang/go_maestro/internal/output"
	clip "github.com/atotto/clipboard"
	"log"
	"os"
	"os/exec"
	"strconv"
	"strings"
)

func main() {

	newCmds := []cmd.Cmd{
		commands.NewMergeFetch(),
	}

	menu := buildMenu(newCmds)
	selected := callRofiMenu(menu, "i")
	index, _ := strconv.Atoi(strings.TrimRight(selected, "\n"))
	xc := newCmds[index]

	res, err := cmd.Run(&xc)

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	//exception.CheckThrow(err)

	fmt.Printf("resposta: %v", string(res))
	/*
		cmds := []commands.Cmd{
			commands.NewMassaCnisHomCat8(),
			commands.MergeFetch(),'
			commands.MassaListaCatalogos(),
			commands.SibeSibeDeploy(),
			commands.SibeSibeClient(),
			commands.QuickFixQuery(),
		}

		rofiOutput := callRofiMenu(buildRofiMenu(cmds), "i")
		//index, _ := strconv.Atoi(strings.TrimRight(rofiOutput, "\n"))
		chosenCmd := parseChosenCmd(index, cmds)

		if chosenCmd == nil {
			return
		}

		var ua []string
		if len(chosenCmd.UserInput) != 0 {
			ua = appendUserArgs(chosenCmd)
		}
		callCmd(chosenCmd, ua)
	*/
}

func buildMenu(cmds []cmd.Cmd) string {
	var sb strings.Builder
	for i, c := range cmds {
		sb.WriteString(fmt.Sprintf("%d - %s: %s", i+1, c.Binary.Name, c.Binary.Desc))
	}
	return sb.String()
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
		afterFormatStr, hasFormat := tryFormat(cmd, afterFilterStr)
		if hasFormat {
			if cmd.CallNext != nil {
				if cmd.OutputConverter != nil {
					callRofiWithCmd(afterFormatStr, cmd.OutputConverter, cmd)
				}
			} else {
				//callRofiMenu(afterFormatStr, "i")
			}
		} else {
			//callRofiMessage(cmd.Binary, afterFormatStr)
		}

		if cmd.CopyOutput {
			err = clip.WriteAll(afterFormatStr)
			if err != nil {
				log.Fatalf(err.Error())
			}
		}
	}

}

func tryFormat(cmd *commands.Cmd, str string) (string, bool) {
	fmtOut := cmd.FormatOutput
	if cmd.DynamicFormatOutput != nil && len(fmtOut) > 0 {
		log.Fatal("argumento inválido: forneça apenas um dos seguintes parâmetros: cmd.DynamicFormatOutput, cmd.FormatOutput")
	}
	if len(fmtOut) > 0 {
		return output.Format(str, fmtOut), true
	}

	if cmd.DynamicFormatOutput == nil {
		return str, false
	}

	fmtOut = cmd.DynamicFormatOutput(str)
	return output.Format(str, fmtOut), true

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
	rofi2.NewMessage(fmt.Sprintf("%s:\n\n%s", title, msg))
}

func callRofiWithCmd(rofiMenu string, converter func(string) (string, []string), cmd *commands.Cmd) {
	out := callRofiMenu(rofiMenu, "s")

	_, args := converter(out)
	nextCmd := cmd.CallNext(args...)
	callCmd(nextCmd, []string{})
}

func callRofiMenu(rofiMenu string, format string) string {
	dMenu := rofi2.NewDMemuBuilder().Format(format).Prompt("selecione").Build()
	out, _ := dMenu.Exec(rofiMenu)
	return out
}

func parseChosenCmd(index int, list []commands.Cmd) *commands.Cmd {
	return &list[index]
}

func buildRofiMenu(cmdList []commands.Cmd) string {
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
