package logManager

import log "github.com/sirupsen/logrus"

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
