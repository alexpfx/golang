package main

import (
	"flag"
	"fmt"
	"github.com/alexpfx/golang/go_timer/internal/timer"
	"log"
	"os"
	"time"
)

func main() {
	startCmd := flag.NewFlagSet("start", flag.ExitOnError)
	var time time.Duration
	startCmd.DurationVar(&time, "t", 1, `-t ex: ""`)

	args := os.Args
	if len(args) < 2 {
		log.Fatal("número de parâmetros incorreto")
	}

	switch args[1] {
	case "start":
		startCmd.Parse(args[2:])
		startCountdown(time)
	}

}

func startCountdown(pTime time.Duration) {
	fmt.Print(pTime)

	if alreadyStarted() {
		log.Fatal("timer already started")
	}

	countdown := timer.NewCountdown()
	countdown.Start(pTime)
}

func alreadyStarted() bool {
	return false
}
