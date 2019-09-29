package main

import "fmt"

type ParseError struct {
	Filename string
	Line int
}

func (e *ParseError) Error() string  {
	return fmt.Sprintf("%s:%d", e.Filename, e.Line)
}

// 创建一些解析错误
func newParseError(filename string, line int) error  {
	return &ParseError{Filename:filename, Line:line}
}

func main() {
	var e error
	e = newParseError("main.go", 1)
	fmt.Println(e.Error())
	// 根据错误的接口类型，获取详细的错误信息
	switch detail := e.(type) {
	case *ParseError:
		fmt.Printf("Filename:%s Line:%d", detail.Filename, detail.Line)
	default:
		fmt.Println("other error")
		
	}
}


