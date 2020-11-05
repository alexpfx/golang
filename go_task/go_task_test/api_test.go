package go_task_test

import (
	"github.com/alexpfx/golang/go_task"
	"testing"
)

func TestConnectCookies(t *testing.T) {

	ccmSession := go_task.NewCcmSession("alexandre.alessi", "", "alm.dataprev.gov.br")

	ccmSession.Connect()

}


