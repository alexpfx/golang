package script

import (
	"bytes"
	"fmt"
	"log"
	"os"
	"os/exec"
)

type Item struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Script struct {
	CmdPath string
	RunPath string
}

type Runner interface {
	Run(script Script, args ...string) (string, error)
}

func (c Script) Run(script Script, args ...string) (string, error) {
	cmd := exec.Cmd{Path: c.CmdPath, Dir: c.RunPath, Args: args}
	cmd.Env = os.Environ()
	var stdErr, stdOut bytes.Buffer
	cmd.Stdout = &stdOut
	cmd.Stderr = &stdErr

	err := cmd.Run()
	outStr, errStr := string(stdOut.Bytes()), string(stdErr.Bytes())
	if err != nil {
		log.Fatal(err.Error())
	}
	fmt.Printf(outStr)
	fmt.Println(errStr)

	return "", nil
}

var ClientsItems = []Item{
	{
		Name: "All",
		Id:   0,
	},
	{
		Name: "SIBE_COMMONS",
		Id:   1,
	},
	{
		Name: "mcb",
		Id:   16,
	},
	{
		Name: "sisben",
		Id:   21,
	},
	{
		Name: "sibe hiscre",
		Id:   23,
	},
	{
		Name: "tcb",
		Id:   26,
	},
	{
		Name: "negocio",
		Id:   27,
	},
	{
		Name: "broker",
		Id:   28,
	},
	{
		Name: "PortalSibe",
		Id:   33,
	},
}

var DeployItems = []Item{
	{
		Name: "All",
		Id:   0,
	},
	{
		Name: "LIBS-SIBE",
		Id:   1,
	}, {
		Name: "sibe cache",
		Id:   2,
	},
	{
		Name: "atuben",
		Id:   4,
	},
	{
		Name: "camvri",
		Id:   5,
	},
	{
		Name: "comsub",
		Id:   7,
	},
	{
		Name: "MCB",
		Id:   16,
	},
	{
		Name: "migra",
		Id:   18,
	},
	{
		Name: "reavdir",
		Id:   20,
	},
	{
		Name: "sisben",
		Id:   21,
	},
	{
		Name: "hiscre",
		Id:   23,
	},
	{
		Name: "tcb",
		Id:   26,
	},
	{
		Name: "negocio",
		Id:   27,
	},
	{
		Name: "broker",
		Id:   28,
	},
	{
		Name: "conev processamento",
		Id:   29,
	},
	{
		Name: "PortalSibe",
		Id:   33,
	},
	{
		Name: "PortalPericia",
		Id:   34,
	},
	{
		Name: "PortalSibeExploded",
		Id:   333,
	},
	{
		Name: "webservices",
		Id:   40,
	},
}
