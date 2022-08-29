package timeUtil

import (
	"time"

	"github.com/PanDa30ss/core/service"
)

func MakeTimer(d time.Duration, f func(params ...interface{}), params ...interface{}) *time.Timer {
	return time.AfterFunc(d, func() {
		service.Post(service.MakeCommand(service.ExecuteFunc(f), params...))
	})
}
