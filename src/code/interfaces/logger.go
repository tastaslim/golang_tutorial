package interfaces

type Logger interface {
	Log(message string)
}

type ConsoleLogger struct {
	Name string
}

type FileLogger struct {
	Name string
}

func (c ConsoleLogger) Log(message string) {
	c.Log(message)

}

func (f FileLogger) Log(message string) {
	f.Log(message)
}

func Logging(l Logger, message string) {
	l.Log(message)
}
