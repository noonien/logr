package logr

import (
	"fmt"
	"os"
)

type Level struct {
	Out Logger
}

func (l *Level) Log(e Entry) {
	l.Out.Log(e)
}

func (l *Level) log(lvl, s string) {
	l.Log(Entry{
		"level":   lvl,
		"message": s,
	})
}

func (l Level) Debug(v ...interface{}) {
	l.log("debug", fmt.Sprint(v...))
}

func (l Level) Debugf(format string, v ...interface{}) {
	l.log("debug", fmt.Sprintf(format, v...))
}

func (l Level) Debugln(v ...interface{}) {
	l.log("debug", fmt.Sprint(v...))
}

func (l Level) Info(v ...interface{}) {
	l.log("info", fmt.Sprint(v...))
}

func (l Level) Infof(format string, v ...interface{}) {
	l.log("info", fmt.Sprintf(format, v...))
}

func (l Level) Infoln(v ...interface{}) {
	l.log("info", fmt.Sprint(v...))
}

func (l Level) Warn(v ...interface{}) {
	l.log("warn", fmt.Sprint(v...))
}

func (l Level) Warnf(format string, v ...interface{}) {
	l.log("warn", fmt.Sprintf(format, v...))
}

func (l Level) Warnln(v ...interface{}) {
	l.log("warn", fmt.Sprint(v...))
}

func (l Level) Error(v ...interface{}) {
	l.log("error", fmt.Sprint(v...))
}

func (l Level) Errorf(format string, v ...interface{}) {
	l.log("error", fmt.Sprintf(format, v...))
}

func (l Level) Errorln(v ...interface{}) {
	l.log("error", fmt.Sprint(v...))
}

func (l Level) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log("panic", s)
	panic(s)
}

func (l Level) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.log("panic", s)
	panic(s)
}

func (l Level) Panicln(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log("panic", s)
	panic(s)
}

func (l Level) Fatal(v ...interface{}) {
	l.log("fatal", fmt.Sprint(v...))
	os.Exit(1)
}

func (l Level) Fatalf(format string, v ...interface{}) {
	l.log("fatal", fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l Level) Fatalln(v ...interface{}) {
	l.log("fatal", fmt.Sprint(v...))
	os.Exit(1)
}
