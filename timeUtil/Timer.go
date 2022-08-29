package timeUtil

import (
	"core/service"
	"time"
)

func MakeTimer(d time.Duration, f func(params ...interface{}), params ...interface{}) *time.Timer {
	return time.AfterFunc(d, func() {
		service.Post(service.MakeCommand(service.ExecuteFunc(f), params...))
	})
}
