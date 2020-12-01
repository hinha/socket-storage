package log

import (
	"fmt"
	"github.com/rifflock/lfshook"
	"github.com/sirupsen/logrus"
	"os"
	"time"
)

var Log *logrus.Logger

func NewLogger() *logrus.Logger {
	if Log != nil {
		return Log
	}

	basePath, _ := os.Getwd()
	path := fmt.Sprintf("%s/temp/log", basePath)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		err := os.Mkdir(path, os.ModePerm)
		if err != nil {
			panic(err)
		}
	}
	y, m, _ := time.Now().Date()
	pathMap := lfshook.PathMap{
		logrus.InfoLevel:  fmt.Sprintf("%s/%d%d_%s", path, int(m), y, "info.log"),
		logrus.ErrorLevel: fmt.Sprintf("%s/%d%d_%s", path, int(m), y, "error.log"),
	}

	Log = logrus.New()
	Log.Hooks.Add(lfshook.NewHook(
		pathMap,
		&logrus.JSONFormatter{},
	))
	return Log
}
