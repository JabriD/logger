package logger

import (
	"fmt"
	"io"
)

type Logger interface {
	Log(...interface{})
}

func New(w io.Writer) Logger {
	return &logger{out: w}

}

type logger struct {
	out io.Writer
}

func (lg *logger) Log(args ...interface{}) {
	fmt.Fprint(lg.out, args...)
	fmt.Fprintln(lg.out)
}
