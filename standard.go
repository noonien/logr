package logr

import (
	"fmt"
	"os"
)

type Standard struct {
	Out Logger
}

func (l *Standard) Log(e Entry) {
	if l.Out != nil {
		l.Out.Log(e)
	}
}

func (l *Standard) log(s string) {
	l.Log(Entry{
		"message": s,
	})
}

func (l *Standard) Print(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

func (l *Standard) Printf(format string, v ...interface{}) {
	l.log(fmt.Sprintf(format, v...))
}

func (l *Standard) Println(v ...interface{}) {
	l.log(fmt.Sprint(v...))
}

func (l *Standard) Panic(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log(s)
	panic(s)
}

func (l *Standard) Panicf(format string, v ...interface{}) {
	s := fmt.Sprintf(format, v...)
	l.log(s)
	panic(s)
}

func (l *Standard) Panicln(v ...interface{}) {
	s := fmt.Sprint(v...)
	l.log(s)
	panic(s)
}

func (l *Standard) Fatal(v ...interface{}) {
	l.log(fmt.Sprint(v...))
	os.Exit(1)
}

func (l *Standard) Fatalf(format string, v ...interface{}) {
	l.log(fmt.Sprintf(format, v...))
	os.Exit(1)
}

func (l *Standard) Fatalln(v ...interface{}) {
	l.log(fmt.Sprint(v...))
	os.Exit(1)
}
