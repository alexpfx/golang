package main

import (
	"bytes"
	"flag"
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

var isHelp bool
var command string

func main() {
	flag.BoolVar(&isHelp, "help", false, "imprimir a ajuda")
	flag.StringVar(&command, "c", "", "especifica o comando a ser executado, em vez de montar o menu com todas as opções disponíveis")

	flag.Parse()
	if isHelp {
		flag.PrintDefaults()
		os.Exit(0)
	}

	newCmds := append(commands.FocusMonitor(), commands.NewGoMassaCustomInput(),
		commands.NewMergeFetch())

	var menu string
	var selected string
	if command == "" {
		menu = buildMenu(newCmds)
		selected = callRofiMenu(menu, "i")
	} else {
		selected = findByIdentifier(command, newCmds)
		if selected == "" {
			log.Fatalf("comando não encontrado: %s", command)
		}
	}

	index, _ := strconv.Atoi(strings.TrimRight(selected, "\n"))
	xc := &newCmds[index]

	res, err := cmd.Run(xc)
	if err != nil {
		fmt.Printf("%T\n", err.Error())
		fmt.Printf("%s\n", err.Error())
		//callRofiMessage(err.Error())
		return
	}
	//exception.CheckThrow(err)

	fmt.Printf("result: %s", string(res))
}

func findByIdentifier(c string, cmds []cmd.Cmd) string {

	for i, cmd := range cmds {
		fmt.Println(cmd.Identifier)
		if strings.EqualFold(cmd.Identifier, c) {
			return strconv.Itoa(i)
		}
	}
	return ""

}

func buildMenu(cmds []cmd.Cmd) string {
	var sb strings.Builder
	for i, c := range cmds {
		formatName := "%d - %s\n"
		formatDesc := "%d - %s: %s\n"

		if c.Binary.Desc != "" {
			sb.WriteString(fmt.Sprintf(formatName, i+1, c.Binary.Name))
		} else {
			sb.WriteString(fmt.Sprintf(formatDesc, i+1, c.Binary.Name, c.Binary.Desc))
		}

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
		callRofiMessage(errStr)
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

func callRofiMessage(msg string) {
	message := rofi2.NewMessage(fmt.Sprintf("%s:\n", msg))
	message.Exec(msg)

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
