package glog

import (
	"fmt"
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

func (v Verbose) Info(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Verbose.Info")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Debug(logger).Log(pairs...)
}

func (v Verbose) Infoln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Verbose.Infoln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Debug(logger).Log(pairs...)
}

func (v Verbose) Infof(format string, args ...interface{}) {
	level.Debug(logger).Log(
		"func", "Infof",
		"msg", fmt.Sprintf(format, args...),
	)
}

func Info(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Info")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Debug(logger).Log(pairs...)
}

func InfoDepth(depth int, args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "InfoDepth")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Debug(logger).Log(pairs...)
}

func Infoln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Infoln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Debug(logger).Log(pairs...)
}

func Infof(format string, args ...interface{}) {
	level.Debug(logger).Log(
		"func", "Infof",
		"msg", fmt.Sprintf(format, args...),
	)
}

func Warning(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Warning")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Warn(logger).Log(pairs...)
}

func WarningDepth(depth int, args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "WarningDepth")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Warn(logger).Log(pairs...)
}

func Warningln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Warningln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Warn(logger).Log(pairs...)
}

func Warningf(format string, args ...interface{}) {
	level.Warn(logger).Log(
		"func", "Warningf",
		"msg", fmt.Sprintf(format, args...),
	)
}

func Error(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Error")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func ErrorDepth(depth int, args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "ErrorDepth")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Errorln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Errorln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Errorf(format string, args ...interface{}) {
	level.Error(logger).Log(
		"func", "Errorf",
		"msg", fmt.Sprintf(format, args...),
	)
}

func Fatal(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Fatal")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func FatalDepth(depth int, args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "FatalDepth")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Fatalln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Fatalln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Fatalf(format string, args ...interface{}) {
	level.Error(logger).Log(
		"func", "Fatalf",
		"msg", fmt.Sprintf(format, args...),
	)
}

func Exit(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Exit")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func ExitDepth(depth int, args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "ExitDepth")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Exitln(args ...interface{}) {
	var pairs []interface{}
	pairs = append(pairs, "func", "Exitln")
	for key, value := range args {
		pairs = append(pairs, key, value)
	}
	level.Error(logger).Log(pairs...)
}

func Exitf(format string, args ...interface{}) {
	level.Error(logger).Log(
		"func", "Exitf",
		"msg", fmt.Sprintf(format, args...),
	)
}
