package log

type Level byte

func (l Level) String() string {
	switch l {
	case Off:
		return "OFF"
	case Fatal:
		return "FATAL"
	case Error:
		return "ERROR"
	case Warn:
		return "WARN"
	case Info:
		return "INFO"
	case Debug:
		return "DEBUG"
	default:
		return "UNKNOWN"
	}
}

const (
	Off Level = iota
	Fatal
	Error
	Warn
	Info
	Debug
)

//var DefaultLogger Logger = &inner{}

type Printer interface {
	Print(...interface{})
	Printf(string, ...interface{})
	Println(...interface{})
}

type Logger interface {
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
