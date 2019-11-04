package logger

import (
	"fmt"
	"chapter7/logger/logger"
)

func createLogger() *Logger {
	l := NewLogger()
	cw := newConsoleWriter()
	l.RegisterWriter(cw)
	// 创建文件写入器
	fw := newFileWriter()
	// 设置文件名称
	if err := fw.SetFile("log.log"); err != nil {
		fmt.Println(err)
	}

	l.RegisterWriter(fw)
	return l
}

func main() {
	l := createLogger()
	l.Log("hello")
}
