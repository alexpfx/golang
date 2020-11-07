package go_task_test

import (
	"fmt"
	"github.com/alexpfx/golang/go_task"
	"testing"
)

func TestConnectCookies(t *testing.T) {

	ccmSession := go_task.NewCcmSession("", "", "")

	r := ccmSession.Get("20000")

	fmt.Println(r)

}


