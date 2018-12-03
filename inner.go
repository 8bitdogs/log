package log

import (
	"fmt"
	"io"
	l "log"
	"os"
)

type cpPrinter interface {
	Printer
	Copy(string) cpPrinter
}

type empty struct{}

func (*empty) Print(args ...interface{}) {}

func (*empty) Printf(format string, args ...interface{}) {}

func (*empty) Println(args ...interface{}) {}

func (e *empty) Copy(prefix string) cpPrinter { return e }

type inner struct {
	logger *l.Logger
	prefix string
	level  string
}

func newInner(out io.Writer, prefix string, level Level, flag int) *inner {
	return &inner{
		level:  fmt.Stringer(level).String(),
		prefix: prefix,
		logger: l.New(out, "", flag),
	}
}

func (i *inner) Print(args ...interface{}) {
	i.logger.Output(2, i.level+" "+i.prefix+" "+fmt.Sprint(args...))
}

func (i *inner) Printf(format string, args ...interface{}) {
	i.logger.Output(2, i.level+" "+i.prefix+" "+fmt.Sprintf(format, args...))
}

func (i *inner) Println(args ...interface{}) {
	i.logger.Output(2, i.level+" "+i.prefix+" "+fmt.Sprintln(args...))
}

func (i *inner) Copy(prefix string) cpPrinter {
	return &inner{logger: i.logger, level: i.level, prefix: prefix}
}

type logger struct {
	fatal cpPrinter
	err   cpPrinter
	warn  cpPrinter
	info  cpPrinter
	debug cpPrinter
}

func New(prefix string, level Level) Logger {
	return NewWithOutput(os.Stdout, prefix, level, l.LstdFlags)
}

func NewWithOutput(out io.Writer, prefix string, level Level, flag int) Logger {
	e := &empty{}
	l := &logger{fatal: e, err: e, warn: e, info: e, debug: e}
	if level == Off {
		return l
	}
	if level <= Off {
		return l
	}
	l.fatal = newInner(out, prefix, Fatal, flag)
	if level <= Fatal {
		return l
	}
	l.err = newInner(out, prefix, Error, flag)
	if level <= Error {
		return l
	}
	l.warn = newInner(out, prefix, Warn, flag)
	if level <= Warn {
		return l
	}
	l.info = newInner(out, prefix, Info, flag)
	if level <= Info {
		return l
	}
	l.debug = newInner(out, prefix, Debug, flag)
	return l
}

func (l *logger) Copy(prefix string) Logger {
	return &logger{
		fatal: l.fatal.Copy(prefix),
		err:   l.err.Copy(prefix),
		warn:  l.warn.Copy(prefix),
		info:  l.info.Copy(prefix),
		debug: l.debug.Copy(prefix),
	}
}

func (l *logger) Debug(args ...interface{})                 { l.debug.Print(args...) }
func (l *logger) Debugf(format string, args ...interface{}) { l.debug.Printf(format, args...) }
func (l *logger) Debugln(args ...interface{})               { l.debug.Println(args...) }

func (l *logger) Info(args ...interface{})                 { l.info.Print(args...) }
func (l *logger) Infof(format string, args ...interface{}) { l.info.Printf(format, args...) }
func (l *logger) Infoln(args ...interface{})               { l.info.Println(args...) }

func (l *logger) Warn(args ...interface{})                 { l.warn.Print(args...) }
func (l *logger) Warnf(format string, args ...interface{}) { l.warn.Printf(format, args...) }
func (l *logger) Warnln(args ...interface{})               { l.warn.Println(args...) }

func (l *logger) Error(args ...interface{})                 { l.err.Print(args...) }
func (l *logger) Errorf(format string, args ...interface{}) { l.err.Printf(format, args...) }
func (l *logger) Errorln(args ...interface{})               { l.err.Println(args...) }

func (l *logger) Fatal(args ...interface{}) {
	l.fatal.Print(args...)
	os.Exit(1)
}

func (l *logger) Fatalf(format string, args ...interface{}) {
	l.fatal.Printf(format, args...)
	os.Exit(1)
}

func (l *logger) Fatalln(args ...interface{}) {
	l.fatal.Println(args...)
	os.Exit(1)
}
