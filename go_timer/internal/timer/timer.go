package timer

import (
	"fmt"
	"log"
	"time"
)

type State int

const (
	Stopped State = iota
	Running
)


type Timer interface {
	Start(duration time.Duration)
	Stop()
	Status() (time.Duration, State)
}

func NewCountdown() Timer {
	return countdown{
	}
}

type countdown struct {
	ticker *time.Ticker
	actual time.Duration
	total time.Duration
	state State
}

func (c countdown) Start(duration time.Duration) {
	if c.state == Running{
		log.Fatal("existe um timer em andamento")
	}
	c.total = duration
	c.state = Running
	c.ticker = time.NewTicker(1 * time.Second)
	c.actual = duration
	for {
		msg := <- c.ticker.C
		c.actual = msg.Sub(time.Now())
		fmt.Println(msg)
		if c.actual < 0 {
			break
		}
	}

}



func (c countdown) Stop() {
	c.ticker.Reset(c.total)
	c.ticker.Stop()
	c.state = Stopped
}

func (c countdown) Status() (time.Duration, State) {
	return c.actual, c.state
}



