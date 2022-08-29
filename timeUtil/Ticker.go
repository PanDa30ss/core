package timeUtil

import (
	"time"

	"github.com/PanDa30ss/core/service"
)

type Ticker struct {
	t    *time.Ticker
	done chan bool
}

func MakeTicker(d time.Duration, f func(params ...interface{}), params ...interface{}) *Ticker {
	ticker := &Ticker{}
	ticker.t = time.NewTicker(d)
	ticker.done = make(chan bool, 1)
	go func(ticker *Ticker) {
		for {
			select {
			case <-ticker.done:
				return
			case <-ticker.t.C:
				service.Post(service.MakeCommand(service.ExecuteFunc(f), params...))
			}
		}
	}(ticker)
	return ticker
}

func (this *Ticker) Stop() {
	this.done <- true
	this.t.Stop()
}
