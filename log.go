package log

import (
	golog "log"
)

var coloredErrors = false

func Printf(format string, v ...interface{}) {
	golog.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	if coloredErrors {
		format = "\x1b[93;41m" + format + "\x1b[0m"
	} else {
		format = "!ERROR: " + format
	}
	Printf(format, v...)
}

func SetColoredErrors(state bool) {
	coloredErrors = state
}
