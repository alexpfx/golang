package bookmark

import (
	"fmt"
	"log"
	"os"
)

func check(err error, msg string) {
	if err != nil {
		fmt.Println(msg)
	}
}

func checkPrint(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(0)
	}
}
func checkPanic(err error) {
	if err != nil {
		log.Panic(err)
	}
}
