package logger

import (
	"fmt"
	"os"
	"runtime/debug"

	"github.com/go-kit/log"
	"github.com/go-kit/log/level"
)

var Logger log.Logger

func Init() {

	baseLogger := log.NewLogfmtLogger(os.Stdout)

	Logger = log.With(
		baseLogger,
		"ts", log.DefaultTimestampUTC,
		"caller", log.Caller(4),
	)
}

func Debug(msg string, keyvals ...interface{}) {
	level.Debug(Logger).Log(
		append([]interface{}{"msg", msg}, keyvals...)...,
	)
}

func Info(msg string, keyvals ...interface{}) {
	level.Info(Logger).Log(
		append([]interface{}{"msg", msg}, keyvals...)...,
	)
}

func Error(msg string, err error, keyvals ...interface{}) {

	args := []interface{}{
		"msg", msg,
		"error", err,
	}

	args = append(args, keyvals...)

	level.Error(Logger).Log(args...)
}

func Trace(msg string, err error) {

	level.Error(Logger).Log(
		"msg", msg,
		"error", err,
		"trace", fmt.Sprintf("%s", debug.Stack()),
	)
}
