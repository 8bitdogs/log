package log

import (
	"errors"
	l "log"
	"os"
	"strings"
)

const (
	OffLevel Level = iota
	FatalLevel
	ErrorLevel
	WarnLevel
	InfoLevel
	DebugLevel
)

const (
	OffLevelString   = "OFF"
	FatalLevelString = "FATAL"
	ErrorLevelString = "ERROR"
	WarnLevelString  = "WARN"
	InfoLevelString  = "INFO"
	DebugLevelString = "DEBUG"
)

var ErrInvalidStringLevel = errors.New("string level doesn't match to any of levels")

var DefaultLogger Logger = NewWithOutput(os.Stderr, "", InfoLevel, l.LstdFlags)

type Level byte

func (l Level) String() string {
	switch l {
	case OffLevel:
		return OffLevelString
	case FatalLevel:
		return FatalLevelString
	case ErrorLevel:
		return ErrorLevelString
	case WarnLevel:
		return WarnLevelString
	case InfoLevel:
		return InfoLevelString
	case DebugLevel:
		return DebugLevelString
	default:
		return "UNKNOWN"
	}
}

type Printer interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

type Logger interface {
	Copy(string) Logger

	Debug(...interface{})
	Debugf(string, ...interface{})
	Debugln(...interface{})

	Info(...interface{})
	Infof(string, ...interface{})
	Infoln(...interface{})

	Warn(...interface{})
	Warnf(string, ...interface{})
	Warnln(...interface{})

	Error(...interface{})
	Errorf(string, ...interface{})
	Errorln(...interface{})

	Fatal(...interface{})
	Fatalf(string, ...interface{})
	Fatalln(...interface{})
}

type Printers interface {
	DebugPrinter() Printer
	InfoPrinter() Printer
	WarnPrinter() Printer
	ErrorPrinter() Printer
	FatalPrinter() Printer
}

func DefaultPrinters() Printers {
	prs, ok := DefaultLogger.(Printers)
	if !ok {
		return nil
	}
	return prs
}

func ParseLevel(level string) (Level, error) {
	switch strings.ToUpper(level) {
	case OffLevelString:
		return OffLevel, nil
	case FatalLevelString:
		return FatalLevel, nil
	case ErrorLevelString:
		return ErrorLevel, nil
	case WarnLevelString:
		return WarnLevel, nil
	case InfoLevelString:
		return InfoLevel, nil
	case DebugLevelString:
		return DebugLevel, nil
	default:
		return OffLevel, ErrInvalidStringLevel
	}
}

func Copy(prefix string) Logger { return DefaultLogger.Copy(prefix) }

func Debug(args ...interface{})                 { DefaultLogger.Debug(args...) }
func Debugf(format string, args ...interface{}) { DefaultLogger.Debugf(format, args...) }
func Debugln(args ...interface{})               { DefaultLogger.Debugln(args...) }

func Info(args ...interface{})                 { DefaultLogger.Info(args...) }
func Infof(format string, args ...interface{}) { DefaultLogger.Infof(format, args...) }
func Infoln(args ...interface{})               { DefaultLogger.Infoln(args...) }

func Warn(args ...interface{})                 { DefaultLogger.Warn(args...) }
func Warnf(format string, args ...interface{}) { DefaultLogger.Warnf(format, args...) }
func Warnln(args ...interface{})               { DefaultLogger.Warnln(args...) }

func Error(args ...interface{})                 { DefaultLogger.Error(args...) }
func Errorf(format string, args ...interface{}) { DefaultLogger.Errorf(format, args...) }
func Errorln(args ...interface{})               { DefaultLogger.Errorln(args...) }

func Fatal(args ...interface{}) {
	DefaultLogger.Fatal(args...)
	os.Exit(1)
}

func Fatalf(format string, args ...interface{}) {
	DefaultLogger.Fatalf(format, args...)
	os.Exit(1)
}

func Fatalln(args ...interface{}) {
	DefaultLogger.Fatalln(args...)
	os.Exit(1)
}
