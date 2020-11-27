package main

import (
	"fmt"
	"io"
	"log"
	"os/exec"
	"strings"
)

func main() {
	list := [][]string{
		{"merge", ""},
	}

	rofi := exec.Command("rofi", "-dmenu", "-format", "s", "-p", "tools")

	stdin, err := rofi.StdinPipe()
	if err != nil {
		log.Fatal(err)
	}

	go func() {
		defer stdin.Close()
		for _, s := range list {
			_, _ = io.WriteString(stdin, strings.Join([]string{s[0]}, "\n"))
		}
	}()

	output, err := rofi.CombinedOutput()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%s\n", output)

}
