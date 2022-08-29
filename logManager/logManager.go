package logManager

import (
	"os"
	"sync"
	"time"

	log "github.com/sirupsen/logrus"
)

type logManager struct {
	timer  *time.Timer
	logger *log.Logger
}

var instance *logManager
var once sync.Once

func getInstance() *logManager {
	once.Do(func() {
		instance = makeLogManager()
	})
	return instance
}

//var register = service.GetInstance().RegisterModule(getInstance(), false)

func makeLogManager() *logManager {
	ret := &logManager{}
	ret.logger = log.New()
	ret.setFile()
	go ret.setTimer()
	return ret
}

func (this *logManager) setFile() {

	fileName := "./" + "log_" + time.Now().Format("20060102") + ".log"
	logFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0766)
	if err != nil {
		panic(err)
	}
	this.logger.SetOutput(logFile)

}

func (this *logManager) setTimer() {
	now := time.Now()
	nextTime := time.Date(now.Year(), now.Month(), now.Day()+1, 0, 0, 0, 0, now.Location())
	this.timer = time.NewTimer(nextTime.Sub(now))
	<-this.timer.C
	this.setFile()
	go this.setTimer()
}
