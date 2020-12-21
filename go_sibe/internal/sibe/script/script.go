package script

import (
	"bufio"
	"log"
	"os"
	"os/exec"
	"strings"
)

type Option struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

type Script struct {
	Path string
	Cmd  string
}

type Runner interface {
	Run(args []string, ch chan string) error
}

type runner struct {
	cmd string
}

func SibeClient() Script {
	return Script{
		Path: os.Getenv("SIBE_DIR") + "/scripts",
		Cmd:  "sibeClient.sh",
	}
}

func SibeDeploy() Script {
	return Script{
		Path: os.Getenv("SIBE_DIR") + "/scripts",
		Cmd:  "sibeDeploy.sh",
	}
}

func NewRunner(script Script) Runner {
	return runner{
		cmd: strings.Join([]string{script.Path, script.Cmd}, "/"),
	}
}

func (c runner) Run(args []string, ch chan string) error {
	cmd := exec.Command(c.cmd, args...)
	pipe, err := cmd.StdoutPipe()

	defer close(ch)

	err = cmd.Start()
	if err != nil {
		log.Fatal((err))
	}

	scanner := bufio.NewScanner(pipe)

	ch <- "teste"
	go func() {
		for scanner.Scan() {
			ch <- scanner.Text()
		}
	}()

	cmd.Wait()
	return nil

}

var ClientScripts = []Option{
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

var DeployScripts = []Option{
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
