package logger

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

var log *logrus.Logger

func initLog() *logrus.Logger {
	pathMap := lfshook.PathMap{
		logrus.DebugLevel: "./log/debug.log",
		logrus.InfoLevel:  "./log/info.log",
		logrus.WarnLevel:  "./log/warning.log",
		logrus.ErrorLevel: "./log/error.log",
	}

	log = logrus.New()
	log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	log.Info("logger.initLog: init log success")
	return log
}

func GetLog() *logrus.Logger {
	if log == nil {
		return initLog()
	}
	return log
}
