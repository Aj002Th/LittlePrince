package logging

import (
	"io"
	"log"
	"os"

	"github.com/sirupsen/logrus"
)

var Log = logrus.New()

func init() {
	logFile, err := GetLogFile("router_")
	if err != nil {
		log.Fatalf("logging.Setup err: %v", err)
	}

	Log.SetOutput(io.MultiWriter(logFile, os.Stdout))
	Log.SetReportCaller(true)
	Log.SetFormatter(&logrus.TextFormatter{})
}

// Debug output logs at debug level
func Debug(v ...interface{}) {
	Log.Debug(v...)
}

// Info output logs at info level
func Info(v ...interface{}) {
	Log.Info(v...)
}

// Warn output logs at warn level
func Warn(v ...interface{}) {
	Log.Warn(v...)
}

// Error output logs at error level
func Error(v ...interface{}) {
	Log.Error(v...)
}

// Fatal output logs at fatal level
func Fatal(v ...interface{}) {
	Log.Fatal(v...)
}
