package logger

type LogWriter interface {
	Write(data interface{}) error
}

type Logger struct {
	writeList []LogWriter
}

func (l *Logger) RegisterWriter(writer LogWriter)  {
	l.writeList = append(l.writeList, writer)
}

// 将一个data类型的数据写入日志
func (l *Logger) Log(data interface{})  {
	for _, writer := range l.writeList {
		writer.Write(data)
	}
}

func NewLogger() *Logger {
	return &Logger{}
}
