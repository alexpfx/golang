package main

import (
	"fmt"
	"github.com/alexpfx/go_action/action"
	"github.com/alexpfx/golang/go_maestro/internal/trees"
	"github.com/urfave/cli/v2"
	"log"
	
	"os"
)

func main() {
	app := &cli.App{
		Name: "go maestro",
		Commands: []*cli.Command{
			{
				Name:    "tree",
				Aliases: []string{},
				Usage: "monta uma árvore de menu de nome passado como parâmetro. " +
					"para ver as opções disponíveis usar o comando o argumento list",
				ArgsUsage: "<tree_name>",
				Flags: []cli.Flag{
					&cli.StringFlag{
						Name:    "menu",
						Aliases: []string{"m"},
						Usage:   "especifica o comando usado para montar o menu [fzf|rofi]",
						Value:   "fzf",
					},
				},
				Action: func(ctx *cli.Context) error {
					if ctx.NArg() < 1 {
						_ = cli.ShowAppHelp(ctx)
						return nil
					}
					var tree action.Tree
					tree = searchTree(ctx.Args().First())
					selectedAction, found, err := tree.Show()
					if err != nil {
						return err
					}
					if !found {
						return fmt.Errorf("selecione uma action")
					}
					
					err = log.Output(1, fmt.Sprintf("executando action: %s", selectedAction.Name))
					if err != nil {
						return err
					}
					actRes, err := selectedAction.Run()
					if err != nil {
						return err
					}
					
					fmt.Println(string(actRes))
					
					return nil
				},
			},
		},
	}
	
	err := app.Run(os.Args)
	if err != nil {
		_, _ = fmt.Fprintf(os.Stderr, "%s!\n", err.Error())
		os.Exit(1)
	}
	
}

func searchTree(treeName string) action.Tree {
	return trees.DtpTree
}

//func findByIdentifier(c string, cmds []action.Action) string {
//
//	for i, cmd := range cmds {
//		fmt.Println(cmd.Identifier)
//		if strings.EqualFold(cmd.Identifier, c) {
//			return strconv.Itoa(i)
//		}
//	}
//	return ""
//
//}
//
//func buildMenu(cmds []action.Action) string {
//	var sb strings.Builder
//	for i, c := range cmds {
//		formatName := "%d - %s\n"
//		formatDesc := "%d - %s: %s\n"
//
//		if c.Binary.Desc != "" {
//			sb.WriteString(fmt.Sprintf(formatName, i+1, c.Binary.Name))
//		} else {
//			sb.WriteString(fmt.Sprintf(formatDesc, i+1, c.Binary.Name, c.Binary.Desc))
//		}
//
//	}
//	return sb.String()
//}
//
///*func callCmd(cmd *cmd.Action, ua []string) {
//	args := append(cmd.Args, ua...)
//
//	command := exec.Command(cmd.Binary.CmdPath, args...)
//
//	command.Env = os.Environ()
//
//	var stdOut, stdErr bytes.Buffer
//
//	command.Stdout = &stdOut
//	command.Stderr = &stdErr
//
//	err := command.Run()
//	outStr, errStr := string(stdOut.Bytes()), string(stdErr.Bytes())
//	if err != nil {
//		callRofiMessage(errStr)
//		log.Fatal(err.Error())
//	}
//
//	if outStr != "" {
//		afterFilterStr := tryFilter(cmd.FilterOutput, outStr)
//		afterFormatStr, hasFormat := tryFormat(cmd, afterFilterStr)
//		if hasFormat {
//			if cmd.CallNext != nil {
//				if cmd.OutputConverter != nil {
//					callRofiWithCmd(afterFormatStr, cmd.OutputConverter, cmd)
//				}
//			} else {
//				//callRofiMenu(afterFormatStr, "i")
//			}
//		} else {
//			//callRofiMessage(cmd.Binary, afterFormatStr)
//		}
//
//		if cmd.CopyOutput {
//			err = clip.WriteAll(afterFormatStr)
//			if err != nil {
//				log.Fatalf(err.Error())
//			}
//		}
//	}
//
//}*/
//
//func tryFormat(cmd *action.Action, str string) (string, bool) {
//	fmtOut := cmd.FormatOutput
//	if cmd.DynamicFormatOutput != nil && len(fmtOut) > 0 {
//		log.Fatal("argumento inválido: forneça apenas um dos seguintes parâmetros: cmd.DynamicFormatOutput, cmd.FormatOutput")
//	}
//	if len(fmtOut) > 0 {
//		return output.Format(str, fmtOut), true
//	}
//
//	if cmd.DynamicFormatOutput == nil {
//		return str, false
//	}
//
//	fmtOut = cmd.DynamicFormatOutput(str)
//	return output.Format(str, fmtOut), true
//
//}
//
//func tryFilter(filter []string, str string) string {
//	if len(filter) == 0 {
//		return str
//	}
//	return output.Filter(str, filter)
//
//}
//
//func appendUserArgs(chosenCmd *action.Action) []string {
//	var moreArgs []string
//	for _, argName := range chosenCmd.InputConfig {
//		if argName != "" {
//			moreArgs = append(moreArgs, argName)
//		}
//		input, err := clip.ReadAll()
//		if err != nil {
//			log.Fatal(err.Error())
//		}
//		moreArgs = append(moreArgs, input)
//	}
//	return moreArgs
//}
//
///*func callRofiWithCmd(rofiMenu string, converter func(string) (string, []string), cmd *action.Action) {
//	out := callRofiMenu(rofiMenu, "s", false)
//
//	_, args := converter(out)
//	nextCmd := cmd.CallNext(args...)
//	callCmd(nextCmd, []string{})
//}*/
//
//func callRofiMenu(rofiMenu string, format string, isFzf bool) string {
//	if isFzf {
//		menu := fzf.NewBuilder().Prompt("selecione:\n").Build()
//		res, _ := menu.Run(rofiMenu)
//		return res
//	}
//
//	dMenu := rofi2.NewDMenuBuilder().Format(format).Prompt("selecione").Build()
//	out, _ := dMenu.Exec(rofiMenu)
//	return out
//}
//
//func callRofiMessage(msg string) {
//	message := rofi2.NewMessage(fmt.Sprintf("%s:\n", msg))
//	message.Exec(msg)
//
//}
//
//func parseChosenCmd(index int, list []action.Action) *action.Action {
//	return &list[index]
//}
//
//func buildRofiMenu(cmdList []action.Action) string {
//	rofiMenu := strings.Builder{}
//
//	max := 0
//	for _, cmd := range cmdList {
//		nLen := len(cmd.Name)
//		if nLen > max {
//			max = nLen
//		}
//	}
//
//	for _, c := range cmdList {
//		rofiMenu.WriteString(fmt.Sprintf("%-*s", max, c.Name))
//		rofiMenu.WriteString("\t")
//		rofiMenu.WriteString(c.Desc)
//
//		rofiMenu.WriteString("\n")
//	}
//	return rofiMenu.String()
//}
