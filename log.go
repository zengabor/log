package log

import (
	"fmt"
	"log"
	"os"
)

var (
	colors  bool
	_prefix string
)

func Printf(format string, v ...interface{}) {
	log.Printf(_prefix+format, v...)
}

func Errorf(format string, v ...interface{}) {
	if colors {
		format = "\x1b[93;41m" + format + "\x1b[0m"
	} else {
		format = "[ERROR] " + format
	}
	format = _prefix + format
	Printf(format, v...)
	if !colors {
		fmt.Fprintf(os.Stderr, format, v...) // maybe Stackdriver will pick it up?
	}
}

func EnableColors() {
	colors = true
}

func SetPrefix(prefix string) {
	_prefix = prefix
}
