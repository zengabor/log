package log

import (
	"fmt"
	"log"
	"os"
)

var colors bool

func Printf(format string, v ...interface{}) {
	log.Printf(format, v...)
}

func Errorf(format string, v ...interface{}) {
	if colors {
		format = "\x1b[93;41m" + format + "\x1b[0m"
	} else {
		format = "[ERROR] " + format
	}
	Printf(format, v...)
	if !colors {
		fmt.Fprintf(os.Stderr, format, v...) // maybe Stackdriver will pick it up?
	}
}

func EnableColors() {
	colors = true
}
