package logger

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"os"
)

var Log = logrus.New()

func init() {
	Log.Out = os.Stdout
	Log.SetLevel(logrus.InfoLevel)
	Log.SetFormatter(&logrus.TextFormatter{
		FullTimestamp: true,
	})
}

func Info(msg string, args ...interface{}) {
	Log.Info(fmt.Sprintf(msg, args...))
}

func Warn(msg string, args ...interface{}) {
	Log.Warn(fmt.Sprintf(msg, args...))
}

func Error(msg string, args ...interface{}) {
	Log.Error(fmt.Sprintf(msg, args...))
}

func Debug(msg string, args ...interface{}) {
	if Log.GetLevel() >= logrus.DebugLevel {
		Log.Debug(fmt.Sprintf(msg, args...))
	}
}

func Fatal(msg string, args ...interface{}) {
	Log.Fatal(fmt.Sprintf(msg, args...))
}

func Panic(msg string, args ...interface{}) {
	Log.Panic(fmt.Sprintf(msg, args...))
}
