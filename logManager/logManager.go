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

func SetLevel(v string) {
	var level log.Level = log.TraceLevel
	if v == "release" {
		level = log.InfoLevel
	}
	getInstance().logger.SetLevel(level)
	log.SetLevel(level)
}

func Info(args ...interface{}) {
	log.Info(args)
	getInstance().logger.Info(args)
}

func Warning(args ...interface{}) {
	log.Warning(args)
	getInstance().logger.Warning(args)

}

func Error(args ...interface{}) {
	log.Error(args)
	getInstance().logger.Error(args)

}

func Debug(args ...interface{}) {
	log.Debug(args)
	getInstance().logger.Debug(args)

}

func Infof(format string, args ...interface{}) {
	log.Infof(format, args...)
	getInstance().logger.Infof(format, args...)
}

func Warningf(format string, args ...interface{}) {
	log.Warningf(format, args...)
	getInstance().logger.Warningf(format, args...)
}

func Errorf(format string, args ...interface{}) {
	log.Errorf(format, args...)
	getInstance().logger.Errorf(format, args...)
}

func Debugf(format string, args ...interface{}) {
	log.Debugf(format, args...)
	getInstance().logger.Debugf(format, args...)
}

func Infoln(args ...interface{}) {
	log.Infoln(args)
	getInstance().logger.Infoln(args)
}

func Warningln(args ...interface{}) {
	log.Warningln(args)
	getInstance().logger.Warningln(args)

}

func Errorln(args ...interface{}) {
	log.Errorln(args)
	getInstance().logger.Errorln(args)

}

func Debugln(args ...interface{}) {
	log.Debugln(args)
	getInstance().logger.Debugln(args)

}

func InfoFn(fn log.LogFunction) {
	log.InfoFn(fn)
	getInstance().logger.InfoFn(fn)
}

func WarningFn(fn log.LogFunction) {
	log.WarningFn(fn)
	getInstance().logger.WarningFn(fn)
}

func ErrorFn(fn log.LogFunction) {
	log.ErrorFn(fn)
	getInstance().logger.ErrorFn(fn)
}

func DebugFn(fn log.LogFunction) {
	log.DebugFn(fn)
	getInstance().logger.DebugFn(fn)
}
