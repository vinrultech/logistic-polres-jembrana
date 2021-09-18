package loggers

import (
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
)

//Log untuk logrus
var Log *logrus.Logger

//NewLogger fungsi untuk logrus
func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  "info.log",
		logrus.ErrorLevel: "error.log",
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}
