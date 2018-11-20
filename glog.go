package glog

import (
	"fmt"
	"os"
	"sync"

	"github.com/go-kit/kit/log"
	"github.com/go-kit/kit/log/level"
)

var logger = log.NewNopLogger()
var mu = sync.Mutex{}

func SetLogger(l log.Logger) {
	mu.Lock()
	logger = l
	mu.Unlock()
}

type Level int32

type Verbose bool

func V(level Level) Verbose { return true }

func logMsg(l log.Logger, f string, msg string) {
	l.Log("func", f, "msg", msg)
}

func (v Verbose) Info(args ...interface{}) {
	logMsg(level.Debug(logger), "Verbose.Info", fmt.Sprint(args...))
}

func (v Verbose) Infoln(args ...interface{}) {
	logMsg(level.Debug(logger), "Verbose.Infoln", fmt.Sprint(args...))
}

func (v Verbose) Infof(format string, args ...interface{}) {
	logMsg(level.Debug(logger), "Verbose.Infof", fmt.Sprintf(format, args...))
}

func Info(args ...interface{}) {
	logMsg(level.Debug(logger), "Info", fmt.Sprint(args...))
}

func InfoDepth(depth int, args ...interface{}) {
	logMsg(level.Debug(logger), "InfoDepth", fmt.Sprint(args...))
}

func Infoln(args ...interface{}) {
	logMsg(level.Debug(logger), "Infoln", fmt.Sprint(args...))
}

func Infof(format string, args ...interface{}) {
	logMsg(level.Debug(logger), "Infof", fmt.Sprintf(format, args...))
}

func Warning(args ...interface{}) {
	logMsg(level.Warn(logger), "Warning", fmt.Sprint(args...))
}

func WarningDepth(depth int, args ...interface{}) {
	logMsg(level.Warn(logger), "WarningDepth", fmt.Sprint(args...))
}

func Warningln(args ...interface{}) {
	logMsg(level.Warn(logger), "Warningln", fmt.Sprint(args...))
}

func Warningf(format string, args ...interface{}) {
	logMsg(level.Warn(logger), "Warningf", fmt.Sprintf(format, args...))
}

func Error(args ...interface{}) {
	logMsg(level.Error(logger), "Error", fmt.Sprint(args...))
}

func ErrorDepth(depth int, args ...interface{}) {
	logMsg(level.Error(logger), "ErrorDepth", fmt.Sprint(args...))
}

func Errorln(args ...interface{}) {
	logMsg(level.Error(logger), "Errorln", fmt.Sprint(args...))
}

func Errorf(format string, args ...interface{}) {
	logMsg(level.Error(logger), "Errorf", fmt.Sprintf(format, args...))
}

func Fatal(args ...interface{}) {
	logMsg(level.Error(logger), "Fatal", fmt.Sprint(args...))
	os.Exit(255)
}

func FatalDepth(depth int, args ...interface{}) {
	logMsg(level.Error(logger), "FatalDepth", fmt.Sprint(args...))
	os.Exit(255)
}

func Fatalln(args ...interface{}) {
	logMsg(level.Error(logger), "Fatalln", fmt.Sprint(args...))
	os.Exit(255)
}

func Fatalf(format string, args ...interface{}) {
	logMsg(level.Error(logger), "Fatalf", fmt.Sprintf(format, args...))
	os.Exit(255)
}

func Exit(args ...interface{}) {
	logMsg(level.Error(logger), "Exit", fmt.Sprint(args...))
	os.Exit(1)
}

func ExitDepth(depth int, args ...interface{}) {
	logMsg(level.Error(logger), "ExitDepth", fmt.Sprint(args...))
	os.Exit(1)
}

func Exitln(args ...interface{}) {
	logMsg(level.Error(logger), "ExitDepth", fmt.Sprint(args...))
	os.Exit(1)
}

func Exitf(format string, args ...interface{}) {
	logMsg(level.Error(logger), "Exit", fmt.Sprintf(format, args...))
	os.Exit(1)
}
